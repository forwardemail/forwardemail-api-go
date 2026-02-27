// Copyright Forward Email 2026
// SPDX-License-Identifier: MIT

package forwardemail

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
	"time"
)

// Invite represents a domain invitation for a user to join with a specific access group.
type Invite struct {
	Email     string    `json:"email"`
	Group     string    `json:"group"`
	ID        string    `json:"id"`
	Object    string    `json:"object"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// CreateDomainInvite sends an invitation to a user to join a domain with the specified group permissions.
func (c *Client) CreateDomainInvite(ctx context.Context, domain, email, group string) (*Invite, error) {
	if ctx == nil {
		return nil, ErrNilContext
	}
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	if strings.TrimSpace(domain) == "" {
		return nil, ErrEmptyDomain
	}
	if strings.TrimSpace(email) == "" {
		return nil, ErrEmptyEmail
	}
	if strings.TrimSpace(group) == "" {
		return nil, ErrEmptyGroup
	}

	encodedDomain := url.PathEscape(domain)

	params := url.Values{}
	params.Add("email", email)
	params.Add("group", group)

	req, err := c.newRequest(ctx, "POST", fmt.Sprintf("/v1/domains/%s/invites", encodedDomain), strings.NewReader(params.Encode()))
	if err != nil {
		return nil, fmt.Errorf("failed to create request for CreateDomainInvite: %w", err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	res, err := c.doRequest(req)
	if err != nil {
		return nil, fmt.Errorf("failed to create domain invite: %w", err)
	}

	var item Invite
	if err := json.Unmarshal(res, &item); err != nil {
		return nil, fmt.Errorf("failed to parse domain invite response: %w", err)
	}

	return &item, nil
}
