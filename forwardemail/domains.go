package forwardemail

import (
	"encoding/json"
	"fmt"
	"io"
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
func (c *Client) GetDomains() ([]Domain, error) {
	req, err := c.newRequest("GET", "/v1/domains")
	if err != nil {
		return nil, err
	}

	res, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var items []Domain

	err = json.Unmarshal(res, &items)
	if err != nil {
		return nil, err
	}

	return items, nil
}

// GetDomain retrieves a specific domain by name.
func (c *Client) GetDomain(name string) (*Domain, error) {
	req, err := c.newRequest("GET", fmt.Sprintf("/v1/domains/%s", name))
	if err != nil {
		return nil, err
	}

	res, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var item Domain

	err = json.Unmarshal(res, &item)
	if err != nil {
		return nil, err
	}

	return &item, nil
}

// CreateDomain adds a new domain to the account with the specified configuration parameters.
func (c *Client) CreateDomain(name string, parameters DomainParameters) (*Domain, error) {
	req, err := c.newRequest("POST", "/v1/domains")
	if err != nil {
		return nil, err
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

	req.Body = io.NopCloser(strings.NewReader(params.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	res, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var item Domain

	err = json.Unmarshal(res, &item)
	if err != nil {
		return nil, err
	}

	return &item, nil
}

// UpdateDomain modifies an existing domain's configuration parameters.
func (c *Client) UpdateDomain(name string, parameters DomainParameters) (*Domain, error) {
	req, err := c.newRequest("PUT", fmt.Sprintf("/v1/domains/%s", name))
	if err != nil {
		return nil, err
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

	req.Body = io.NopCloser(strings.NewReader(params.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	res, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var item Domain

	err = json.Unmarshal(res, &item)
	if err != nil {
		return nil, err
	}

	return &item, nil
}

// DeleteDomain removes a domain from the account.
func (c *Client) DeleteDomain(name string) error {
	req, err := c.newRequest("DELETE", fmt.Sprintf("/v1/domains/%s", name))
	if err != nil {
		return err
	}

	_, err = c.doRequest(req)
	if err != nil {
		return err
	}

	return nil
}
