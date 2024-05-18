package gpx

import (
	"context"
	"embed"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

//go:embed testdata
var testDataFS embed.FS

func TestRoute_Empty(t *testing.T) {
	testCases := []struct {
		name     string
		file     string
		expected bool
	}{
		{
			name:     "empty",
			file:     "testdata/empty.xml",
			expected: true,
		},
		{
			name:     "not empty",
			file:     "testdata/non_empty.xml",
			expected: false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			xml, err := testDataFS.ReadFile(tc.file)
			require.NoError(t, err)

			route, err := ParseRoute(xml)
			require.NoError(t, err)
			actual, err := route.Empty(context.Background())
			require.NoError(t, err)

			assert.Equal(t, tc.expected, actual)
		})
	}
}

func TestRoute_MarshalJSON(t *testing.T) {
	testCases := []struct {
		name string
		file string
	}{
		{
			name: "empty",
			file: "testdata/empty.xml",
		},
		{
			name: "not empty",
			file: "testdata/non_empty.xml",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			xml, err := testDataFS.ReadFile(tc.file)
			require.NoError(t, err)

			route, err := ParseRoute(xml)
			require.NoError(t, err)

			_, err = json.Marshal(route)
			assert.NoError(t, err)
		})
	}
}
