package main

import (
	"strings"
)

func findFirstStringInBracket(str string) string {
	start := strings.Index(str, "(")
	end := strings.Index(str, ")")

	if start >= 0 && end > start+1 {
		return str[start+1 : end]
	}

	return str
}
