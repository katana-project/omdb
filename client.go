package omdb

import (
	"context"
	"encoding/json"
	"net/http"
)

// DefaultServerBaseURL is the default OMDB API base URL.
const DefaultServerBaseURL = "https://www.omdbapi.com/"

// NewDefaultClient creates a new Client with the DefaultServerBaseURL base URL and a WithToken ClientOption.
func NewDefaultClient(token string, opts ...ClientOption) (*Client, error) {
	return NewClient(DefaultServerBaseURL, append(opts, WithToken(token))...)
}

// WithToken authenticates the client's requests with an `apikey` query parameter.
func WithToken(token string) ClientOption {
	return WithRequestEditorFn(func(_ context.Context, req *http.Request) error {
		values := req.URL.Query()
		values.Set("apikey", token)

		req.URL.RawQuery = values.Encode()
		return nil
	})
}

// Response is an abstraction of a typed http.Response wrapper.
type Response interface {
	// Status returns the status text.
	Status() string
	// StatusCode returns the status code.
	StatusCode() int
}

// ParseBody parses JSON bytes into an untyped map.
func ParseBody(b []byte) (map[string]interface{}, error) {
	var d map[string]interface{}
	if err := json.Unmarshal(b, &d); err != nil {
		return nil, err
	}

	return d, nil
}
