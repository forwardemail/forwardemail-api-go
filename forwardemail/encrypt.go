// Copyright Forward Email 2026
// SPDX-License-Identifier: MIT

package forwardemail

// Forward Email allows you to encrypt records even on the free plan at no cost. As highly requested in a Privacy Guides discussion and on our GitHub issues we've added this.

import (
	"context"
	"encoding/json"
	"strings"
)

// EncryptionResponse represents the API Response for encrypting a TXT record.
type EncryptionResponse struct {
	Encrypted string `json:"encrypted"`
}

// EncryptRecord allows you to encrypt a TXT record value so that you can safely add it to your DNS without exposing the value in plaintext.
func (c *Client) EncryptRecord(ctx context.Context, input string) (*EncryptionResponse, error) {
	if ctx == nil {
		return nil, ErrNilContext
	}
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	if input == "" {
		return nil, ErrMissingEncryptionInput
	}

	payload := strings.NewReader(`{"input": "` + input + `"}`)
	req, err := c.newRequest(ctx, "POST", "/v1/encrypt", payload)
	if err != nil {
		return nil, ErrFailedToCreateRequest
	}

	res, err := c.doRequest(req)
	if err != nil {
		return nil, ErrFailedToDoRequest
	}

	var encryptionResponse EncryptionResponse
	if err := json.Unmarshal(res, &encryptionResponse); err != nil {
		return nil, ErrFailedToUnmarshalResponse
	}

	return &encryptionResponse, nil
}
