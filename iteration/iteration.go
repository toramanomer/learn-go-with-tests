package iteration

import "strings"

func Repeat(char string, count int) string {
	var stringBuilder strings.Builder
	for range count {
		stringBuilder.WriteString(char)
	}
	return stringBuilder.String()
}
