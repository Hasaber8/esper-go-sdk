package resources

import (
	"fmt"
	"time"

	"github.com/Hasaber8/esper-go-sdk/requests"
)

// Commands handles command-related API operations
type Commands struct {
	Request *requests.Request
}

// CommandType represents the type of command target
type CommandType string

const (
	CommandTypeDevice  CommandType = "DEVICE"  // Command for specific devices
	CommandTypeGroup   CommandType = "GROUP"   // Command for device groups
	CommandTypeDynamic CommandType = "DYNAMIC" // Command for dynamic device sets
)

// Command represents the actual command to execute
type Command string

// Device management commands
const (
	CommandAddToWhitelist          Command = "ADD_TO_WHITELIST"
	CommandAddWifiAP               Command = "ADD_WIFI_AP"
	CommandClearAppData            Command = "CLEAR_APP_DATA"
	CommandInstall                 Command = "INSTALL"
	CommandLock                    Command = "LOCK"
	CommandReboot                  Command = "REBOOT"
	CommandRemoveFromWhitelist     Command = "REMOVE_FROM_WHITELIST"
	CommandRemoveWifiAP            Command = "REMOVE_WIFI_AP"
	CommandSetAppPermission        Command = "SET_APP_PERMISSION"
	CommandSetAppState             Command = "SET_APP_STATE"
	CommandSetBluetoothState       Command = "SET_BLUETOOTH_STATE"
	CommandSetBrightnessScale      Command = "SET_BRIGHTNESS_SCALE"
	CommandSetDeviceLockdownState  Command = "SET_DEVICE_LOCKDOWN_STATE"
	CommandSetGPSState             Command = "SET_GPS_STATE"
	CommandSetKioskApp             Command = "SET_KIOSK_APP"
	CommandSetNewPolicy            Command = "SET_NEW_POLICY"
	CommandSetRotationState        Command = "SET_ROTATION_STATE"
	CommandSetScreenOffTimeout     Command = "SET_SCREEN_OFF_TIMEOUT"
	CommandSetStreamVolume         Command = "SET_STREAM_VOLUME"
	CommandSetTimezone             Command = "SET_TIMEZONE"
	CommandSetWifiState            Command = "SET_WIFI_STATE"
	CommandUninstall               Command = "UNINSTALL"
	CommandUpdateDeviceConfig      Command = "UPDATE_DEVICE_CONFIG"
	CommandUpdateHeartbeat         Command = "UPDATE_HEARTBEAT"
	CommandUpdateLatestDPC         Command = "UPDATE_LATEST_DPC"
	CommandWipe                    Command = "WIPE"
	CommandResetLockscreenPassword Command = "RESET_LOCKSCREEN_PASSWORD"
	CommandCaptureScreenshot       Command = "CAPTURE_SCREENSHOT"
	CommandUpdateBlueprint         Command = "UPDATE_BLUEPRINT"
	CommandNotifyDevice            Command = "NOTIFY_DEVICE"
	CommandSetDeviceLanguage       Command = "SET_DEVICE_LANGUAGE"
	CommandSetEthernetSettings     Command = "SET_ETHERNET_SETTINGS"
	CommandSetStaticIP             Command = "SET_STATIC_IP"
	CommandBeepDevice              Command = "BEEP_DEVICE"
	CommandSetAppNotifications     Command = "SET_APP_NOTIFICATIONS"
	CommandUseOnlySavedWifiAP      Command = "USE_ONLY_SAVED_WIFI_AP"
	CommandConverge                Command = "CONVERGE"
)

// ScheduleType represents when the command should be executed
type ScheduleType string

const (
	ScheduleImmediate ScheduleType = "IMMEDIATE" // Execute immediately
	ScheduleWindow    ScheduleType = "WINDOW"    // Execute within a time window
	ScheduleRecurring ScheduleType = "RECURRING" // Execute repeatedly until success
)

// SendCommand sends a command with the given body to the commands endpoint
// This maintains backward compatibility with your existing code
func (c *Commands) SendCommand(body map[string]interface{}) (*requests.APIResponse, error) {
	endpoint := fmt.Sprintf("/api/v0/enterprise/%s/command/", c.Request.EnterpriseID)
	return c.Request.Post(endpoint, body)
}

