package gpx

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/tkrajina/gpxgo/gpx"

	"github.com/microhod/sweaty-swapper/internal/domain"
)

type Route struct {
	data *gpx.GPX
	fetch func(context.Context) ([]byte, error)
}

func ParseRoute(b []byte) (*Route, error) {
	data, err := gpx.ParseBytes(b)
	if err != nil {
		return nil, err
	}
	return &Route{data: data}, nil
}

func NewLazyLoadedRoute(f func(context.Context) ([]byte, error)) *Route {
	return &Route{
		fetch: f,
	}
}

func (r *Route) Type() domain.RouteType {
	return domain.RouteTypeGpx
}

func (r *Route) Data(ctx context.Context) ([]byte, error) {
	if r.data == nil && r.fetch != nil {
		if err := r.load(ctx); err != nil {
			return nil, err
		}
	}

	return r.data.ToXml(gpx.ToXmlParams{
		Indent: false,
	})
}

func (r *Route) Empty(ctx context.Context) (bool, error) {
	if r.data == nil && r.fetch != nil {
		if err := r.load(ctx); err != nil {
			return true, err
		}
	}

	for _, track := range r.data.Tracks {
		for _, segment := range track.Segments {
			if len(segment.Points) > 0 {
				return false, nil
			}
		}
	}
	return true, nil
}

func (r *Route) load(ctx context.Context) error {
	data, err := r.fetch(ctx)
	if err != nil {
		return fmt.Errorf("lazy loading gpx: %s", err)
	}
	parsed, err := gpx.ParseBytes(data)
	if err != nil {
		return fmt.Errorf("parsing lazy loaded gpx: %w", err)
	}

	r.data = parsed
	return nil
}

func (r *Route) MarshalJSON() ([]byte, error) {
	if empty, _ := r.Empty(context.Background()); empty {
		return json.Marshal(nil)
	}
	data, err := r.Data(context.Background())
	if err != nil {
		return nil, fmt.Errorf("converting gpx to xml: %w", err)
	}
	return json.Marshal(data)
}
