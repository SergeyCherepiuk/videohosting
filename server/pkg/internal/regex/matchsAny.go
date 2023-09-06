package regex

import "regexp"

func MatchesAny(s string, patterns []string) bool {
	for _, pattern := range patterns {
		if matches, err := regexp.MatchString(pattern, s); matches && err == nil {
			return true
		}
	}
	return false
}