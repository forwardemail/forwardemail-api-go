// Copyright Forward Email 2026
// SPDX-License-Identifier: MIT

package forwardemail

import (
	"errors"
	"net/http"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

func TestNewClient(t *testing.T) {
	tests := []struct {
		name      string
		options   ClientOptions
		want      *Client
		wantError error
	}{
		{
			name:      "empty options returns error",
			options:   ClientOptions{},
			want:      nil,
			wantError: ErrMissingAPIKey,
		},
		{
			name: "with api key",
			options: ClientOptions{
				APIKey: "4e4d6c332b6fe62a63afe56171fd3725",
			},
			want: &Client{
				APIKey:     "4e4d6c332b6fe62a63afe56171fd3725",
				APIURL:     "https://api.forwardemail.net",
				HTTPClient: &http.Client{Timeout: 30 * time.Second},
			},
		},
		{
			name: "with api url but no api key returns error",
			options: ClientOptions{
				APIURL: "https://google.com",
			},
			want:      nil,
			wantError: ErrMissingAPIKey,
		},
		{
			name: "with everything at once",
			options: ClientOptions{
				APIKey: "4e4d6c332b6fe62a63afe56171fd3725",
				APIURL: "https://google.com",
			},
			want: &Client{
				APIKey:     "4e4d6c332b6fe62a63afe56171fd3725",
				APIURL:     "https://google.com",
				HTTPClient: &http.Client{Timeout: 30 * time.Second},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewClient(tt.options)
			if tt.wantError != nil {
				if !errors.Is(err, tt.wantError) {
					t.Fatalf("expected error %v, got %v", tt.wantError, err)
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Fatalf("values are not the same %s", diff)
			}
		})
	}
}
