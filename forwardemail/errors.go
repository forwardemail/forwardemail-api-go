// Copyright Forward Email 2026
// SPDX-License-Identifier: MIT

package forwardemail

import "errors"

var (
	// ErrRequestFailure is returned when an API request fails.
	ErrRequestFailure = errors.New("failed to complete request")
	// ErrMissingAPIKey is returned when no API key is provided.
	ErrMissingAPIKey = errors.New("no API key provided")
)
