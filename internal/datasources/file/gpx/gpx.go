package gpx

import (
	"encoding/json"
	"fmt"

	"github.com/tkrajina/gpxgo/gpx"

	"github.com/microhod/sweaty-swapper/internal/domain"
)

type Route struct {
	data *gpx.GPX
}

func ParseRoute(b []byte) (Route, error) {
	data, err := gpx.ParseBytes(b)
	if err != nil {
		return Route{}, err
	}
	return Route{data: data}, nil
}

func (r Route) Type() domain.RouteType {
	return domain.RouteTypeGpx
}

func (r Route) Data() ([]byte, error) {
	return r.data.ToXml(gpx.ToXmlParams{
		Indent: false,
	})
}

func (r Route) Empty() bool {
	for _, track := range r.data.Tracks {
		for _, segment := range track.Segments {
			if len(segment.Points) > 0 {
				return false
			}
		}
	}
	return true
}

func (r Route) MarshalJSON() ([]byte, error) {
	if r.Empty() {
		return json.Marshal(nil)
	}
	data, err := r.Data()
	if err != nil {
		return nil, fmt.Errorf("converting gpx to xml: %w", err)
	}
	return json.Marshal(data)
}
