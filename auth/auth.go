package auth

import (
	"errors"
	"net/http"
	"strings"
)

// GetAPIKey extract api key from http header
// Example:
// Authorization: ApiKey {api_key}
func GetAPIKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("no auth header")
	}

	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("invalid auth header")
	}

	if vals[0] != "ApiKey" {
		return "", errors.New("invalid auth header")
	}

	return vals[1], nil
}
