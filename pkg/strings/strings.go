package strings

import (
	"unicode"
)

func ToTitleCase(input string) string {
	inputRunes := []rune(input)
	outputRunes := make([]rune, len(inputRunes))

	// initialising this to true ensures the first character is
	// capitalised if it's non-empty
	lastWasSpace := true
	for index, r := range inputRunes {
		if unicode.IsSpace(r) {
			lastWasSpace = true
		} else if lastWasSpace {
			r = unicode.ToTitle(r)
			lastWasSpace = false
		}

		outputRunes[index] = r
	}
	return string(outputRunes)
}