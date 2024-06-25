package embloy_test

import (
	"io"
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/embloy/embloy-go/embloy"
)

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type MockHTTPClient struct {
	LastRequest *http.Request
}

func (m *MockHTTPClient) Do(req *http.Request) (*http.Response, error) {
	m.LastRequest = req
	// Return a dummy response with a simple JSON body
	responseBody := `{"request_token":"dummyToken", "status":"ok"}`
	return &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(strings.NewReader(responseBody)),
		Header:     http.Header{"Content-Type": []string{"application/json"}},
	}, nil
}

func TestMakeRequest(t *testing.T) {
	mockClient := &MockHTTPClient{}

	client := &embloy.EmbloyClient{
		ClientToken: "testToken",
		Session: embloy.SessionData{
			Mode:       "testMode",
			SuccessURL: "https://success.url",
			CancelURL:  "https://cancel.url",
			JobSlug:    "testJobSlug",
		},
		BaseURL:    "https://embloy.com",
		APIURL:     "https://api.embloy.com",
		APIVersion: "api/v0",
		HTTPClient: mockClient,
	}

	_, err := client.MakeRequest()
	if err != nil {
		t.Fatalf("MakeRequest failed: %v", err)
	}

	req := mockClient.LastRequest
	if req == nil {
		t.Fatal("Expected a request to be made, but none was captured")
	}

	// Verify the request URL
	expectedURL := "https://api.embloy.com/api/v0/sdk/request/auth/token"
	if req.URL.String() != expectedURL {
		t.Errorf("Expected URL to be %s, got %s", expectedURL, req.URL.String())
	}

	// Verify the request method
	if req.Method != "POST" {
		t.Errorf("Expected method POST, got %s", req.Method)
	}

	// Verify headers
	if req.Header.Get("client_token") != "testToken" {
		t.Errorf("Expected client_token to be 'testToken', got '%s'", req.Header.Get("client_token"))
	}

	// Verify the request body
	expectedData := url.Values{
		"mode":        {"testMode"},
		"success_url": {"https://success.url"},
		"cancel_url":  {"https://cancel.url"},
		"job_slug":    {"testJobSlug"},
	}.Encode()
	body, _ := io.ReadAll(req.Body)
	if string(body) != expectedData {
		t.Errorf("Expected body to be %s, got %s", expectedData, string(body))
	}
}
