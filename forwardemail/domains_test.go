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

func TestClient_GetDomain(t *testing.T) {
	tests := []struct {
		name     string
		domain   string
		response string
		want     *Domain
	}{
		{
			name:   "no data",
			domain: "stark.com",
		},
		{
			name:   "ok",
			domain: "stark.com",
			response: `{
				  "has_adult_content_protection": true,
				  "has_phishing_protection": true,
				  "has_executable_protection": true,
				  "has_virus_protection": true,
				  "is_catchall_regex_disabled": false,
				  "plan": "enhanced_protection",
				  "max_recipients_per_alias": 10,
				  "smtp_port": "25",
				  "name": "stark.com",
				  "has_mx_record": true,
				  "has_txt_record": true,
				  "has_recipient_verification": false,
				  "has_custom_verification": false,
				  "verification_record": "v8O0S8JjRv",
				  "id": "15ff615b6180f1fc7faf40e6",
				  "object": "domain",
				  "created_at": "2023-09-21T20:18:24.790Z",
				  "updated_at": "2023-10-07T21:21:01.992Z",
				  "last_allowlist_sync_at": "2023-10-07T13:06:02.630Z",
				  "link": "https://forwardemail.net/my-account/domains/stark.com"
			}`,
			want: &Domain{
				HasAdultContentProtection: true,
				HasPhishingProtection:     true,
				HasExecutableProtection:   true,
				HasVirusProtection:        true,
				Plan:                      "enhanced_protection",
				MaxRecipientsPerAlias:     10,
				SMTPPort:                  "25",
				Name:                      "stark.com",
				HasMxRecord:               true,
				HasTxtRecord:              true,
				VerificationRecord:        "v8O0S8JjRv",
				ID:                        "15ff615b6180f1fc7faf40e6",
				Object:                    "domain",
				CreatedAt:                 parseTime("2023-09-21T20:18:24.790Z"),
				UpdatedAt:                 parseTime("2023-10-07T21:21:01.992Z"),
				Link:                      "https://forwardemail.net/my-account/domains/stark.com",
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

			got, _ := c.GetDomain(context.Background(), tt.domain)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Fatalf("values are not the same %s", diff)
			}
		})
	}
}

