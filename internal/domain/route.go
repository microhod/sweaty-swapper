package domain

type Route struct {
	Type RouteType `json:"type"`
	Data []byte    `json:"data"`
}

type RouteType int

const (
	RouteTypeGpx RouteType = iota
)
