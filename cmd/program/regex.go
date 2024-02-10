package program

import "regexp"

const (
	regexEmailAddress = `[\w._%+-]+@[\w.-]+\.[A-Za-z]{2,}`
	regexName         = `^[a-zA-ZÀ-ÿ0-9 ()-]*$`
)

// parseAddresses
func parseAddresses(s string) []string {
	return GetStringsByRegexp(s, regexEmailAddress)
}

// parseNames
func parseNames(s string) []string {
	return GetStringsByRegexp(s, regexName)
}

// GetStringsByRegexp
func GetStringsByRegexp(s string, regex string) []string {
	return regexp.MustCompile(regex).FindAllString(s, -1)
}
