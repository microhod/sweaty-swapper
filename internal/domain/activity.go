package domain

import "fmt"

type ActivityType int

const (
	ActivityTypeRunning ActivityType = iota
	ActivityTypeWalking
	ActivityTypeHiking
	ActivityTypeCycling
	ActivityTypeGym
	ActivityTypeYoga
)

type Activity struct {
	ID          ActivityID `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Route       Route      `json:"route"`
}

type ActivityID struct {
	Source string
	ID     string
}

func (id ActivityID) MarshalText() ([]byte, error) {
	return []byte(fmt.Sprintf("%s:%s", id.Source, id.ID)), nil
}
