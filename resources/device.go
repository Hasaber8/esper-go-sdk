package resources

import (
	"net/url"

	"github.com/Hasaber8/esper-go-sdk/requests"
)

type Device struct {
	Request *requests.Request
}

// List devices with optional filters
func (d *Device) List(filters map[string]string) (*requests.APIResponse, error) {
	endpoint := "/api/v2/devices"

	// Build query parameters
	queryParams := url.Values{}
	for key, value := range filters {
		queryParams.Add(key, value)
	}

	return d.Request.Get(endpoint, queryParams)
}
