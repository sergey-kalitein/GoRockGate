package helpers

import (
	"regexp"
	"strings"
)

func IsValidDomainName(domainName string) bool {
	RegExp := regexp.MustCompile(`^(([a-zA-Z]{1})|([a-zA-Z]{1}[a-zA-Z]{1})|([a-zA-Z]{1}[0-9]{1})|([0-9]{1}[a-zA-Z]{1})|([a-zA-Z0-9][a-zA-Z0-9-_]{1,61}[a-zA-Z0-9]))\.([a-zA-Z]{2,6}|[a-zA-Z0-9-]{2,30}\.[a-zA-Z]{2,4})$`)
	return RegExp.MatchString(domainName)
}

// Strips protocol from the domain name
func StripDomainName(webOriginDomain string) string {
	domainStripped := strings.TrimSpace(strings.ToLower(webOriginDomain))
	return regexp.MustCompile(`^(?:http|https)://(.*?)`).ReplaceAllString(domainStripped, `$1`)
}