// Convenience methods for common operations

// Reboot reboots the specified devices
func (c *Commands) Reboot(devices []string) (*requests.APIResponse, error) {
	body := map[string]interface{}{
		"command_type": string(CommandTypeDevice),
		"devices":      devices,
		"command":      string(CommandReboot),
		"schedule":     string(ScheduleImmediate),
	}
	return c.SendCommand(body)
}

// Lock locks the specified devices
func (c *Commands) Lock(devices []string) (*requests.APIResponse, error) {
	body := map[string]interface{}{
		"command_type": string(CommandTypeDevice),
		"devices":      devices,
		"command":      string(CommandLock),
		"schedule":     string(ScheduleImmediate),
	}
	return c.SendCommand(body)
}

// Wipe wipes the specified devices
func (c *Commands) Wipe(devices []string) (*requests.APIResponse, error) {
	body := map[string]interface{}{
		"command_type": string(CommandTypeDevice),
		"devices":      devices,
		"command":      string(CommandWipe),
		"schedule":     string(ScheduleImmediate),
	}
	return c.SendCommand(body)
}

// InstallApp installs an app on devices
func (c *Commands) InstallApp(devices []string, appVersionID string) (*requests.APIResponse, error) {
	body := map[string]interface{}{
		"command_type": string(CommandTypeDevice),
		"devices":      devices,
		"command":      string(CommandInstall),
		"command_args": map[string]interface{}{
			"app_version": appVersionID,
		},
		"schedule": string(ScheduleImmediate),
	}
	return c.SendCommand(body)
}

// UninstallApp uninstalls an app from devices
func (c *Commands) UninstallApp(devices []string, packageName string) (*requests.APIResponse, error) {
	body := map[string]interface{}{
		"command_type": string(CommandTypeDevice),
		"devices":      devices,
		"command":      string(CommandUninstall),
		"command_args": map[string]interface{}{
			"package_name": packageName,
		},
		"schedule": string(ScheduleImmediate),
	}
	return c.SendCommand(body)
}

// ClearAppData clears app data on devices
func (c *Commands) ClearAppData(devices []string, packageName string) (*requests.APIResponse, error) {
	body := map[string]interface{}{
		"command_type": string(CommandTypeDevice),
		"devices":      devices,
		"command":      string(CommandClearAppData),
		"command_args": map[string]interface{}{
			"package_name": packageName,
		},
		"schedule": string(ScheduleImmediate),
	}
	return c.SendCommand(body)
}

// SetKioskApp sets the kiosk app on devices
func (c *Commands) SetKioskApp(devices []string, packageName string) (*requests.APIResponse, error) {
	body := map[string]interface{}{
		"command_type": string(CommandTypeDevice),
		"devices":      devices,
		"command":      string(CommandSetKioskApp),
		"command_args": map[string]interface{}{
			"package_name": packageName,
		},
		"schedule": string(ScheduleImmediate),
	}
	return c.SendCommand(body)
}

// SetAppState sets the state of an app (SHOW/HIDE/DISABLE)
func (c *Commands) SetAppState(devices []string, packageName string, state string) (*requests.APIResponse, error) {
	body := map[string]interface{}{
		"command_type": string(CommandTypeDevice),
		"devices":      devices,
		"command":      string(CommandSetAppState),
		"command_args": map[string]interface{}{
			"package_name": packageName,
			"app_state":    state,
		},
		"schedule": string(ScheduleImmediate),
	}
	return c.SendCommand(body)
}

// SetBrightness sets the brightness on devices (1-100)
func (c *Commands) SetBrightness(devices []string, brightness int) (*requests.APIResponse, error) {
	if brightness < 1 || brightness > 100 {
		return nil, fmt.Errorf("brightness must be between 1 and 100")
	}

	body := map[string]interface{}{
		"command_type": string(CommandTypeDevice),
		"devices":      devices,
		"command":      string(CommandSetBrightnessScale),
		"command_args": map[string]interface{}{
			"brightness_value": brightness,
		},
		"schedule": string(ScheduleImmediate),
	}
	return c.SendCommand(body)
}

