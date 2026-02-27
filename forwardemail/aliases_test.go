// Copyright Forward Email 2026
// SPDX-License-Identifier: MIT

package forwardemail

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestClient_GetAlias(t *testing.T) {
	type request struct {
		domain string
		alias  string
	}

	tests := []struct {
		name     string
		req      request
		domain   string
		response string
		want     *Alias
	}{
		{
			name: "no data",
			req: request{
				domain: "stark.com",
				alias:  "tony",
			},
		},
		{
			name: "ok",
			req: request{
				domain: "stark.com",
				alias:  "tony",
			},
			response: `{
				"user": {
				  "email": "tony@stark.com",
				  "display_name": "tony@stark.com",
				  "id": "59ad551ae6fb4a4c53427ca38079f029"
				},
				"domain": {
				  "name": "stark.com",
				  "id": "15ff615b6180f1fc7faf40e6"
				},
				"name": "tony",
				"labels": [
				  "catch-all"
				],
				"description": "main email",
				"is_enabled": true,
				"has_recipient_verification": true,
				"recipients": [
				  "james@rhodes.com"
				],
				"id": "6525b03e0bde8f333ace5824",
				"object": "alias",
				"created_at": "2023-10-10T20:12:46.588Z",
				"updated_at": "2023-10-10T20:12:46.588Z"
			}`,
			want: &Alias{
				Account: Account{
					Email:       "tony@stark.com",
					DisplayName: "tony@stark.com",
					ID:          "59ad551ae6fb4a4c53427ca38079f029",
				},
				Domain: Domain{
					Name: "stark.com",
					ID:   "15ff615b6180f1fc7faf40e6",
				},
				Name:                     "tony",
				Labels:                   []string{"catch-all"},
				IsEnabled:                true,
				HasRecipientVerification: true,
				Recipients:               []string{"james@rhodes.com"},
				Description:              "main email",
				ID:                       "6525b03e0bde8f333ace5824",
				Object:                   "alias",
				CreatedAt:                parseTime("2023-10-10T20:12:46.588Z"),
				UpdatedAt:                parseTime("2023-10-10T20:12:46.588Z"),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
				fmt.Fprint(w, tt.response)
			}))
			defer svr.Close()

			c, _ := NewClient("test-key", WithAPIURL(svr.URL))

			got, _ := c.GetAlias(context.Background(), tt.req.domain, tt.req.alias)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Fatalf("values are not the same %s", diff)
			}
		})
	}
}

