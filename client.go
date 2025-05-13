package esperio

import (
	"fmt"
	"net/http"
	"time"

	"github.com/esper-io/esper-go/requests"
	"github.com/esper-io/esper-go/resources"
)

var Request *requests.Request

type Client struct {

	// Resources # todo
	Device *resources.Device
}

func NewClient(tenant string, enterpriseID string, token string) *Client {
	baseURL := fmt.Sprintf("https://%s-api.esper.cloud", tenant)
	auth := requests.Auth{Token: token}
	httpClient := &http.Client{Timeout: 30 * time.Second}

	Request = &requests.Request{
		BaseURL:      baseURL,
		EnterpriseID: enterpriseID,
		Auth:         auth,
		HTTPClient:   httpClient,
	}

	device := resources.Device{Request: Request}

	client := &Client{
		Device: &device,
	}

	return client

}
