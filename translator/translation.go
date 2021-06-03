package translator

import (
	"unicode/utf8"
)

func Translate(input string) (string, error) {
	output := make([]rune, 0, utf8.RuneCountInString(input))

	for _, r := range input {
		replace, ok := mapping[r]
		if ok {
			output = append(output, []rune(replace)...)
		} else {
			output = append(output, r)
		}
	}

	return string(output), nil
}
