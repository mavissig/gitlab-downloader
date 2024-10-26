package utils

import "strings"

func IsEmpty(s string) bool {
	pool := []string{
		"", " ", "\n", "\t", "\r", "\v",
		"empty", "nil", "null", "undefined", "unknown", "none",
	}

	s = strings.ToLower(s)

	for _, v := range pool {
		if s == v {
			return true
		}
	}

	return false
}
