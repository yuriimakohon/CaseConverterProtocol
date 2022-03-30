package ccp

import "strings"

func Camel(input string) string {
	input = strings.ToUpper(input)

	output := strings.Builder{}
	output.Grow(len(input))

	for i, char := range input {
		// Lower every second alphabetic character
		if 'A' <= char && char <= 'Z' && i%2 == 1 {
			output.WriteByte(byte(char + 32))
			continue
		}
		// Or just write character to resulting string
		output.WriteByte(byte(char))
	}

	return output.String()
}
