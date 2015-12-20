package elastic

import (
	"encoding/json"
	"errors"
	"fmt"

	"gopkg.in/olivere/elastic.v3"

	"github.com/unchartedsoftware/prism/generation/tile"
)

// GetTopicCountParams returns a map of tiling parameters.
func GetTopicCountParams(tileReq *tile.Request) map[string]tile.Param {
	return map[string]tile.Param{
		"binning": NewBinningParams(tileReq),
		"topic": NewTopicParams(tileReq),
		"time": NewTimeParams(tileReq),
	}
}

// GetTopicCountTile returns a marshalled tile containing topics and counts.
func GetTopicCountTile(tileReq *tile.Request, params map[string]tile.Param) ([]byte, error) {
	binning, _ := params["binning"].(*BinningParams)
	time, _ := params["time"].(*TimeParams)
	topic, _ := params["topic"].(*TopicParams)
	if binning == nil {
		return nil, errors.New("No binning information has been provided")
	}
	if topic == nil {
		return nil, errors.New("No topics have been provided")
	}
	// get client
	client, err := getClient(tileReq.Endpoint)
	if err != nil {
		return nil, err
	}
	// create x and y range queries
	boolQuery := elastic.NewBoolQuery().Must(
		binning.GetXQuery(),
		binning.GetYQuery())
	// if time params are provided, add time range query
	if time != nil {
		boolQuery.Must(time.GetTimeQuery())
	}
	// build query
	query := client.
		Search(tileReq.Index).
		Size(0).
		Query(boolQuery)
	// add all filter aggregations
	topicAggs := topic.GetTopicAggregations()
	for topic, topicAgg := range topicAggs {
		query.Aggregation(topic, topicAgg)
	}
	// send query
	result, err := query.Do()
	if err != nil {
		return nil, err
	}
	// build map of topics and counts
	topicCounts := make(map[string]int64)
	for _, topic := range topic.Topics {
		filter, ok := result.Aggregations.Filter(topic)
		if !ok {
			return nil, fmt.Errorf("Filter aggregation '%s' was not found in response", topic)
		}
		if filter.DocCount > 0 {
			topicCounts[topic] = filter.DocCount
		}
	}
	// marshal results map
	return json.Marshal(topicCounts)
}
