package forwardemail

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

func TestClient_GetAccount(t *testing.T) {
	tests := []struct {
		name     string
		response string
		want     *Account
	}{
		{
			name: "no data",
		},
		{
			name: "ok",
			response: `{
  				"plan": "enhanced_protection",
  				"email": "tony@stark.com",
  				"full_email": "tony@stark.com",
  				"display_name": "tony@stark.com",
  				"last_locale": "en",
  				"address_country": "None",
  				"id": "59ad551ae6fb4a4c53427ca38079f029",
  				"object": "user",
  				"locale": "en",
  				"created_at": "2023-09-21T20:14:27.964Z",
  				"updated_at": "2023-10-07T17:47:54.595Z",
  				"address_html": ""
			}`,
			want: &Account{
				Plan:           "enhanced_protection",
				Email:          "tony@stark.com",
				FullEmail:      "tony@stark.com",
				DisplayName:    "tony@stark.com",
				LastLocale:     "en",
				AddressCountry: "None",
				ID:             "59ad551ae6fb4a4c53427ca38079f029",
				Object:         "user",
				Locale:         "en",
				CreatedAt:      parseTime("2023-09-21T20:14:27.964Z"),
				UpdatedAt:      parseTime("2023-10-07T17:47:54.595Z"),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
				fmt.Fprintf(w, tt.response)
			}))
			defer svr.Close()

			c := NewClient(ClientOptions{
				APIURL: svr.URL,
			})

			got, _ := c.GetAccount()
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Fatalf("values are not the same %s", diff)
			}
		})
	}
}

func parseTime(str string) time.Time {
	t, _ := time.Parse(time.RFC3339, str)

	return t
}
