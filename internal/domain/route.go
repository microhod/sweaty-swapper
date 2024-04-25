package domain

import "encoding/json"

type Route interface {
	Type() RouteType
	Data() ([]byte, error)
	Empty() bool
}

type RouteType int

const (
	RouteTypeGpx RouteType = iota
)

type GPXRoute struct {
	data []byte
}

func NewGPXRoute(d []byte) GPXRoute {
	return GPXRoute{data: d}
}

func (g GPXRoute) Type() RouteType {
	return RouteTypeGpx
}

func (g GPXRoute) Data() ([]byte, error) {
	return g.data, nil
}

func (g GPXRoute) Empty() bool {
	// not implemented
	return false
}

func (g GPXRoute) MarshalJSON() ([]byte, error) {
	if g.Empty() {
		return nil, nil
	}
	return json.Marshal(g.data)
}
