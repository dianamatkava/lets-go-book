package validator

import (
	"slices"
	"strings"
	"unicode/utf8"
)

func NoBlank(val string) bool {
	return strings.TrimSpace(val) != ""
}

func MaxChar(val string, maxChar int) bool {
	return utf8.RuneCountInString(val) <= maxChar
}

func Contains(val int, permitedValues []int) bool {
	return slices.Contains(permitedValues, val)
}