// SetVolume sets the volume on devices
// stream: 0=Ring, 1=Notification, 2=Alarm, 3=Music
// volume: 0-100
func (c *Commands) SetVolume(devices []string, stream, volume int) (*requests.APIResponse, error) {
	if stream < 0 || stream > 3 {
		return nil, fmt.Errorf("stream must be 0-3")
	}
	if volume < 0 || volume > 100 {
		return nil, fmt.Errorf("volume must be 0-100")
	}

	body := map[string]interface{}{
		"command_type": string(CommandTypeDevice),
		"devices":      devices,
		"command":      string(CommandSetStreamVolume),
		"command_args": map[string]interface{}{
			"stream":       stream,
			"volume_level": volume,
		},
		"schedule": string(ScheduleImmediate),
	}
	return c.SendCommand(body)
}

// SetWifiState enables or disables WiFi on devices
func (c *Commands) SetWifiState(devices []string, enabled bool) (*requests.APIResponse, error) {
	body := map[string]interface{}{
		"command_type": string(CommandTypeDevice),
		"devices":      devices,
		"command":      string(CommandSetWifiState),
		"command_args": map[string]interface{}{
			"wifi_state": enabled,
		},
		"schedule": string(ScheduleImmediate),
	}
	return c.SendCommand(body)
}

// SetBluetoothState enables or disables Bluetooth on devices
func (c *Commands) SetBluetoothState(devices []string, enabled bool) (*requests.APIResponse, error) {
	body := map[string]interface{}{
		"command_type": string(CommandTypeDevice),
		"devices":      devices,
		"command":      string(CommandSetBluetoothState),
		"command_args": map[string]interface{}{
			"bluetooth_state": enabled,
		},
		"schedule": string(ScheduleImmediate),
	}
	return c.SendCommand(body)
}

// UpdateDeviceConfig updates device configuration
func (c *Commands) UpdateDeviceConfig(devices []string, config map[string]interface{}) (*requests.APIResponse, error) {
	body := map[string]interface{}{
		"command_type": string(CommandTypeDevice),
		"devices":      devices,
		"command":      string(CommandUpdateDeviceConfig),
		"command_args": config,
		"schedule":     string(ScheduleImmediate),
	}
	// Special handling for device_type if not in config
	if _, ok := config["device_type"]; !ok {
		body["device_type"] = "all"
	}
	return c.SendCommand(body)
}

// NotifyDevice sends a notification to devices
func (c *Commands) NotifyDevice(devices []string, title, message string, url ...string) (*requests.APIResponse, error) {
	args := map[string]interface{}{
		"title":   title,
		"message": message,
	}
	if len(url) > 0 {
		args["url"] = url[0]
	}

	body := map[string]interface{}{
		"command_type": string(CommandTypeDevice),
		"devices":      devices,
		"command":      string(CommandNotifyDevice),
		"command_args": args,
		"schedule":     string(ScheduleImmediate),
	}
	return c.SendCommand(body)
}

// CaptureScreenshot captures a screenshot on devices
func (c *Commands) CaptureScreenshot(devices []string, tag ...string) (*requests.APIResponse, error) {
	args := map[string]interface{}{}
	if len(tag) > 0 {
		args["tag"] = tag[0]
	}

	body := map[string]interface{}{
		"command_type": string(CommandTypeDevice),
		"devices":      devices,
		"command":      string(CommandCaptureScreenshot),
		"command_args": args,
		"schedule":     string(ScheduleImmediate),
	}
	return c.SendCommand(body)
}

// SetDeviceLanguage sets the language on devices
func (c *Commands) SetDeviceLanguage(devices []string, locale string) (*requests.APIResponse, error) {
	body := map[string]interface{}{
		"command_type": string(CommandTypeDevice),
		"devices":      devices,
		"command":      string(CommandSetDeviceLanguage),
		"command_args": map[string]interface{}{
			"locale": locale,
		},
		"schedule": string(ScheduleImmediate),
	}
	return c.SendCommand(body)
}

