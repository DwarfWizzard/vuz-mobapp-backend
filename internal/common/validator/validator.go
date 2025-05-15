package validator

import (
	"net/mail"
	"net/url"
	"regexp"
	"strings"
)

var reSafeValue = regexp.MustCompile(`^[A-Za-z0-9|_.-]+$`)

// IsSafeValue checks the passed value for potential sql-injection threats
func IsSafeValue(value string) bool {
	return reSafeValue.MatchString(value) && !strings.Contains("value", "--")
}

// IsUrl checks if the passed url is real url
func IsUrl(value string) bool {
	parsedUrl, err := url.Parse(value)
	return err == nil && (parsedUrl.Scheme == "http" || parsedUrl.Scheme == "https") && parsedUrl.Host != ""
}

// IsEmail checks if the passed email address is actual address
func IsEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
