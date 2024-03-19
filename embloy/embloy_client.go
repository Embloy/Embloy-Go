package embloy

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

// EmbloyClient represents the Embloy API client.
type EmbloyClient struct {
	ClientToken string
	Session     map[string]string
	BaseURL     string
	APIURL      string
	APIVersion  string
}

// NewEmbloyClient creates a new instance of EmbloyClient.
func NewEmbloyClient(clientToken string, session map[string]string) *EmbloyClient {
	return &EmbloyClient{
		ClientToken: clientToken,
		Session:     session,
		BaseURL:     "https://embloy.com",
		APIURL:      "https://api.embloy.com",
		APIVersion:  "api/v0",
	}
}

// MakeRequest makes a request to the Embloy API.
func (c *EmbloyClient) MakeRequest() (string, error) {
	requestURL := fmt.Sprintf("%s/%s/sdk/request/auth/token", c.APIURL, c.APIVersion) // Renamed 'url' to 'requestURL'
	headers := map[string]string{"client_token": c.ClientToken}
	data := url.Values{
		"mode":        {c.Session["mode"]},
		"success_url": {c.Session["success_url"]},
		"cancel_url":  {c.Session["cancel_url"]},
		"job_slug":    {c.Session["job_slug"]},
	}

	request, err := http.NewRequest("POST", requestURL, strings.NewReader(data.Encode())) // Use 'requestURL' here
	if err != nil {
		return "", err
	}

	for key, value := range headers {
		request.Header.Set(key, value)
	}
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	response, err := client.Do(request)
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
		requestToken := result["request_token"].(string)
		return fmt.Sprintf("%s/sdk/apply?request_token=%s", c.BaseURL, requestToken), nil
	}

	return "", fmt.Errorf("Error making request: %s", response.Status)
}