// BeepDevice makes devices beep for a specified duration
func (c *Commands) BeepDevice(devices []string, duration string) (*requests.APIResponse, error) {
	body := map[string]interface{}{
		"command_type": string(CommandTypeDevice),
		"devices":      devices,
		"command":      string(CommandBeepDevice),
		"command_args": map[string]interface{}{
			"duration": duration,
		},
		"schedule": string(ScheduleImmediate),
	}
	return c.SendCommand(body)
}

// ResetPassword resets the lockscreen password
func (c *Commands) ResetPassword(devices []string, newPassword string) (*requests.APIResponse, error) {
	body := map[string]interface{}{
		"command_type": string(CommandTypeDevice),
		"devices":      devices,
		"command":      string(CommandResetLockscreenPassword),
		"command_args": map[string]interface{}{
			"new_lockscreen_password": newPassword,
		},
		"schedule": string(ScheduleImmediate),
	}
	return c.SendCommand(body)
}

// UpdateBlueprint pushes or reapplies the current Blueprint to devices
func (c *Commands) UpdateBlueprint(devices []string) (*requests.APIResponse, error) {
	body := map[string]interface{}{
		"command_type": string(CommandTypeDevice),
		"devices":      devices,
		"command":      string(CommandUpdateBlueprint),
		"schedule":     string(ScheduleImmediate),
	}
	return c.SendCommand(body)
}

// SetGPSState sets GPS state
// state: 0=High Accuracy, 1=Sensors Only, 2=Battery Saving, 3=Off, 4=On
func (c *Commands) SetGPSState(devices []string, state int) (*requests.APIResponse, error) {
	if state < 0 || state > 4 {
		return nil, fmt.Errorf("gps_state must be between 0 and 4")
	}

	body := map[string]interface{}{
		"command_type": string(CommandTypeDevice),
		"devices":      devices,
		"command":      string(CommandSetGPSState),
		"command_args": map[string]interface{}{
			"gps_state": state,
		},
		"schedule": string(ScheduleImmediate),
	}
	return c.SendCommand(body)
}

// SetRotationState sets screen orientation
// state: 0=Auto, 1=Portrait Only, 2=Landscape Only
func (c *Commands) SetRotationState(devices []string, state int) (*requests.APIResponse, error) {
	if state < 0 || state > 2 {
		return nil, fmt.Errorf("rotate_state must be 0, 1, or 2")
	}

	body := map[string]interface{}{
		"command_type": string(CommandTypeDevice),
		"devices":      devices,
		"command":      string(CommandSetRotationState),
		"command_args": map[string]interface{}{
			"rotate_state": state,
		},
		"schedule": string(ScheduleImmediate),
	}
	return c.SendCommand(body)
}

// SetScreenOffTimeout sets screen off timeout
// timeout: -1 or between 5000 and 1800000 milliseconds
func (c *Commands) SetScreenOffTimeout(devices []string, timeout int) (*requests.APIResponse, error) {
	if timeout != -1 && (timeout < 5000 || timeout > 1800000) {
		return nil, fmt.Errorf("screen_off_timeout must be -1 or between 5000 and 1800000")
	}

	body := map[string]interface{}{
		"command_type": string(CommandTypeDevice),
		"devices":      devices,
		"command":      string(CommandSetScreenOffTimeout),
		"command_args": map[string]interface{}{
			"screen_off_timeout": timeout,
		},
		"schedule": string(ScheduleImmediate),
	}
	return c.SendCommand(body)
}

// SetTimezone sets the timezone for devices
func (c *Commands) SetTimezone(devices []string, timezone string) (*requests.APIResponse, error) {
	body := map[string]interface{}{
		"command_type": string(CommandTypeDevice),
		"devices":      devices,
		"command":      string(CommandSetTimezone),
		"command_args": map[string]interface{}{
			"timezone_string": timezone,
		},
		"schedule": string(ScheduleImmediate),
	}
	return c.SendCommand(body)
}

