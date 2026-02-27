// Copyright Forward Email 2026
// SPDX-License-Identifier: MIT

package forwardemail

import (
	"errors"
	"fmt"
	"net/http"
)

var (
	// ErrRequestFailure is returned when an API request fails.
	ErrRequestFailure = errors.New("failed to complete request")
	// ErrMissingAPIKey is returned when no API key is provided.
	ErrMissingAPIKey = errors.New("no API key provided")
	// ErrNilContext is returned when a nil context is passed to an API method.
	ErrNilContext = errors.New("context cannot be nil")
	// ErrEmptyMethod is returned when an empty HTTP method is passed to newRequest.
	ErrEmptyMethod = errors.New("HTTP method cannot be empty")
	// ErrEmptyPath is returned when an empty request path is passed to newRequest.
	ErrEmptyPath = errors.New("request path cannot be empty")
	// ErrNilHTTPClient is returned when the client has a nil HTTPClient.
	ErrNilHTTPClient = errors.New("HTTP client cannot be nil")
	// ErrEmptyDomain is returned when a domain parameter is empty.
	ErrEmptyDomain = errors.New("domain cannot be empty")
	// ErrEmptyAlias is returned when an alias parameter is empty.
	ErrEmptyAlias = errors.New("alias cannot be empty")
	// ErrEmptyDomainName is returned when a domain name parameter is empty.
	ErrEmptyDomainName = errors.New("domain name cannot be empty")
	// ErrEmptyEmail is returned when an email parameter is empty.
	ErrEmptyEmail = errors.New("email cannot be empty")
	// ErrEmptyGroup is returned when a group parameter is empty.
	ErrEmptyGroup = errors.New("group cannot be empty")
)

// APIError represents an error response from the Forward Email API.
type APIError struct {
	StatusCode int
	Body       []byte
}

func (e *APIError) Error() string {
	return fmt.Sprintf("forwardemail: API error %d: %s", e.StatusCode, http.StatusText(e.StatusCode))
}
