// Copyright Forward Email 2026
// SPDX-License-Identifier: MIT

// Package forwardemail provides a client library for interacting with the Forward Email API.
package forwardemail

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

// Account represents a Forward Email account with plan details, user information, and metadata.
type Account struct {
	Plan           string    `json:"plan"`
	Email          string    `json:"email"`
	FullEmail      string    `json:"full_email"`
	DisplayName    string    `json:"display_name"`
	LastLocale     string    `json:"last_locale"`
	AddressCountry string    `json:"address_country"`
	ID             string    `json:"id"`
	Object         string    `json:"object"`
	Locale         string    `json:"locale"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	AddressHTML    string    `json:"address_html"`
}

// GetAccount retrieves the authenticated user's account information from the Forward Email API.
func (c *Client) GetAccount(ctx context.Context) (*Account, error) {
	if ctx == nil {
		return nil, ErrNilContext
	}
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	req, err := c.newRequest(ctx, "GET", "/v1/account", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request for GetAccount: %w", err)
	}

	res, err := c.doRequest(req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch account: %w", err)
	}

	var item Account
	if err := json.Unmarshal(res, &item); err != nil {
		return nil, fmt.Errorf("failed to parse account response: %w", err)
	}

	return &item, nil
}
