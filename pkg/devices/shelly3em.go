// shelly3em is a package for collecting the status of
// shelly3em devices https://shelly-api-docs.shelly.cloud/gen1/#shelly-3em
package devices

// TODO reset totals each month ?
// See https://shelly-api-docs.shelly.cloud/gen1/#shelly-3em-emeter-index
// reset_totals

type Wifi struct {
	Connected bool   `json:"connected"`
	SSID      string `json:"ssid"`
	IP        string `json:"ip"`
	RSSI      int    `json:"rssi"`
}

// See also https://shelly-api-docs.shelly.cloud/gen1/#ota
type UpdateStatus string

const (
	UpdateIdle    UpdateStatus = "idle"
	UpdatePending UpdateStatus = "pending"
	UpdateRunning UpdateStatus = "updating"
)

// See also https://shelly-api-docs.shelly.cloud/gen1/#ota
type OTA struct {
	Status          UpdateStatus `json:"status"`
	UpdateAvailable bool         `json:"has_update"`
	NewVersion      string       `json:"new_version"`
	CurrentVersion  string       `json:"old_version"`
}

// https://shelly-api-docs.shelly.cloud/gen1/#status
type NodeStatus struct {
	// MAC address of the device
	MAC string `json:"mac"`
	// Whether sysflash is mounted
	FSMounted bool `json:"fs_mounted"`
	// Total amount of system memory in bytes
	RamTotal int `json:"ram_total"`
	// Available amount of system memory in bytes
	RamFree int `json:"ram_free"`
	// Total amount of the file system in bytes
	FilesystemSize int `json:"fs_size"`
	// Available amount of the file system in bytes
	FilesystemFree int `json:"fs_free"`
	// Seconds elapsed since boot
	Uptime int `json:"uptime"`
}

// See also https://shelly-api-docs.shelly.cloud/gen1/#shelly-3em-emeter-index
type Emeter struct {
	// Instantaneous power, Watts
	Power float64 `json:"power"`
	// Power factor (dimensionless)
	PowerFactor float64 `json:"pf"`
	// Current, A
	Current float64 `json:"current"`
	// RMS voltage, Volts
	Voltage float64 `json:"voltage"`
	// Whether the associated meter is functioning properly
	Valid bool `json:"is_valid"`
	// Total consumed energy, Wh
	Total float64 `json:"total"`
	// Total returned energy, Wh
	TotalReturned float64 `json:"total_returned"`
}

type Shelly3EM struct {
	NodeStatus
	Wifi    `json:"wifi_sta"`
	OTA     `json:"update"`
	Emeters []Emeter `json:"emeters"`
	// Sum of the power of the three channels, Watts
	TotalPower float64 `json:"total_power"`
}