// ApplyPolicy applies a policy to devices
func (c *Commands) ApplyPolicy(devices []string, policyURL string) (*requests.APIResponse, error) {
	body := map[string]interface{}{
		"command_type": string(CommandTypeDevice),
		"devices":      devices,
		"command":      string(CommandSetNewPolicy),
		"command_args": map[string]interface{}{
			"policy_url": policyURL,
		},
		"schedule": string(ScheduleImmediate),
	}
	return c.SendCommand(body)
}

// SetDeviceLockdown sets lockdown state for devices
func (c *Commands) SetDeviceLockdown(devices []string, locked bool, message string) (*requests.APIResponse, error) {
	state := "UNLOCKED"
	if locked {
		state = "LOCKED"
	}

	body := map[string]interface{}{
		"command_type": string(CommandTypeDevice),
		"devices":      devices,
		"command":      string(CommandSetDeviceLockdownState),
		"command_args": map[string]interface{}{
			"state":   state,
			"message": message,
		},
		"schedule": string(ScheduleImmediate),
	}
	return c.SendCommand(body)
}

// Group command methods

// SendGroupCommand sends a command to device groups
func (c *Commands) SendGroupCommand(groups []string, command Command, args map[string]interface{}) (*requests.APIResponse, error) {
	body := map[string]interface{}{
		"command_type": string(CommandTypeGroup),
		"groups":       groups,
		"command":      string(command),
		"schedule":     string(ScheduleImmediate),
	}
	if args != nil {
		body["command_args"] = args
	}
	return c.SendCommand(body)
}

// RebootGroups reboots all devices in specified groups
func (c *Commands) RebootGroups(groups []string) (*requests.APIResponse, error) {
	return c.SendGroupCommand(groups, CommandReboot, nil)
}

// LockGroups locks all devices in specified groups
func (c *Commands) LockGroups(groups []string) (*requests.APIResponse, error) {
	return c.SendGroupCommand(groups, CommandLock, nil)
}

// ApplyPolicyToGroups applies a policy to all devices in groups
func (c *Commands) ApplyPolicyToGroups(groups []string, policyURL string) (*requests.APIResponse, error) {
	args := map[string]interface{}{
		"policy_url": policyURL,
	}
	return c.SendGroupCommand(groups, CommandSetNewPolicy, args)
}

// Scheduled command helpers

// SendScheduledCommand sends a command with custom scheduling
func (c *Commands) SendScheduledCommand(body map[string]interface{}, scheduleType ScheduleType, scheduleArgs map[string]interface{}) (*requests.APIResponse, error) {
	body["schedule"] = string(scheduleType)
	if scheduleArgs != nil {
		body["schedule_args"] = scheduleArgs
	}
	return c.SendCommand(body)
}

// ScheduleRebootWindow schedules a reboot within a time window
func (c *Commands) ScheduleRebootWindow(devices []string, startTime, endTime time.Time, windowStart, windowEnd string) (*requests.APIResponse, error) {
	body := map[string]interface{}{
		"command_type": string(CommandTypeDevice),
		"devices":      devices,
		"command":      string(CommandReboot),
	}

	scheduleArgs := map[string]interface{}{
		"start_datetime":    startTime.Format(time.RFC3339),
		"end_datetime":      endTime.Format(time.RFC3339),
		"window_start_time": windowStart,
		"window_end_time":   windowEnd,
		"time_type":         "console",
	}

	return c.SendScheduledCommand(body, ScheduleWindow, scheduleArgs)
}

// ScheduleRecurringNotification schedules recurring notifications
func (c *Commands) ScheduleRecurringNotification(devices []string, name, title, message string, startTime, endTime time.Time, days []string) (*requests.APIResponse, error) {
	body := map[string]interface{}{
		"command_type": string(CommandTypeDevice),
		"devices":      devices,
		"command":      string(CommandNotifyDevice),
		"command_args": map[string]interface{}{
			"title":   title,
			"message": message,
		},
	}

	scheduleArgs := map[string]interface{}{
		"name":           name,
		"start_datetime": startTime.Format(time.RFC3339),
		"end_datetime":   endTime.Format(time.RFC3339),
		"days":           days,
		"time_type":      "console",
	}

	return c.SendScheduledCommand(body, ScheduleRecurring, scheduleArgs)
}
