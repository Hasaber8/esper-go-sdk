package resources

import (
	"fmt"
	"net/url"

	"github.com/esper-io/esper-go/requests"
)

type Device struct {
	Request *requests.Request
}

// List devices with optional filters
func (d *Device) List(filters map[string]string) (map[string]interface{}, error) {
	endpoint := fmt.Sprintf("/api/enterprise/%s/device/", d.Request.EnterpriseID)

	// Build query parameters
	queryParams := url.Values{}
	for key, value := range filters {
		queryParams.Add(key, value)
	}

	return d.Request.Get(endpoint, queryParams)
}

// Get a specific device
func (d *Device) Get(deviceID string) (map[string]interface{}, error) {
	endpoint := fmt.Sprintf("/api/enterprise/%s/device/%s/", d.Request.EnterpriseID, deviceID)
	return d.Request.Get(endpoint, nil)
}
