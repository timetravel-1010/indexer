package regex

import "regexp"

const (
	regexEmailAddress = `[\w._%+-]+@[\w.-]+\.[A-Za-z]{2,}`
	regexName         = `^[a-zA-ZÀ-ÿ0-9 ()-]*$`
)

// parseAddresses
func ParseAddresses(s string) []string {
	return getStringsByRegexp(s, regexEmailAddress)
}

// parseNames
func ParseNames(s string) []string {
	return getStringsByRegexp(s, regexName)
}

// GetStringsByRegexp
func getStringsByRegexp(s string, regex string) []string {
	return regexp.MustCompile(regex).FindAllString(s, -1)
}
