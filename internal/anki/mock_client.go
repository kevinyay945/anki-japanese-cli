package anki

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

// MockHTTPClient is a mock implementation of the http.Client for testing
type MockHTTPClient struct {
	// DoFunc will be executed when Do is called
	DoFunc func(req *http.Request) (*http.Response, error)
}

// Do implements the http.Client interface
func (m *MockHTTPClient) Do(req *http.Request) (*http.Response, error) {
	return m.DoFunc(req)
}

// NewMockHTTPClient creates a new mock HTTP client with the given response
func NewMockHTTPClient(statusCode int, responseBody string, err error) *MockHTTPClient {
	return &MockHTTPClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			if err != nil {
				return nil, err
			}
			return &http.Response{
				StatusCode: statusCode,
				Body:       io.NopCloser(strings.NewReader(responseBody)),
			}, nil
		},
	}
}

// NewMockHTTPClientWithRequestCheck creates a new mock HTTP client that checks the request
func NewMockHTTPClientWithRequestCheck(statusCode int, responseBody string, err error, requestCheck func(req *http.Request) bool) *MockHTTPClient {
	return &MockHTTPClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			if err != nil {
				return nil, err
			}
			if requestCheck != nil && !requestCheck(req) {
				return &http.Response{
					StatusCode: http.StatusBadRequest,
					Body:       io.NopCloser(strings.NewReader(`{"error": "Invalid request"}`)),
				}, nil
			}
			return &http.Response{
				StatusCode: statusCode,
				Body:       io.NopCloser(strings.NewReader(responseBody)),
			}, nil
		},
	}
}

// NewMockHTTPClientWithMultipleResponses creates a new mock HTTP client that returns different responses for different requests
func NewMockHTTPClientWithMultipleResponses(responses map[string]struct {
	StatusCode int
	Body       string
	Error      error
}) *MockHTTPClient {
	return &MockHTTPClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			// Read the request body
			body, err := io.ReadAll(req.Body)
			if err != nil {
				return nil, err
			}
			// Close the original body
			req.Body.Close()
			// Create a new body with the same content for further reading
			req.Body = io.NopCloser(bytes.NewBuffer(body))

			// Parse the request to get the action
			var request Request
			if err := json.Unmarshal(body, &request); err != nil {
				return &http.Response{
					StatusCode: http.StatusBadRequest,
					Body:       io.NopCloser(strings.NewReader(`{"error": "Invalid JSON"}`)),
				}, nil
			}

			// Get the response for this action
			response, ok := responses[request.Action]
			if !ok {
				return &http.Response{
					StatusCode: http.StatusBadRequest,
					Body:       io.NopCloser(strings.NewReader(`{"error": "Unknown action"}`)),
				}, nil
			}

			if response.Error != nil {
				return nil, response.Error
			}

			return &http.Response{
				StatusCode: response.StatusCode,
				Body:       io.NopCloser(strings.NewReader(response.Body)),
			}, nil
		},
	}
}