func TestClient_GetDomains(t *testing.T) {
	tests := []struct {
		name     string
		response string
		want     []Domain
	}{
		{
			name: "no data",
		},
		{
			name: "ok",
			response: `[
				{
					"has_adult_content_protection": true,
					"has_phishing_protection": true,
					"has_executable_protection": true,
					"has_virus_protection": true,
					"is_catchall_regex_disabled": false,
					"plan": "enhanced_protection",
					"max_recipients_per_alias": 10,
					"smtp_port": "25",
					"name": "stark.com",
					"has_mx_record": true,
					"has_txt_record": true,
					"has_recipient_verification": false,
					"has_custom_verification": false,
					"verification_record": "v8O0S8JjRv",
					"id": "15ff615b6180f1fc7faf40e6",
					"object": "domain",
					"created_at": "2023-09-21T20:18:24.790Z",
					"updated_at": "2023-10-07T21:21:01.992Z",
					"last_allowlist_sync_at": "2023-10-07T13:06:02.630Z",
					"link": "https://forwardemail.net/my-account/domains/stark.com"
				},
				{
					"has_adult_content_protection": true,
					"has_phishing_protection": true,
					"has_executable_protection": true,
					"has_virus_protection": true,
					"is_catchall_regex_disabled": false,
					"plan": "enhanced_protection",
					"max_recipients_per_alias": 10,
					"smtp_port": "25",
					"name": "rhodes.com",
					"has_mx_record": true,
					"has_txt_record": true,
					"has_recipient_verification": false,
					"has_custom_verification": false,
					"verification_record": "v0jJ88SROv",
					"id": "e61ffff601c7fb14185af506",
					"object": "domain",
					"created_at": "2023-04-04T12:13:55.723Z",
					"updated_at": "2023-11-03T22:22:02.724Z",
					"last_allowlist_sync_at": "2023-11-03T22:23:08.123Z",
					"link": "https://forwardemail.net/my-account/domains/rhodes.com"
				}
			]`,
			want: []Domain{
				{
					HasAdultContentProtection: true,
					HasPhishingProtection:     true,
					HasExecutableProtection:   true,
					HasVirusProtection:        true,
					Plan:                      "enhanced_protection",
					MaxRecipientsPerAlias:     10,
					SMTPPort:                  "25",
					Name:                      "stark.com",
					HasMxRecord:               true,
					HasTxtRecord:              true,
					VerificationRecord:        "v8O0S8JjRv",
					ID:                        "15ff615b6180f1fc7faf40e6",
					Object:                    "domain",
					CreatedAt:                 parseTime("2023-09-21T20:18:24.790Z"),
					UpdatedAt:                 parseTime("2023-10-07T21:21:01.992Z"),
					Link:                      "https://forwardemail.net/my-account/domains/stark.com",
				},
				{
					HasAdultContentProtection: true,
					HasPhishingProtection:     true,
					HasExecutableProtection:   true,
					HasVirusProtection:        true,
					Plan:                      "enhanced_protection",
					MaxRecipientsPerAlias:     10,
					SMTPPort:                  "25",
					Name:                      "rhodes.com",
					HasMxRecord:               true,
					HasTxtRecord:              true,
					VerificationRecord:        "v0jJ88SROv",
					ID:                        "e61ffff601c7fb14185af506",
					Object:                    "domain",
					CreatedAt:                 parseTime("2023-04-04T12:13:55.723Z"),
					UpdatedAt:                 parseTime("2023-11-03T22:22:02.724Z"),
					Link:                      "https://forwardemail.net/my-account/domains/rhodes.com",
				},
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

			got, _ := c.GetDomains(context.Background())
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Fatalf("values are not the same %s", diff)
			}
		})
	}
}

func TestClient_CreateDomain(t *testing.T) {
	tests := []struct {
		name       string
		domain     string
		parameters DomainParameters
		response   string
		want       *Domain
	}{
		{
			name:   "no data",
			domain: "stark.com",
		},
		{
			name:   "ok",
			domain: "stark.com",
			parameters: DomainParameters{
				HasAdultContentProtection: pointBool(true),
				HasPhishingProtection:     pointBool(true),
				HasExecutableProtection:   pointBool(true),
				HasVirusProtection:        pointBool(true),
				HasRecipientVerification:  pointBool(true),
			},
			response: `{
				  "has_adult_content_protection": true,
				  "has_phishing_protection": true,
				  "has_executable_protection": true,
				  "has_virus_protection": true,
				  "is_catchall_regex_disabled": false,
				  "plan": "enhanced_protection",
				  "max_recipients_per_alias": 10,
				  "smtp_port": "25",
				  "name": "stark.com",
				  "has_mx_record": true,
				  "has_txt_record": true,
				  "has_recipient_verification": false,
				  "has_custom_verification": false,
				  "verification_record": "v8O0S8JjRv",
				  "id": "15ff615b6180f1fc7faf40e6",
				  "object": "domain",
				  "created_at": "2023-09-21T20:18:24.790Z",
				  "updated_at": "2023-10-07T21:21:01.992Z",
				  "last_allowlist_sync_at": "2023-10-07T13:06:02.630Z",
				  "link": "https://forwardemail.net/my-account/domains/stark.com"
			}`,
			want: &Domain{
				HasAdultContentProtection: true,
				HasPhishingProtection:     true,
				HasExecutableProtection:   true,
				HasVirusProtection:        true,
				Plan:                      "enhanced_protection",
				MaxRecipientsPerAlias:     10,
				SMTPPort:                  "25",
				Name:                      "stark.com",
				HasMxRecord:               true,
				HasTxtRecord:              true,
				VerificationRecord:        "v8O0S8JjRv",
				ID:                        "15ff615b6180f1fc7faf40e6",
				Object:                    "domain",
				CreatedAt:                 parseTime("2023-09-21T20:18:24.790Z"),
				UpdatedAt:                 parseTime("2023-10-07T21:21:01.992Z"),
				Link:                      "https://forwardemail.net/my-account/domains/stark.com",
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

			got, _ := c.CreateDomain(context.Background(), tt.domain, tt.parameters)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Fatalf("values are not the same %s", diff)
			}
		})
	}
}

