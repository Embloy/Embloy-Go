package embloy

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// SessionData represents the allowed session values for an EmbloyClient.
type SessionData struct {
	Mode       string      `json:"mode"`
	SuccessURL string      `json:"success_url"`
	CancelURL  string      `json:"cancel_url"`
	JobSlug    string      `json:"job_slug"`
	Proxy      interface{} `json:"proxy,omitempty"` // Used for internal purposes. Ignore this field.
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
	if c.Session.Proxy != nil {
		requestURL = fmt.Sprintf("%s/%s/sdk/request/auth/proxy", c.APIURL, c.APIVersion)
	}

	headers := map[string]string{"client_token": c.ClientToken}
	data := SessionData{
		Mode:       c.Session.Mode,
		SuccessURL: c.Session.SuccessURL,
		CancelURL:  c.Session.CancelURL,
		JobSlug:    c.Session.JobSlug,
	}
	if c.Session.Proxy != nil {
		data.Proxy = c.Session.Proxy
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error marshaling data: ", err)
		return "", err
	}

	request, err := http.NewRequest("POST", requestURL, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error creating request: ", err)
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
