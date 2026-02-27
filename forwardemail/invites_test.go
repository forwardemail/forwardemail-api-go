// Copyright Forward Email 2026
// SPDX-License-Identifier: MIT

package forwardemail

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestClient_CreateDomainInvite(t *testing.T) {
	tests := []struct {
		name     string
		domain   string
		email    string
		group    string
		response string
		want     *Invite
	}{
		{
			name:   "no data",
			domain: "stark.com",
			email:  "tony@stark.com",
			group:  "admin",
		},
		{
			name:   "ok",
			domain: "stark.com",
			email:  "tony@stark.com",
			group:  "admin",
			response: `{
				"email": "tony@stark.com",
				"group": "admin",
				"id": "12345",
				"object": "invite",
				"created_at": "2024-12-10T18:14:31.378Z",
				"updated_at": "2024-12-10T18:14:31.378Z"
			}`,
			want: &Invite{
				Email:     "tony@stark.com",
				Group:     "admin",
				ID:        "12345",
				Object:    "invite",
				CreatedAt: parseTime("2024-12-10T18:14:31.378Z"),
				UpdatedAt: parseTime("2024-12-10T18:14:31.378Z"),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.Method != http.MethodPost {
					t.Errorf("expected POST method, got %s", r.Method)
				}
				if r.URL.Path != fmt.Sprintf("/v1/domains/%s/invites", tt.domain) {
					t.Errorf("expected URL %s, got %s", fmt.Sprintf("/v1/domains/%s/invites", tt.domain), r.URL.Path)
				}
				fmt.Fprint(w, tt.response)
			}))
			defer svr.Close()

			c, _ := NewClient("test-key", WithAPIURL(svr.URL))

			got, _ := c.CreateDomainInvite(context.Background(), tt.domain, tt.email, tt.group)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Fatalf("values are not the same %s", diff)
			}
		})
	}
}
