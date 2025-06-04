package esperio

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Hasaber8/esper-go-sdk/requests"
	"github.com/Hasaber8/esper-go-sdk/resources"
)

var Request *requests.Request

type Client struct {

	// Resources # todo
	Device   *resources.Device
	Commands *resources.Commands
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
	commands := resources.Commands{Request: Request}

	client := &Client{
		Device:   &device,
		Commands: &commands,
	}

	return client

}
