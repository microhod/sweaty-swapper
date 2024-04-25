package domain

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/microhod/sweaty-swapper/pkg/strings"
)

type ActivityType string

const (
	ActivityTypeRun   ActivityType = "run"
	ActivityTypeWalk  ActivityType = "walk"
	ActivityTypeHike  ActivityType = "hike"
	ActivityTypeCycle ActivityType = "cycle"
	ActivityTypeGym   ActivityType = "gym"
	ActivityTypeYoga  ActivityType = "yoga"
)

type Activity struct {
	IDs         ActivityIDs   `json:"ids"`
	Type        ActivityType  `json:"type"`
	Title       string        `json:"title"`
	Description string        `json:"description"`
	CreatedAt   time.Time     `json:"created_at"`
	Stats       ActivityStats `json:"stats"`
	Route       Route         `json:"route,omitempty"`
}

func (a Activity) DefaultTitle() string {
	var timeOfDay string

	hour := a.CreatedAt.Hour()
	switch {
	case hour < 12:
		timeOfDay = "Morning"
	case hour < 14:
		timeOfDay = "Lunch"
	case hour < 18:
		timeOfDay = "Afternoon"
	default:
		timeOfDay = "Evening"
	}

	return fmt.Sprintf("%s %s", timeOfDay, strings.ToTitleCase(string(a.Type)))
}

type ActivityIDs map[Platform]string

// Platform is the name of the service storing the activity data e.g. strava.com
// can be any string e.g. domain name might be a sensible choice if the platform has a website
type Platform string

type ActivityStats struct {
	TotalTime     ActivityDuration
	TotalDistance Meters
}

type ActivityDuration time.Duration

func (ad ActivityDuration) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Duration(ad).String())
}

// Meters is distance expressed in meters
type Meters float64
