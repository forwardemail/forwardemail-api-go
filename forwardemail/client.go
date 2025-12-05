package forwardemail

import (
	"fmt"
	"io"
	"net/http"
)

const (
	forwardemailAPIURL = "https://api.forwardemail.net"
)

// ClientOptions contains configuration options for creating a new Forward Email API client.
type ClientOptions struct {
	APIKey string
	APIURL string
}

// Client is the main client for interacting with the Forward Email API.
type Client struct {
	APIKey string
	APIURL string

	HTTPClient *http.Client
}

// NewClient returns a new Forward Email API Client.
func NewClient(options ClientOptions) *Client {
	apiURL := forwardemailAPIURL
	if options.APIURL != "" {
		apiURL = options.APIURL
	}

	return &Client{
		APIKey:     options.APIKey,
		APIURL:     apiURL,
		HTTPClient: http.DefaultClient,
	}
}

func (c *Client) newRequest(method, path string) (*http.Request, error) {
	req, err := http.NewRequest(method, c.APIURL+path, http.NoBody)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(c.APIKey, "")

	return req, nil
}

func (c *Client) doRequest(req *http.Request) ([]byte, error) {
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = res.Body.Close()
	}()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode == http.StatusOK || res.StatusCode == http.StatusNoContent {
		return body, err
	}

	return nil, fmt.Errorf("%w: %s", ErrRequestFailure, body)
}
