package camel

import "strings"

func Camel(str string) string {
	var sb strings.Builder
	for _, r := range str {
		if r >= 'A' && r <= 'Z' {
			sb.WriteRune('_')
		}
		sb.WriteRune(r)
	}
	return strings.ToLower(sb.String())
}
