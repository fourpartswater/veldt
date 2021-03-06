package elastic

import (
	"fmt"

	"github.com/unchartedsoftware/veldt"
	"github.com/unchartedsoftware/veldt/binning"
)

// Count represents an elasticsearch implementation of the count tile.
type Count struct {
	Elastic
	Bivariate
}

// NewCountTile instantiates and returns a new tile struct.
func NewCountTile(host, port string) veldt.TileCtor {
	return func() (veldt.Tile, error) {
		t := &Count{}
		t.Host = host
		t.Port = port
		return t, nil
	}
}

// Parse parses the provided JSON object and populates the tiles attributes.
func (t *Count) Parse(params map[string]interface{}) error {
	return t.Bivariate.Parse(params)
}

// Create generates a tile from the provided URI, tile coordinate and query
// parameters.
func (t *Count) Create(uri string, coord *binning.TileCoord, query veldt.Query) ([]byte, error) {
	// create search service
	search, err := t.CreateSearchService(uri)
	if err != nil {
		return nil, err
	}

	// create root query
	q, err := t.CreateQuery(query)
	if err != nil {
		return nil, err
	}
	// add tiling query
	q.Must(t.Bivariate.GetQuery(coord))
	// set the query
	search.Query(q)

	// send query
	res, err := search.Do()
	if err != nil {
		return nil, err
	}

	return []byte(fmt.Sprintf(`{"count":%d}`, res.Hits.TotalHits)), nil
}
