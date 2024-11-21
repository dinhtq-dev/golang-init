package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// HTTPClient with custom retry logic
type HTTPClient struct {
	client  *http.Client
	maxRetries int
}

// NewHTTPClient creates a new HTTP client with retry functionality
func NewHTTPClient(maxRetries int) *HTTPClient {
	return &HTTPClient{
		client: &http.Client{
			Timeout: time.Second * 10,
		},
		maxRetries: maxRetries,
	}
}

// SendRequest sends an HTTP request and retries on failure
func (c *HTTPClient) SendRequest(url string, method string, body interface{}) (*http.Response, error) {
	var requestBody []byte
	if body != nil {
		var err error
		requestBody, err = json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal request body: %v", err)
		}
	}

	var resp *http.Response
	var err error
	for i := 0; i < c.maxRetries; i++ {
		req, err := http.NewRequest(method, url, bytes.NewBuffer(requestBody))
		if err != nil {
			return nil, fmt.Errorf("failed to create request: %v", err)
		}
		resp, err = c.client.Do(req)
		if err == nil {
			return resp, nil
		}
		time.Sleep(time.Second * time.Duration(i)) // Retry with backoff
	}

	return nil, fmt.Errorf("request failed after %d retries: %v", c.maxRetries, err)
}
