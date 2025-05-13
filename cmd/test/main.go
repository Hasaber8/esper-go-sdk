package main

import (
	"fmt"

	esperio "github.com/esper-io/esper-go"
)

func main() {
	client := esperio.NewClient("develop", "f44373cb-1800-43c6-aab3-c81f8b1f435c", "hS5Ha4VpmZ0SkUBTBG0ZFLdEiYmJ2F")

	// List devices with filters
	filters := map[string]string{
		"limit":  "20",
		"offset": "0",
		"brand":  "Samsung",
	}

	devices, err := client.Device.List(filters)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Devices: %v\n", devices)
}
