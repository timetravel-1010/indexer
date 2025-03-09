package regex

import (
	"regexp"
)

const (
	regexEmailAddress = `[\w._%+-]+@[\w.-]+\.[A-Za-z]{2,}`
	regexName         = `^[a-zA-ZÀ-ÿ0-9 ()-]*$`
)

var (
	namePattern      = regexp.MustCompile(regexName)
	emailAddrPattern = regexp.MustCompile(regexEmailAddress)
)

// getAddress
func GetEmailAddress(s string) string {
	return getStringByRegexp(s, emailAddrPattern)
}

// parseNames
func GetName(s string) string {
	return getStringByRegexp(s, namePattern)
}

// GetStringByRegexp
func getStringByRegexp(s string, regex *regexp.Regexp) string {
	matches := regex.FindStringSubmatch(s)
	if len(matches) == 2 {
		return matches[1]
	} else if len(matches) == 1 {
		return matches[0]
	}
	return ""
}

// GetStringsByRegexp
func getStringsByRegexp(s string, regex string) []string {
	return regexp.MustCompile(regex).FindAllString(s, -1)
}
