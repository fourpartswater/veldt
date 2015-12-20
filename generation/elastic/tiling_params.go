package elastic

import (
	"fmt"

	"gopkg.in/olivere/elastic.v3"

	"github.com/unchartedsoftware/prism/binning"
	"github.com/unchartedsoftware/prism/generation/tile"
	"github.com/unchartedsoftware/prism/util/json"
)

// TilingParams represents params for binning the data within the tile.
type TilingParams struct {
	X    string
	Y    string
	MinX int64
	MaxX int64
	MinY int64
	MaxY int64
}

// NewTilingParams parses the params map returns a pointer to the param struct.
func NewTilingParams(tileReq *tile.Request) *TilingParams {
	params := tileReq.Params
	extents := &binning.Bounds{
		TopLeft: &binning.Coord{
			X: json.GetNumberDefault(params, "minX", 0.0),
			Y: json.GetNumberDefault(params, "maxX", binning.MaxPixels),
		},
		BottomRight: &binning.Coord{
			X: json.GetNumberDefault(params, "maxY", 0.0),
			Y: json.GetNumberDefault(params, "minY", binning.MaxPixels),
		},
	}
	bounds := binning.GetTileBounds(tileReq.TileCoord, extents)
	return &TilingParams{
		X: json.GetStringDefault(params, "x", "pixel.x"),
		Y: json.GetStringDefault(params, "y", "pixel.y"),
		MinX: int64(bounds.TopLeft.X),
		MaxX: int64(bounds.BottomRight.X - 1),
		MinY: int64(bounds.TopLeft.Y),
		MaxY: int64(bounds.BottomRight.Y - 1),
	}
}

// GetHash returns a string hash of the parameter state.
func (p *TilingParams) GetHash() string {
	return fmt.Sprintf("%s:%s:%d:%d:%d:%d",
		p.X,
		p.Y,
		p.MinX,
		p.MaxX,
		p.MinY,
		p.MaxY)
}

// GetXQuery returns an elastic query.
func (p *TilingParams) GetXQuery() *elastic.RangeQuery {
	return elastic.NewRangeQuery(p.X).
		Gte(p.MinX).
		Lte(p.MaxX)
}

// GetYQuery returns an elastic query.
func (p *TilingParams) GetYQuery() *elastic.RangeQuery {
	return elastic.NewRangeQuery(p.Y).
		Gte(p.MinY).
		Lte(p.MaxY)
}
