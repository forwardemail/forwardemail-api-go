// Copyright Forward Email 2026
// SPDX-License-Identifier: MIT

package forwardemail

import (
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNewClient(t *testing.T) {
	tests := []struct {
		name    string
		options ClientOptions
		want    *Client
	}{
		{
			name:    "empty options",
			options: ClientOptions{},
			want: &Client{
				APIURL:     "https://api.forwardemail.net",
				HTTPClient: &http.Client{},
			},
		},
		{
			name: "with api key",
			options: ClientOptions{
				APIKey: "4e4d6c332b6fe62a63afe56171fd3725",
			},
			want: &Client{
				APIKey:     "4e4d6c332b6fe62a63afe56171fd3725",
				APIURL:     "https://api.forwardemail.net",
				HTTPClient: &http.Client{},
			},
		},
		{
			name: "with api url",
			options: ClientOptions{
				APIURL: "https://google.com",
			},
			want: &Client{
				APIURL:     "https://google.com",
				HTTPClient: &http.Client{},
			},
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
				HTTPClient: &http.Client{},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := NewClient(tt.options)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Fatalf("values are not the same %s", diff)
			}
		})
	}
}