func TestClient_GetAliases(t *testing.T) {
	type request struct {
		domain string
	}

	tests := []struct {
		name string
		req  request
		res  string
		want []Alias
	}{
		{
			name: "no data",
			req: request{
				domain: "stark.com",
			},
		},
		{
			name: "ok",
			req: request{
				domain: "stark.com",
			},
			res: `[
				{
					"user": {
					  "email": "tony@stark.com",
					  "display_name": "tony@stark.com",
					  "id": "59ad551ae6fb4a4c53427ca38079f029"
					},
					"domain": {
					  "name": "stark.com",
					  "id": "15ff615b6180f1fc7faf40e6"
					},
					"name": "tony",
					"labels": [
					  "catch-all"
					],
					"is_enabled": true,
					"description": "main email",
					"has_recipient_verification": true,
					"recipients": [
					  "james@rhodes.com"
					],
					"id": "6525b03e0bde8f333ace5824",
					"object": "alias",
					"created_at": "2023-10-10T20:12:46.588Z",
					"updated_at": "2023-10-10T20:12:46.588Z"
				},
				{
					"user": {
					  "email": "tony@stark.com",
					  "display_name": "tony@stark.com",
					  "id": "59ad551ae6fb4a4c53427ca38079f029"
					},
					"domain": {
					  "name": "stark.com",
					  "id": "15ff615b6180f1fc7faf40e6"
					},
					"name": "james",
					"labels": [
					  "catch-all"
					],
					"is_enabled": true,
					"has_recipient_verification": true,
					"recipients": [
					  "james@rhodes.com"
					],
					"id": "b078f60f2636c4d6cf668d9b36a3e42e",
					"object": "alias",
					"created_at": "2023-10-12T18:11:22.123Z",
					"updated_at": "2023-10-12T19:55:56.534Z"
				}
			]`,
			want: []Alias{
				{
					Account: Account{
						Email:       "tony@stark.com",
						DisplayName: "tony@stark.com",
						ID:          "59ad551ae6fb4a4c53427ca38079f029",
					},
					Domain: Domain{
						Name: "stark.com",
						ID:   "15ff615b6180f1fc7faf40e6",
					},
					Name:                     "tony",
					Labels:                   []string{"catch-all"},
					IsEnabled:                true,
					HasRecipientVerification: true,
					Description:              "main email",
					Recipients:               []string{"james@rhodes.com"},
					ID:                       "6525b03e0bde8f333ace5824",
					Object:                   "alias",
					CreatedAt:                parseTime("2023-10-10T20:12:46.588Z"),
					UpdatedAt:                parseTime("2023-10-10T20:12:46.588Z"),
				},
				{
					Account: Account{
						Email:       "tony@stark.com",
						DisplayName: "tony@stark.com",
						ID:          "59ad551ae6fb4a4c53427ca38079f029",
					},
					Domain: Domain{
						Name: "stark.com",
						ID:   "15ff615b6180f1fc7faf40e6",
					},
					Name:                     "james",
					Labels:                   []string{"catch-all"},
					IsEnabled:                true,
					HasRecipientVerification: true,
					Recipients:               []string{"james@rhodes.com"},
					ID:                       "b078f60f2636c4d6cf668d9b36a3e42e",
					Object:                   "alias",
					CreatedAt:                parseTime("2023-10-12T18:11:22.123Z"),
					UpdatedAt:                parseTime("2023-10-12T19:55:56.534Z"),
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
				fmt.Fprint(w, tt.res)
			}))
			defer svr.Close()

			c, _ := NewClient("test-key", WithAPIURL(svr.URL))

			got, _ := c.GetAliases(context.Background(), tt.req.domain)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Fatalf("values are not the same %s", diff)
			}
		})
	}
}

func TestClient_CreateAlias(t *testing.T) {
	type request struct {
		domain string
		alias  string
		params AliasParameters
	}

	tests := []struct {
		name string
		req  request
		res  string
		want *Alias
	}{
		{
			name: "no data",
			req: request{
				domain: "stark.com",
				alias:  "tony",
			},
		},
		{
			name: "ok",
			req: request{
				domain: "stark.com",
				alias:  "*",
				params: AliasParameters{
					Recipients:               pointSliceOfStrings([]string{"james@rhodes.com"}),
					Labels:                   pointSliceOfStrings([]string{"catch-all"}),
					Description:              "main email",
					IsEnabled:                pointBool(true),
					HasRecipientVerification: pointBool(true),
				},
			},
			res: `{
				"user": {
				  "email": "tony@stark.com",
				  "display_name": "tony@stark.com",
				  "id": "59ad551ae6fb4a4c53427ca38079f029"
				},
				"domain": {
				  "name": "stark.com",
				  "id": "15ff615b6180f1fc7faf40e6"
				},
				"name": "*",
				"labels": [
				  "catch-all"
				],
				"description": "main email",
				"is_enabled": true,
				"has_recipient_verification": true,
				"recipients": [
				  "james@rhodes.com"
				],
				"id": "6525b03e0bde8f333ace5824",
				"object": "alias",
				"created_at": "2023-10-10T20:12:46.588Z",
				"updated_at": "2023-11-11T22:12:42.533Z"
			}`,
			want: &Alias{
				Account: Account{
					Email:       "tony@stark.com",
					DisplayName: "tony@stark.com",
					ID:          "59ad551ae6fb4a4c53427ca38079f029",
				},
				Domain: Domain{
					Name: "stark.com",
					ID:   "15ff615b6180f1fc7faf40e6",
				},
				Name:                     "*",
				Labels:                   []string{"catch-all"},
				IsEnabled:                true,
				HasRecipientVerification: true,
				Description:              "main email",
				Recipients:               []string{"james@rhodes.com"},
				ID:                       "6525b03e0bde8f333ace5824",
				Object:                   "alias",
				CreatedAt:                parseTime("2023-10-10T20:12:46.588Z"),
				UpdatedAt:                parseTime("2023-11-11T22:12:42.533Z"),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
				fmt.Fprint(w, tt.res)
			}))
			defer svr.Close()

			c, _ := NewClient("test-key", WithAPIURL(svr.URL))

			got, _ := c.CreateAlias(context.Background(), tt.req.domain, tt.req.alias, tt.req.params)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Fatalf("values are not the same %s", diff)
			}
		})
	}
}

