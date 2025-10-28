package urlbuilder_test

import (
	"testing"

	"github.com/pujidjayanto/choochoohub/api-gateway/pkg/urlbuilder"
)

func TestBuildURL(t *testing.T) {
	tests := []struct {
		name     string
		host     string
		port     string
		path     string
		expected string
		wantErr  bool
	}{
		{
			name:     "normal case with port",
			host:     "http://localhost",
			port:     "8080",
			path:     "/v1/signup",
			expected: "http://localhost:8080/v1/signup",
		},
		{
			name:     "host with trailing slash",
			host:     "http://localhost/",
			port:     "8080",
			path:     "/v1/signup",
			expected: "http://localhost:8080/v1/signup",
		},
		{
			name:     "path without leading slash",
			host:     "http://localhost",
			port:     "8080",
			path:     "v1/signup",
			expected: "http://localhost:8080/v1/signup",
		},
		{
			name:     "no port",
			host:     "http://example.com",
			port:     "",
			path:     "/api",
			expected: "http://example.com/api",
		},
		{
			name:    "invalid URL",
			host:    "://invalid-host",
			port:    "",
			path:    "/api",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := urlbuilder.Build(tt.host, tt.port, tt.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("BuildURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.expected && !tt.wantErr {
				t.Errorf("BuildURL() = %v, expected %v", got, tt.expected)
			}
		})
	}
}
