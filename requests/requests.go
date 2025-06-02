package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Request struct {
	BaseURL      string
	EnterpriseID string
	Auth         Auth
	HTTPClient   *http.Client
}

type Auth struct {
	Token string
}

type APIResponse struct {
	Data map[string]interface{}
}

func (r *APIResponse) PrettyString() string {
	prettyJson, err := json.MarshalIndent(r.Data, "", " ")
	if err != nil {
		return fmt.Sprintf("Error formatting JSON: %v", err)
	}
	return string(prettyJson)
}

// String implements the Stringer interface for default fmt.Printf behavior
func (r *APIResponse) String() string {
	return r.PrettyString()
}

// Get returns the underlying data
func (r *APIResponse) Get() map[string]interface{} {
	return r.Data
}

// Modified Get method to return APIResponse
func (request *Request) Get(endpoint string, queryParam url.Values) (*APIResponse, error) {
	fullURL := request.BaseURL + endpoint
	if queryParam != nil && len(queryParam) > 0 {
		fullURL += "?" + queryParam.Encode()
	}

	// Create request
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return nil, err
	}

	// Add headers
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", request.Auth.Token))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-Caller-Id", "Esper-sdk")
	req.Header.Add("X-Tenant-Id", request.EnterpriseID)

	// Make the request
	resp, err := request.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Handle error responses
	if resp.StatusCode >= 400 {
		var errorResp map[string]interface{}
		if err = json.Unmarshal(body, &errorResp); err != nil {
			return nil, fmt.Errorf("Error %d: %s", resp.StatusCode, string(body))
		}
		return nil, fmt.Errorf("API Error: %v", errorResp)
	}

	// Parse JSON response
	var result map[string]interface{}
	if err = json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return &APIResponse{Data: result}, nil
}
