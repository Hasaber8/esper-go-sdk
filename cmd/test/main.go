package main

import (
	"fmt"
	"os"

	esperio "github.com/Hasaber8/esper-go-sdk"
)

func main() {
	enterpriseID := os.Getenv("ESPER_ENTERPRISE_ID")
	token := os.Getenv("ESPER_TOKEN")

	if enterpriseID == "" {
		fmt.Println("Error: ESPER_ENTERPRISE_ID environment variable is required")
		return
	}
	if token == "" {
		fmt.Println("Error: ESPER_TOKEN environment variable is required")
		return
	}

	client := esperio.NewClient("develop", enterpriseID, token)

	// List devices with filters
	filters := map[string]string{
		"limit":  "2",
		"offset": "0",
	}

	devices, err := client.Device.List(filters)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Devices: %v\n", devices)
}