func TestClient_UpdateDomain(t *testing.T) {
	tests := []struct {
		name       string
		domain     string
		parameters DomainParameters
		response   string
		want       *Domain
	}{
		{
			name:   "no data",
			domain: "stark.com",
		},
		{
			name:   "ok",
			domain: "stark.com",
			parameters: DomainParameters{
				HasAdultContentProtection: pointBool(true),
				HasPhishingProtection:     pointBool(true),
				HasExecutableProtection:   pointBool(true),
				HasVirusProtection:        pointBool(true),
				HasRecipientVerification:  pointBool(true),
			},
			response: `{
				  "has_adult_content_protection": true,
				  "has_phishing_protection": true,
				  "has_executable_protection": true,
				  "has_virus_protection": true,
				  "is_catchall_regex_disabled": false,
				  "plan": "enhanced_protection",
				  "max_recipients_per_alias": 10,
				  "smtp_port": "25",
				  "name": "stark.com",
				  "has_mx_record": true,
				  "has_txt_record": true,
				  "has_recipient_verification": false,
				  "has_custom_verification": false,
				  "verification_record": "v8O0S8JjRv",
				  "id": "15ff615b6180f1fc7faf40e6",
				  "object": "domain",
				  "created_at": "2023-09-21T20:18:24.790Z",
				  "updated_at": "2023-10-07T21:21:01.992Z",
				  "last_allowlist_sync_at": "2023-10-07T13:06:02.630Z",
				  "link": "https://forwardemail.net/my-account/domains/stark.com"
			}`,
			want: &Domain{
				HasAdultContentProtection: true,
				HasPhishingProtection:     true,
				HasExecutableProtection:   true,
				HasVirusProtection:        true,
				Plan:                      "enhanced_protection",
				MaxRecipientsPerAlias:     10,
				SMTPPort:                  "25",
				Name:                      "stark.com",
				HasMxRecord:               true,
				HasTxtRecord:              true,
				VerificationRecord:        "v8O0S8JjRv",
				ID:                        "15ff615b6180f1fc7faf40e6",
				Object:                    "domain",
				CreatedAt:                 parseTime("2023-09-21T20:18:24.790Z"),
				UpdatedAt:                 parseTime("2023-10-07T21:21:01.992Z"),
				Link:                      "https://forwardemail.net/my-account/domains/stark.com",
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

			got, _ := c.UpdateDomain(context.Background(), tt.domain, tt.parameters)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Fatalf("values are not the same %s", diff)
			}
		})
	}
}

func TestClient_DeleteDomain(t *testing.T) {
	type response struct {
		code int
		body string
	}

	tests := []struct {
		name      string
		domain    string
		resp      response
		wantError bool
	}{
		{
			name:   "ok",
			domain: "stark.com",
			resp: response{
				code: http.StatusNoContent,
			},
		},
		{
			name:   "not ok",
			domain: "stark.com",
			resp: response{
				code: http.StatusInternalServerError,
				body: "oh no",
			},
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
				w.WriteHeader(tt.resp.code)
				_, _ = w.Write([]byte(tt.resp.body))
			}))
			defer svr.Close()

			c, _ := NewClient("test-key", WithAPIURL(svr.URL))

			got := c.DeleteDomain(context.Background(), tt.domain)
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

func pointBool(b bool) *bool {
	return &b
}
