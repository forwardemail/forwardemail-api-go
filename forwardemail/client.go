package forwardemail

import (
	"fmt"
	"io"
	"net/http"
)

const (
	forwardemailApiUrl = "https://api.forwardemail.net"
)

type ClientOptions struct {
	ApiKey string
	ApiUrl string
}

type Client struct {
	ApiKey string
	ApiUrl string

	HttpClient *http.Client
}

// NewClient returns a new Forward Email API Client.
func NewClient(options ClientOptions) *Client {
	apiUrl := forwardemailApiUrl
	if options.ApiUrl != "" {
		apiUrl = options.ApiUrl
	}

	return &Client{
		ApiKey:     options.ApiKey,
		ApiUrl:     apiUrl,
		HttpClient: http.DefaultClient,
	}
}

func (c *Client) newRequest(method, path string) (*http.Request, error) {
	req, err := http.NewRequest(method, c.ApiUrl+path, nil)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(c.ApiKey, "")

	return req, nil
}

func (c *Client) doRequest(req *http.Request) ([]byte, error) {
	res, err := c.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode == http.StatusOK || res.StatusCode == http.StatusNoContent {
		return body, err
	}

	return nil, fmt.Errorf("%w: %s", ErrRequestFailure, body)
}
