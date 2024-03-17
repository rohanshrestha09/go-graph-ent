package utils

import (
	"regexp"
	"strings"
)

func ToKebabCase(s string) string {
	// Convert to lowercase
	s = strings.ToLower(s)

	// Replace non-alphanumeric characters with hyphens
	re := regexp.MustCompile("[^a-z0-9]+")
	s = re.ReplaceAllString(s, "-")

	// Collapse consecutive hyphens into a single hyphen
	re = regexp.MustCompile("-+")
	s = re.ReplaceAllString(s, "-")

	// Trim leading and trailing hyphens
	s = strings.Trim(s, "-")

	return s
}
