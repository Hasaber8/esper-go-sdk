package main

import (
	"fmt"
	"log"
	"os"
	"time"

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

	devices := []string{"d774ae8c-7466-42df-a472-6f04b39b8907"}

	// Example 1: raw OTA update command
	fmt.Println("=== RAW OTA Update Command ===")
	complexBody := map[string]interface{}{
		"command_type": "DEVICE",
		"command":      "UPDATE_DEVICE_CONFIG",
		"command_args": map[string]interface{}{
			"custom_settings_config": map[string]interface{}{
				"scripts": []map[string]interface{}{
					{
						"action":     "LAUNCH",
						"launchType": "SERVICE",
						"actionParams": map[string]interface{}{
							"componentName": "io.esper.otamanager/io.esper.otamanager.OTAUpdateService",
							"intentAction":  "io.esper.otamanager.INSTALL_OTA",
							"serviceType":   "BACKGROUND",
							"extras": map[string]interface{}{
								"otaType":  "SEAMLESS",
								"metaData": `{"switch_slot_allowed":true,"ab_streaming_metadata":{"property_files":[{"filename":"payload_metadata.bin","size":"67033","offset":"2706"},{"filename":"payload.bin","size":"817266612","offset":"2706"},{"filename":"payload_properties.txt","size":"154","offset":"817269376"},{"filename":"apex_info.pb","size":"925","offset":"1569"},{"filename":"care_map.pb","size":"118","offset":"2541"},{"filename":"metadata","size":"653","offset":"69"},{"filename":"metadata.pb","size":"731","offset":"790"}]},"ab_install_type":"STREAMING","name":"13.6.2.1027-arm64-gsi-20240530-turnkey-TYD-STAGING-nPlusOne","header_key_value_pairs":["FILE_HASH=wPYevFfTewvnucpM3aRHF1RSxA2ELM6HY0Fp97AKBFA=","FILE_SIZE=817266612","METADATA_HASH=qbr5WBzNRvv8mybFXVQ1N6BNhpIX++0wnAN8VUhUAjI=","METADATA_SIZE=66766"],"url":"https://ota.esper.cloud/jenkins_builds/OSBuilds/sparrow/thirteen/arm64/1027/artifacts/foundation-13.6.2.1027-arm64-gsi-20240530-turnkey-TYD-STAGING-nPlusOne-fullota.zip"}`,
							},
						},
					},
				},
			},
		},
		"devices":     devices,
		"device_type": "all",
	}

	// directly using the SendCommand method
	apiResponse, err := client.Commands.SendCommand(complexBody)
	if err != nil {
		fmt.Println("Error sending command:", err)
	} else {
		fmt.Println("Command sent successfully:", apiResponse.PrettyString())
	}

	// Example 2: With validation and convenience methods
	fmt.Println("\n=== Using Convenience Methods ===")

	// Simple device operations
	resp, err := client.Commands.Reboot(devices)
	if err != nil {
		log.Printf("Reboot failed: %v", err)
	} else {
		fmt.Printf("Reboot command sent successfully\n")
	}

	// Set brightness
	resp, err = client.Commands.SetBrightness(devices, 70)
	if err != nil {
		log.Printf("Set brightness failed: %v", err)
	} else {
		fmt.Printf("Brightness set to 70%%\n")
	}

	// Send notification
	resp, err = client.Commands.NotifyDevice(
		devices,
		"Test Notification",
		"This is a test message from the SDK",
		"https://example.com",
	)
	if err != nil {
		log.Printf("Notification failed: %v", err)
	} else {
		fmt.Printf("Notification sent successfully\n")
	}

	// Example 3: Using UpdateDeviceConfig for custom settings
	fmt.Println("\n=== Custom Device Config ===")

	customConfig := map[string]interface{}{
		"custom_settings_config": map[string]interface{}{
			"wifi_settings": map[string]interface{}{
				"ssid":     "Office-WiFi",
				"security": "WPA2",
			},
			"display_settings": map[string]interface{}{
				"brightness": 80,
				"timeout":    30000,
			},
		},
	}

	resp, err = client.Commands.UpdateDeviceConfig(devices, customConfig)
	if err != nil {
		log.Printf("Update config failed: %v", err)
	} else {
		fmt.Printf("Device config updated successfully\n")
	}

	// Example 4: App management
	fmt.Println("\n=== App Management ===")

	// Install an app
	resp, err = client.Commands.InstallApp(devices, "app-version-id-12345")
	if err != nil {
		log.Printf("Install app failed: %v", err)
	} else {
		fmt.Printf("App installation initiated\n")
	}

	// Set kiosk app
	resp, err = client.Commands.SetKioskApp(devices, "com.example.kiosk")
	if err != nil {
		log.Printf("Set kiosk app failed: %v", err)
	} else {
		fmt.Printf("Kiosk app set successfully\n")
	}

	// Example 5: Scheduled commands
	fmt.Println("\n=== Scheduled Commands ===")

	// Schedule a reboot within a maintenance window
	startTime := time.Now().Add(1 * time.Hour)
	endTime := time.Now().Add(3 * time.Hour)

	resp, err = client.Commands.ScheduleRebootWindow(
		devices,
		startTime,
		endTime,
		"02:00", // Window starts at 2 AM
		"04:00", // Window ends at 4 AM
	)
	if err != nil {
		log.Printf("Schedule reboot failed: %v", err)
	} else {
		fmt.Printf("Reboot scheduled for maintenance window\n")
	}

	// Schedule recurring notifications
	resp, err = client.Commands.ScheduleRecurringNotification(
		devices,
		"Daily Reminder",
		"Check-in Reminder",
		"Please complete your daily check-in",
		time.Now(),
		time.Now().Add(30*24*time.Hour), // For 30 days
		[]string{"monday", "tuesday", "wednesday", "thursday", "friday"},
	)
	if err != nil {
		log.Printf("Schedule notification failed: %v", err)
	} else {
		fmt.Printf("Recurring notification scheduled\n")
	}

	// Example 6: Group commands
	fmt.Println("\n=== Group Commands ===")

	groups := []string{"production-devices", "warehouse-tablets"}

	// Apply policy to groups
	resp, err = client.Commands.ApplyPolicyToGroups(
		groups,
		"https://example.com/policies/production-policy.json",
	)
	if err != nil {
		log.Printf("Apply policy to groups failed: %v", err)
	} else {
		fmt.Printf("Policy applied to groups successfully\n")
	}

	// Reboot all devices in groups
	resp, err = client.Commands.RebootGroups(groups)
	if err != nil {
		log.Printf("Reboot groups failed: %v", err)
	} else {
		fmt.Printf("Reboot command sent to all devices in groups\n")
	}

	// Example 7: Advanced device settings
	fmt.Println("\n=== Advanced Device Settings ===")

	// Set GPS to high accuracy mode
	resp, err = client.Commands.SetGPSState(devices, 0) // 0 = High Accuracy
	if err != nil {
		log.Printf("Set GPS failed: %v", err)
	} else {
		fmt.Printf("GPS set to high accuracy mode\n")
	}

	// Set screen rotation to portrait only
	resp, err = client.Commands.SetRotationState(devices, 1) // 1 = Portrait Only
	if err != nil {
		log.Printf("Set rotation failed: %v", err)
	} else {
		fmt.Printf("Screen rotation set to portrait only\n")
	}

	// Set music volume to 50%
	resp, err = client.Commands.SetVolume(devices, 3, 50) // 3 = Music stream
	if err != nil {
		log.Printf("Set volume failed: %v", err)
	} else {
		fmt.Printf("Music volume set to 50%%\n")
	}

	// Enable Bluetooth
	resp, err = client.Commands.SetBluetoothState(devices, true)
	if err != nil {
		log.Printf("Set Bluetooth failed: %v", err)
	} else {
		fmt.Printf("Bluetooth enabled\n")
	}

	// Set device language
	resp, err = client.Commands.SetDeviceLanguage(devices, "en_US")
	if err != nil {
		log.Printf("Set language failed: %v", err)
	} else {
		fmt.Printf("Device language set to en_US\n")
	}

	// Lock device with message
	resp, err = client.Commands.SetDeviceLockdown(
		devices,
		true,
		"Device is under maintenance. Please contact IT support.",
	)
	if err != nil {
		log.Printf("Set lockdown failed: %v", err)
	} else {
		fmt.Printf("Device locked with maintenance message\n")
	}

	// Capture screenshot with tag
	resp, err = client.Commands.CaptureScreenshot(devices, "audit-screenshot-2024")
	if err != nil {
		log.Printf("Capture screenshot failed: %v", err)
	} else {
		fmt.Printf("Screenshot captured with tag\n")
	}

	// Make device beep for 3 seconds
	resp, err = client.Commands.BeepDevice(devices, "3")
	if err != nil {
		log.Printf("Beep device failed: %v", err)
	} else {
		fmt.Printf("Device beep command sent\n")
	}

	_ = resp // Suppress unused variable warning
}
