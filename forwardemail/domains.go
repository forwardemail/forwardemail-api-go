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

// Domain represents a domain configuration with security settings, verification status, and metadata.
type Domain struct {
	HasAdultContentProtection bool      `json:"has_adult_content_protection"`
	HasPhishingProtection     bool      `json:"has_phishing_protection"`
	HasExecutableProtection   bool      `json:"has_executable_protection"`
	HasVirusProtection        bool      `json:"has_virus_protection"`
	IsCatchallRegexDisabled   bool      `json:"is_catchall_regex_disabled"`
	Plan                      string    `json:"plan"`
	MaxRecipientsPerAlias     int       `json:"max_recipients_per_alias"`
	SMTPPort                  string    `json:"smtp_port"`
	Name                      string    `json:"name"`
	HasMxRecord               bool      `json:"has_mx_record"`
	HasTxtRecord              bool      `json:"has_txt_record"`
	HasRecipientVerification  bool      `json:"has_recipient_verification"`
	HasCustomVerification     bool      `json:"has_custom_verification"`
	VerificationRecord        string    `json:"verification_record"`
	ID                        string    `json:"id"`
	Object                    string    `json:"object"`
	CreatedAt                 time.Time `json:"created_at"`
	UpdatedAt                 time.Time `json:"updated_at"`
	Link                      string    `json:"link"`
}

// DomainParameters contains optional parameters for creating or updating a domain.
type DomainParameters struct {
	HasAdultContentProtection *bool
	HasPhishingProtection     *bool
	HasExecutableProtection   *bool
	HasVirusProtection        *bool
	HasRecipientVerification  *bool
}

// GetDomains retrieves all domains associated with the authenticated account.
func (c *Client) GetDomains(ctx context.Context) ([]Domain, error) {
	if ctx == nil {
		return nil, ErrNilContext
	}
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	req, err := c.newRequest(ctx, "GET", "/v1/domains", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request for GetDomains: %w", err)
	}

	res, err := c.doRequest(req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch domains: %w", err)
	}

	var items []Domain
	if err := json.Unmarshal(res, &items); err != nil {
		return nil, fmt.Errorf("failed to parse domains response: %w", err)
	}

	return items, nil
}

// GetDomain retrieves a specific domain by name.
func (c *Client) GetDomain(ctx context.Context, name string) (*Domain, error) {
	if ctx == nil {
		return nil, ErrNilContext
	}
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	if strings.TrimSpace(name) == "" {
		return nil, ErrEmptyDomainName
	}

	encodedName := url.PathEscape(name)

	req, err := c.newRequest(ctx, "GET", fmt.Sprintf("/v1/domains/%s", encodedName), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request for GetDomain: %w", err)
	}

	res, err := c.doRequest(req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch domain: %w", err)
	}

	var item Domain
	if err := json.Unmarshal(res, &item); err != nil {
		return nil, fmt.Errorf("failed to parse domain response: %w", err)
	}

	return &item, nil
}

// CreateDomain adds a new domain to the account with the specified configuration parameters.
func (c *Client) CreateDomain(ctx context.Context, name string, parameters DomainParameters) (*Domain, error) {
	if ctx == nil {
		return nil, ErrNilContext
	}
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	if strings.TrimSpace(name) == "" {
		return nil, ErrEmptyDomainName
	}

	params := url.Values{}
	params.Add("domain", name)

	for k, v := range map[string]*bool{
		"has_adult_content_protection": parameters.HasAdultContentProtection,
		"has_phishing_protection":      parameters.HasPhishingProtection,
		"has_executable_protection":    parameters.HasExecutableProtection,
		"has_virus_protection":         parameters.HasVirusProtection,
		"has_recipient_verification":   parameters.HasRecipientVerification,
	} {
		if v != nil {
			params.Add(k, strconv.FormatBool(*v))
		}
	}

	req, err := c.newRequest(ctx, "POST", "/v1/domains", strings.NewReader(params.Encode()))
	if err != nil {
		return nil, fmt.Errorf("failed to create request for CreateDomain: %w", err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	res, err := c.doRequest(req)
	if err != nil {
		return nil, fmt.Errorf("failed to create domain: %w", err)
	}

	var item Domain
	if err := json.Unmarshal(res, &item); err != nil {
		return nil, fmt.Errorf("failed to parse create domain response: %w", err)
	}

	return &item, nil
}

// UpdateDomain modifies an existing domain's configuration parameters.
func (c *Client) UpdateDomain(ctx context.Context, name string, parameters DomainParameters) (*Domain, error) {
	if ctx == nil {
		return nil, ErrNilContext
	}
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	if strings.TrimSpace(name) == "" {
		return nil, ErrEmptyDomainName
	}

	encodedName := url.PathEscape(name)

	params := url.Values{}
	params.Add("domain", name)

	for k, v := range map[string]*bool{
		"has_adult_content_protection": parameters.HasAdultContentProtection,
		"has_phishing_protection":      parameters.HasPhishingProtection,
		"has_executable_protection":    parameters.HasExecutableProtection,
		"has_virus_protection":         parameters.HasVirusProtection,
		"has_recipient_verification":   parameters.HasRecipientVerification,
	} {
		if v != nil {
			params.Add(k, strconv.FormatBool(*v))
		}
	}

	req, err := c.newRequest(ctx, "PUT", fmt.Sprintf("/v1/domains/%s", encodedName), strings.NewReader(params.Encode()))
	if err != nil {
		return nil, fmt.Errorf("failed to create request for UpdateDomain: %w", err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	res, err := c.doRequest(req)
	if err != nil {
		return nil, fmt.Errorf("failed to update domain: %w", err)
	}

	var item Domain
	if err := json.Unmarshal(res, &item); err != nil {
		return nil, fmt.Errorf("failed to parse update domain response: %w", err)
	}

	return &item, nil
}

// DeleteDomain removes a domain from the account.
func (c *Client) DeleteDomain(ctx context.Context, name string) error {
	if ctx == nil {
		return ErrNilContext
	}
	if err := ctx.Err(); err != nil {
		return err
	}
	if strings.TrimSpace(name) == "" {
		return ErrEmptyDomainName
	}

	encodedName := url.PathEscape(name)

	req, err := c.newRequest(ctx, "DELETE", fmt.Sprintf("/v1/domains/%s", encodedName), nil)
	if err != nil {
		return fmt.Errorf("failed to create request for DeleteDomain: %w", err)
	}

	_, err = c.doRequest(req)
	if err != nil {
		return fmt.Errorf("failed to delete domain: %w", err)
	}

	return nil
}
