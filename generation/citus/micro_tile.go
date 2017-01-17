package citus

import (
	"github.com/unchartedsoftware/prism"
	"github.com/unchartedsoftware/prism/binning"
	"github.com/unchartedsoftware/prism/tile"
)

type MicroTile struct {
	Bivariate
	Tile
	TopHits
	tile.Micro
}

func NewMicroTile(host, port string) prism.TileCtor {
	return func() (prism.Tile, error) {
		m := &MicroTile{}
		m.Host = host
		m.Port = port
		return m, nil
	}
}

func (m *MicroTile) Parse(params map[string]interface{}) error {
	err := m.Bivariate.Parse(params)
	if err != nil {
		return err
	}
	err = m.TopHits.Parse(params)
	if err != nil {
		return err
	}
	err = m.Micro.Parse(params)
	if err != nil {
		return err
	}
	// parse includes
	m.TopHits.IncludeFields = m.Micro.ParseIncludes(
		m.TopHits.IncludeFields,
		m.Bivariate.XField,
		m.Bivariate.YField)
	return nil
}

func (m *MicroTile) Create(uri string, coord *binning.TileCoord, query prism.Query) ([]byte, error) {
	// Initialize the tile processing.
	client, citusQuery, err := m.InitliazeTile(uri, query)

	// add tiling query
	citusQuery = m.Bivariate.AddQuery(coord, citusQuery)

	// get aggs
	citusQuery = m.TopHits.AddAggs(citusQuery)

	// send query
	res, err := client.Query(citusQuery.GetQuery(false), citusQuery.QueryArgs...)
	if err != nil {
		return nil, err
	}

	// get top hits
	hits, err := m.TopHits.GetTopHits(res)
	if err != nil {
		return nil, err
	}

	// convert to point array
	points := make([]float32, len(hits)*2)
	for i, hit := range hits {
		// get hit x/y in tile coords
		x, y, ok := m.Bivariate.GetXY(hit)
		if !ok {
			continue
		}
		// add to point array
		points[i*2] = x
		points[i*2+1] = y
	}

	// encode and return results
	return m.Micro.Encode(hits, points)
}