// Copyright Forward Email 2026
// SPDX-License-Identifier: MIT

package forwardemail

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestClient_EncryptRecord(t *testing.T) {
	const input = "forward-email-site-verification=abc123"

	response := `{
		"encrypted": "forward-email-site-verification=$2a$10$abc123"
	}`
	want := &EncryptionResponse{
		Encrypted: "forward-email-site-verification=$2a$10$abc123",
	}

	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST method, got %s", r.Method)
		}
		if r.URL.Path != "/v1/encrypt" {
			t.Errorf("expected URL /v1/encrypt, got %s", r.URL.Path)
		}

		body, err := io.ReadAll(r.Body)
		if err != nil {
			t.Fatalf("failed to read request body: %v", err)
		}
		if string(body) != fmt.Sprintf(`{"input": "%s"}`, input) {
			t.Errorf("unexpected request body %q", string(body))
		}

		fmt.Fprint(w, response)
	}))
	defer svr.Close()

	c, _ := NewClient("test-key", WithAPIURL(svr.URL))

	got, err := c.EncryptRecord(context.Background(), input)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Fatalf("values are not the same %s", diff)
	}
}

func TestClient_EncryptRecordMissingInput(t *testing.T) {
	c, _ := NewClient("test-key")

	_, err := c.EncryptRecord(context.Background(), "")
	if !errors.Is(err, ErrMissingEncryptionInput) {
		t.Fatalf("expected error %v, got %v", ErrMissingEncryptionInput, err)
	}
}
