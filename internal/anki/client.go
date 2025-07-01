package anki

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"anki-japanese-cli/internal/config"
)

// HTTPClient interface for the HTTP client
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

const (
	// DefaultTimeout is the default timeout for HTTP requests
	DefaultTimeout = 30 * time.Second
	// DefaultRetries is the default number of retries for failed requests
	DefaultRetries = 3
	// DefaultRetryDelay is the default delay between retries
	DefaultRetryDelay = 1 * time.Second
	// APIVersion is the version of the Anki Connect API
	APIVersion = 6
)

// Client represents an Anki Connect API client
type Client struct {
	config     *config.AnkiConfig
	httpClient HTTPClient
	retries    int
	retryDelay time.Duration
}

// Request represents an Anki Connect API request
type Request struct {
	Action  string      `json:"action"`
	Version int         `json:"version"`
	Params  interface{} `json:"params,omitempty"`
}

// Response represents an Anki Connect API response
type Response struct {
	Result interface{} `json:"result"`
	Error  *string     `json:"error"`
}

// NewClient creates a new Anki Connect client
func NewClient(cfg *config.AnkiConfig) *Client {
	return &Client{
		config: cfg,
		httpClient: &http.Client{
			Timeout: DefaultTimeout,
		},
		retries:    DefaultRetries,
		retryDelay: DefaultRetryDelay,
	}
}

// NewClientWithHTTPClient creates a new Anki Connect client with a custom HTTP client
func NewClientWithHTTPClient(cfg *config.AnkiConfig, httpClient HTTPClient) *Client {
	return &Client{
		config:     cfg,
		httpClient: httpClient,
		retries:    DefaultRetries,
		retryDelay: DefaultRetryDelay,
	}
}

// SetRetryOptions sets the retry options for the client
func (c *Client) SetRetryOptions(retries int, retryDelay time.Duration) {
	c.retries = retries
	c.retryDelay = retryDelay
}

// Call makes a request to the Anki Connect API
func (c *Client) Call(action string, params interface{}) (interface{}, error) {
	req := Request{
		Action:  action,
		Version: APIVersion,
		Params:  params,
	}

	var result interface{}
	var lastErr error

	// Retry logic
	for attempt := 0; attempt <= c.retries; attempt++ {
		if attempt > 0 {
			// Wait before retrying
			time.Sleep(c.retryDelay)
		}

		result, lastErr = c.doRequest(req)
		if lastErr == nil {
			return result, nil
		}
	}

	return nil, fmt.Errorf("failed after %d attempts: %w", c.retries+1, lastErr)
}

// doRequest performs the actual HTTP request
func (c *Client) doRequest(req Request) (interface{}, error) {
	jsonData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	httpReq, err := http.NewRequest("POST", c.config.ConnectURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	httpReq.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var apiResp Response
	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	if apiResp.Error != nil {
		return nil, fmt.Errorf("API error: %s", *apiResp.Error)
	}

	return apiResp.Result, nil
}

// Ping checks if the Anki Connect API is available
func (c *Client) Ping() error {
	_, err := c.Call("version", nil)
	if err != nil {
		return fmt.Errorf("failed to connect to Anki: %w", err)
	}
	return nil
}
