package urlbuilder

import (
	"fmt"
	"net/url"
	"strings"
)

// Build constructs a full URL from host, port, and path
func Build(host, port, path string) (string, error) {
	// Remove trailing slash from host
	host = strings.TrimRight(host, "/")

	// Include port if provided
	fullHost := host
	if port != "" {
		fullHost = fmt.Sprintf("%s:%s", host, port)
	}

	// Ensure path starts with a slash
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}

	// Parse and validate URL
	u, err := url.Parse(fullHost + path)
	if err != nil {
		return "", err
	}

	return u.String(), nil
}
