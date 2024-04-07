package domain

import (
	"fmt"
	"strings"
	"time"
)

type ActivityType string

const (
	ActivityTypeRunning ActivityType = "running"
	ActivityTypeWalking ActivityType = "walking"
	ActivityTypeHiking  ActivityType = "hiking"
	ActivityTypeCycling ActivityType = "cycling"
	ActivityTypeGym     ActivityType = "gym"
	ActivityTypeYoga    ActivityType = "yoga"
)

type Activity struct {
	ID          ActivityID   `json:"id"`
	Type        ActivityType `json:"type"`
	Title       string       `json:"title"`
	Description string       `json:"description"`
	CreatedAt   time.Time    `json:"created_at"`
	Route       Route        `json:"route"`
}

func (a Activity) DefaultTitle() string {
	return strings.ToUpper(string(a.Type))
}

type ActivityID struct {
	Source string
	ID     string
}

func (id ActivityID) MarshalText() ([]byte, error) {
	return []byte(fmt.Sprintf("%s:%s", id.Source, id.ID)), nil
}
