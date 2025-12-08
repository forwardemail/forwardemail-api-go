// Package forwardemail provides a client library for interacting with the Forward Email API.
package forwardemail

import (
	"encoding/json"
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
func (c *Client) GetAccount() (*Account, error) {
	req, err := c.newRequest("GET", "/v1/account")
	if err != nil {
		return nil, err
	}

	res, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var item Account

	err = json.Unmarshal(res, &item)
	if err != nil {
		return nil, err
	}

	return &item, nil
}
