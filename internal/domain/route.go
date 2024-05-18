package domain

import "context"

type Route interface {
	Type() RouteType
	Data(context.Context) ([]byte, error)
	Empty(context.Context) (bool, error)
}

type RouteType int

const (
	RouteTypeGpx RouteType = iota
)
