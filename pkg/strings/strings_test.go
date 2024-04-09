package strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToTitleCase(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{
			input:    "single",
			expected: "Single",
		},
		{
			input:    "multiple words",
			expected: "Multiple Words",
		},
		{
			input:    "multiple words	with   differing\nspaces",
			expected: "Multiple Words	With   Differing\nSpaces",
		},
		{
			input:    "étude",
			expected: "Étude",
		},
		{
			input:    "???",
			expected: "???",
		},
		{
			input:    "",
			expected: "",
		},
	}
	for _, tc := range testCases {
		assert.Equal(t, tc.expected, ToTitleCase(tc.input))
	}
}