func TestClient_UpdateAlias(t *testing.T) {
	type request struct {
		domain string
		alias  string
		params AliasParameters
	}

	tests := []struct {
		name string
		req  request
		res  string
		want *Alias
	}{
		{
			name: "no data",
			req: request{
				domain: "stark.com",
				alias:  "tony",
			},
		},
		{
			name: "ok",
			req: request{
				domain: "stark.com",
				alias:  "james",
				params: AliasParameters{
					Recipients:               pointSliceOfStrings([]string{"james@rhodes.com"}),
					Labels:                   pointSliceOfStrings([]string{"catch-all", "friends"}),
					Description:              "main email",
					IsEnabled:                pointBool(true),
					HasRecipientVerification: pointBool(true),
				},
			},
			res: `{
				"user": {
				  "email": "tony@stark.com",
				  "display_name": "tony@stark.com",
				  "id": "59ad551ae6fb4a4c53427ca38079f029"
				},
				"domain": {
				  "name": "stark.com",
				  "id": "15ff615b6180f1fc7faf40e6"
				},
				"name": "james",
				"labels": [
				  "catch-all",
				  "friends"
				],
				"description": "main email",
				"is_enabled": true,
				"has_recipient_verification": true,
				"recipients": [
				  "james@rhodes.com"
				],
				"id": "6525b03e0bde8f333ace5824",
				"object": "alias",
				"created_at": "2023-10-10T20:12:46.588Z",
				"updated_at": "2023-11-11T22:12:42.533Z"
			}`,
			want: &Alias{
				Account: Account{
					Email:       "tony@stark.com",
					DisplayName: "tony@stark.com",
					ID:          "59ad551ae6fb4a4c53427ca38079f029",
				},
				Domain: Domain{
					Name: "stark.com",
					ID:   "15ff615b6180f1fc7faf40e6",
				},
				Name:                     "james",
				Labels:                   []string{"catch-all", "friends"},
				Description:              "main email",
				IsEnabled:                true,
				HasRecipientVerification: true,
				Recipients:               []string{"james@rhodes.com"},
				ID:                       "6525b03e0bde8f333ace5824",
				Object:                   "alias",
				CreatedAt:                parseTime("2023-10-10T20:12:46.588Z"),
				UpdatedAt:                parseTime("2023-11-11T22:12:42.533Z"),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
				fmt.Fprint(w, tt.res)
			}))
			defer svr.Close()

			c, _ := NewClient("test-key", WithAPIURL(svr.URL))

			got, _ := c.UpdateAlias(context.Background(), tt.req.domain, tt.req.alias, tt.req.params)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Fatalf("values are not the same %s", diff)
			}
		})
	}
}

func TestClient_DeleteAlias(t *testing.T) {
	type request struct {
		domain string
		alias  string
	}
	type response struct {
		code int
		body string
	}

	tests := []struct {
		name      string
		req       request
		res       response
		wantError bool
	}{
		{
			name: "ok",
			req: request{
				domain: "stark.com",
				alias:  "tony",
			},
			res: response{
				code: http.StatusNoContent,
			},
		},
		{
			name: "not ok",
			req: request{
				domain: "stark.com",
				alias:  "tony",
			},
			res: response{
				code: http.StatusInternalServerError,
				body: "oh no",
			},
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
				w.WriteHeader(tt.res.code)
				_, _ = w.Write([]byte(tt.res.body))
			}))
			defer svr.Close()

			c, _ := NewClient("test-key", WithAPIURL(svr.URL))

			got := c.DeleteAlias(context.Background(), tt.req.domain, tt.req.alias)
			if tt.wantError {
				if got == nil {
					t.Fatal("expected error, got nil")
				}
				var apiErr *APIError
				if !errors.As(got, &apiErr) {
					t.Fatalf("expected error to wrap *APIError, got %v", got)
				}
			} else if got != nil {
				t.Fatalf("expected no error, got %v", got)
			}
		})
	}
}

func pointSliceOfStrings(s []string) *[]string {
	return &s
}
