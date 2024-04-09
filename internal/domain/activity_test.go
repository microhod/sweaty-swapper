package domain

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestActivityDefaultTitle(t *testing.T) {
	midnight := time.Date(2024, 4, 8, 0, 0, 0, 0, time.Local)

	testCases := []struct {
		name     string
		activity Activity
		expected string
	}{
		{
			name: "midnight run",
			activity: Activity{
				Type: ActivityTypeRun,
				CreatedAt: midnight,
			},
			expected: "Morning Run",
		},
		{
			name: "11am run",
			activity: Activity{
				Type: ActivityTypeRun,
				CreatedAt: midnight.Add(11 * time.Hour),
			},
			expected: "Morning Run",
		},
		{
			name: "12pm run",
			activity: Activity{
				Type: ActivityTypeRun,
				CreatedAt: midnight.Add(12 * time.Hour),
			},
			expected: "Lunch Run",
		},
		{
			name: "2pm run",
			activity: Activity{
				Type: ActivityTypeRun,
				CreatedAt: midnight.Add(14 * time.Hour),
			},
			expected: "Afternoon Run",
		},
		{
			name: "6pm run",
			activity: Activity{
				Type: ActivityTypeRun,
				CreatedAt: midnight.Add(18 * time.Hour),
			},
			expected: "Evening Run",
		},
		{
			name: "11pm run",
			activity: Activity{
				Type: ActivityTypeRun,
				CreatedAt: midnight.Add(23 * time.Hour),
			},
			expected: "Evening Run",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, tc.activity.DefaultTitle())
		})
	}
}
