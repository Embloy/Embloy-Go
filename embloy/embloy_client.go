package embloy

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// EmbloyClient represents the Embloy API client.
type EmbloyClient struct {
	ClientToken string
	Session     map[string]string
	BaseURL     string
	APIVersion  string
}

// NewEmbloyClient creates a new instance of EmbloyClient.
func NewEmbloyClient(clientToken string, session map[string]string) *EmbloyClient {
	return &EmbloyClient{
		ClientToken: clientToken,
		Session:     session,
		BaseURL:     "https://api.embloy.com",
		APIVersion:  "api/v0",
	}
}

// MakeRequest makes a request to the Embloy API.
func (c *EmbloyClient) MakeRequest() (string, error) {
	url := fmt.Sprintf("%s/%s/sdk/request/auth/token", c.BaseURL, c.APIVersion)
	headers := map[string]string{"client_token": c.ClientToken}
	data := map[string]string{
		"mode":        c.Session["mode"],
		"success_url": c.Session["success_url"],
		"cancel_url":  c.Session["cancel_url"],
		"job_slug":    c.Session["job_slug"],
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	request, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}

	for key, value := range headers {
		request.Header.Set(key, value)
	}

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
		return fmt.Sprintf("@%s/sdk/apply?token=%s", c.BaseURL, requestToken), nil
	}

	return "", fmt.Errorf("Error making request: %s", response.Status)
}
