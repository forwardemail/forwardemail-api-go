// Copyright Forward Email 2026
// SPDX-License-Identifier: MIT

package forwardemail

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"time"
)

// Alias represents an email alias configuration including recipients, labels, and verification settings.
type Alias struct {
	Account                  Account   `json:"user"`
	Domain                   Domain    `json:"domain"`
	Name                     string    `json:"name"`
	Labels                   []string  `json:"labels"`
	Description              string    `json:"description"`
	IsEnabled                bool      `json:"is_enabled"`
	HasRecipientVerification bool      `json:"has_recipient_verification"`
	Recipients               []string  `json:"recipients"`
	ID                       string    `json:"id"`
	Object                   string    `json:"object"`
	CreatedAt                time.Time `json:"created_at"`
	UpdatedAt                time.Time `json:"updated_at"`
}

// AliasParameters contains optional parameters for creating or updating an alias.
type AliasParameters struct {
	Recipients               *[]string
	Labels                   *[]string
	Description              string `json:"description"`
	HasRecipientVerification *bool
	IsEnabled                *bool
}

// GetAliases retrieves all email aliases for the specified domain.
func (c *Client) GetAliases(ctx context.Context, domain string) ([]Alias, error) {
	if ctx == nil {
		return nil, ErrNilContext
	}
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	if strings.TrimSpace(domain) == "" {
		return nil, ErrEmptyDomain
	}

	encodedDomain := url.PathEscape(domain)

	req, err := c.newRequest(ctx, "GET", fmt.Sprintf("/v1/domains/%s/aliases", encodedDomain), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request for GetAliases: %w", err)
	}

	res, err := c.doRequest(req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch aliases: %w", err)
	}

	var items []Alias
	if err := json.Unmarshal(res, &items); err != nil {
		return nil, fmt.Errorf("failed to parse aliases response: %w", err)
	}

	return items, nil
}

// GetAlias retrieves a specific email alias by name for the specified domain.
func (c *Client) GetAlias(ctx context.Context, domain, alias string) (*Alias, error) {
	if ctx == nil {
		return nil, ErrNilContext
	}
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	if strings.TrimSpace(domain) == "" {
		return nil, ErrEmptyDomain
	}
	if strings.TrimSpace(alias) == "" {
		return nil, ErrEmptyAlias
	}

	encodedDomain := url.PathEscape(domain)
	encodedAlias := url.PathEscape(alias)

	req, err := c.newRequest(ctx, "GET", fmt.Sprintf("/v1/domains/%s/aliases/%s", encodedDomain, encodedAlias), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request for GetAlias: %w", err)
	}

	res, err := c.doRequest(req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch alias: %w", err)
	}

	var item Alias
	if err := json.Unmarshal(res, &item); err != nil {
		return nil, fmt.Errorf("failed to parse alias response: %w", err)
	}

	return &item, nil
}

// CreateAlias creates a new email alias for the specified domain with the given parameters.
func (c *Client) CreateAlias(ctx context.Context, domain, alias string, parameters AliasParameters) (*Alias, error) {
	if ctx == nil {
		return nil, ErrNilContext
	}
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	if strings.TrimSpace(domain) == "" {
		return nil, ErrEmptyDomain
	}
	if strings.TrimSpace(alias) == "" {
		return nil, ErrEmptyAlias
	}

	encodedDomain := url.PathEscape(domain)

	params := url.Values{}
	params.Add("name", alias)
	encodeAliasParams(params, parameters)

	req, err := c.newRequest(ctx, "POST", fmt.Sprintf("/v1/domains/%s/aliases", encodedDomain), strings.NewReader(params.Encode()))
	if err != nil {
		return nil, fmt.Errorf("failed to create request for CreateAlias: %w", err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	res, err := c.doRequest(req)
	if err != nil {
		return nil, fmt.Errorf("failed to create alias: %w", err)
	}

	var item Alias
	if err := json.Unmarshal(res, &item); err != nil {
		return nil, fmt.Errorf("failed to parse create alias response: %w", err)
	}

	return &item, nil
}

// UpdateAlias updates an existing email alias with new parameters for the specified domain.
func (c *Client) UpdateAlias(ctx context.Context, domain, alias string, parameters AliasParameters) (*Alias, error) {
	if ctx == nil {
		return nil, ErrNilContext
	}
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	if strings.TrimSpace(domain) == "" {
		return nil, ErrEmptyDomain
	}
	if strings.TrimSpace(alias) == "" {
		return nil, ErrEmptyAlias
	}

	encodedDomain := url.PathEscape(domain)
	encodedAlias := url.PathEscape(alias)

	params := url.Values{}
	params.Add("name", alias)
	encodeAliasParams(params, parameters)

	req, err := c.newRequest(ctx, "PUT", fmt.Sprintf("/v1/domains/%s/aliases/%s", encodedDomain, encodedAlias), strings.NewReader(params.Encode()))
	if err != nil {
		return nil, fmt.Errorf("failed to create request for UpdateAlias: %w", err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	res, err := c.doRequest(req)
	if err != nil {
		return nil, fmt.Errorf("failed to update alias: %w", err)
	}

	var item Alias
	if err := json.Unmarshal(res, &item); err != nil {
		return nil, fmt.Errorf("failed to parse update alias response: %w", err)
	}

	return &item, nil
}

func encodeAliasParams(params url.Values, parameters AliasParameters) {
	if parameters.Description != "" {
		params.Add("description", parameters.Description)
	}

	for k, v := range map[string]*bool{
		"has_recipient_verification": parameters.HasRecipientVerification,
		"is_enabled":                 parameters.IsEnabled,
	} {
		if v != nil {
			params.Add(k, strconv.FormatBool(*v))
		}
	}

	for k, v := range map[string]*[]string{
		"recipients[]": parameters.Recipients,
		"labels[]":     parameters.Labels,
	} {
		if v != nil {
			for _, vv := range *v {
				params.Add(k, vv)
			}
		}
	}
}

// DeleteAlias removes an email alias from the specified domain.
func (c *Client) DeleteAlias(ctx context.Context, domain, alias string) error {
	if ctx == nil {
		return ErrNilContext
	}
	if err := ctx.Err(); err != nil {
		return err
	}
	if strings.TrimSpace(domain) == "" {
		return ErrEmptyDomain
	}
	if strings.TrimSpace(alias) == "" {
		return ErrEmptyAlias
	}

	encodedDomain := url.PathEscape(domain)
	encodedAlias := url.PathEscape(alias)

	req, err := c.newRequest(ctx, "DELETE", fmt.Sprintf("/v1/domains/%s/aliases/%s", encodedDomain, encodedAlias), nil)
	if err != nil {
		return fmt.Errorf("failed to create request for DeleteAlias: %w", err)
	}

	_, err = c.doRequest(req)
	if err != nil {
		return fmt.Errorf("failed to delete alias: %w", err)
	}

	return nil
}
