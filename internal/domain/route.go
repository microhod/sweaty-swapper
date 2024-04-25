package domain

type Route interface {
	Type() RouteType
	Data() ([]byte, error)
	Empty() bool
}

type RouteType int

const (
	RouteTypeGpx RouteType = iota
)
