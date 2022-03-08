package utils

import (
	"os"
	"unicode/utf8"
)

//TODO to structure and move to a library

func StringInArray(val string, arr []string) (index int, contains bool) {
	for i, s := range arr {
		if s == val {
			return i, true
		}
	}
	return -1, false
}

func TrimLastChar(s string) string {
	r, size := utf8.DecodeLastRuneInString(s)
	if r == utf8.RuneError && (size == 0 || size == 1) {
		size = 0
	}
	return s[:len(s)-size]
}

func GetEnv(key, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultValue
}
