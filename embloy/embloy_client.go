package embloy

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// SessionData represents the allowed session values for an EmbloyClient.
type SessionData struct {
	Mode       string
	SuccessURL string
	CancelURL  string
	JobSlug    string
}

// EmbloyClient represents the Embloy API client.
type EmbloyClient struct {
	ClientToken string
	Session     SessionData
	BaseURL     string
	APIURL      string
	APIVersion  string
	HTTPClient  HTTPClient
}

// NewEmbloyClient creates a new instance of EmbloyClient.
func NewEmbloyClient(clientToken string, session SessionData) *EmbloyClient {
	return &EmbloyClient{
		ClientToken: clientToken,
		Session:     session,
		BaseURL:     "https://embloy.com",
		APIURL:      "https://api.embloy.com",
		APIVersion:  "api/v0",
		HTTPClient:  &http.Client{},
	}
}

// MakeRequest makes a request to the Embloy API.
func (c *EmbloyClient) MakeRequest() (string, error) {
	requestURL := fmt.Sprintf("%s/%s/sdk/request/auth/token", c.APIURL, c.APIVersion)
	headers := map[string]string{"client_token": c.ClientToken}
	data := url.Values{
		"mode":        {c.Session.Mode},
		"success_url": {c.Session.SuccessURL},
		"cancel_url":  {c.Session.CancelURL},
		"job_slug":    {c.Session.JobSlug},
	}

	request, err := http.NewRequest("POST", requestURL, strings.NewReader(data.Encode()))
	if err != nil {
		return "", err
	}

	for key, value := range headers {
		request.Header.Set(key, value)
	}
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64; rv:122.0) Gecko/20100101 Firefox/122.0")

	response, err := c.HTTPClient.Do(request)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	// Ensure debug info includes response headers if request was made
	debugInfo := map[string]interface{}{
		"client_token":     c.ClientToken,
		"error":            "",
		"request_headers":  headers,
		"response_headers": response.Header,
	}

	fmt.Println("Debug Info:", debugInfo)

	return c.handleResponse(response)
}

func (c *EmbloyClient) handleResponse(response *http.Response) (string, error) {
	if response.StatusCode == http.StatusOK {
		var result map[string]interface{}
		if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
			return "", err
		}
		// Safely assert requestToken to string, checking for nil
		requestToken, ok := result["request_token"]
		if !ok || requestToken == nil {
			return "", errors.New("request_token is missing or nil")
		}
		requestTokenStr, ok := requestToken.(string)
		if !ok {
			return "", errors.New("request_token is not a string")
		}
		return fmt.Sprintf("%s/sdk/apply?request_token=%s", c.BaseURL, requestTokenStr), nil
	}

	return "", fmt.Errorf("error making request: %s", response.Status)
}
