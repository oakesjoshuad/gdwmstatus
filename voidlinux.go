package main

import (
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

const (
	pathPower   = "/sys/class/power_supply"
	pathNetwork = "/sys/class/net"
	pathProc    = "/proc"

	epBattery      = "/BAT0"
	epPoweradapter = "/AC"
	epWifi         = "/wlp3s0"
	epLAN          = "/eno1"

	fdBatteryCapacity    = pathPower + epBattery + "/capacity"    // 0 - 100 || Full
	fdPoweradapterStatus = pathPower + epPoweradapter + "/online" // 1
	fdWifiStatus         = pathNetwork + epWifi + "/operstate"    // up
	fdLANStatus          = pathNetwork + epLAN + "/operstate"     // up
	fdLoadavg            = pathProc + "/loadavg"
)

// parseFile reads a file assumed to be one line terminated by a newline character.
// returns the string content stripped of the newline or an empty string with an error
func parseFile(path string) (string, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return strings.Split(string(content), "\n")[0], nil
}

// powerAdapterStatus returns true if the power adapter is connected, otherwise false or error
func powerAdapterStatus() bool {
	if content, err := parseFile(fdPoweradapterStatus); err != nil {
		return false
	} else {
		if status, err := strconv.Atoi(content); err != nil {
			return false
		} else {
			switch status {
			case 1:
				return true
			case 0:
				return false
			default:
				return false
			}
		}
	}
}

// wifiStatus returns true if a connection is detected, otherwise false or error
func wifiStatus() string {
	if content, err := parseFile(fdWifiStatus); err != nil {
		return ""
	} else {
		switch content {
		case "up":
			return WifiConnected
		default:
			return WifiDisconnected
		}
	}
}

func lanStatus() string {
	if content, err := parseFile(fdLANStatus); err != nil {
		return ""
	} else {
		switch content {
		case "up":
			return LANConnected
		default:
			return LANDisconnected
		}
	}
}

// batteryCapacity returns the current battery capacity in a range between 0-100, otherwise an error
func batteryCapacity() string {
	content, err := parseFile(fdBatteryCapacity)
	if err != nil {
		return ""
	}
	capacity, err := strconv.Atoi(content)
	if err != nil {
		return ""
	}
	present := powerAdapterStatus()
	if present {
		switch {
		case capacity < 10:
			return BatteryAlert
		case capacity < 20:
			return BatteryCharging
		case capacity < 30:
			return BatteryCharging20
		case capacity < 40:
			return BatteryCharging30
		case capacity < 60:
			return BatteryCharging40
		case capacity < 80:
			return BatteryCharging60
		case capacity < 90:
			return BatteryCharging80
		default:
			return BatteryCharging
		}
	} else {
		switch {
		case capacity < 10:
			return BatteryAlert
		case capacity < 20:
			return Battery10
		case capacity < 30:
			return Battery20
		case capacity < 40:
			return Battery30
		case capacity < 50:
			return Battery40
		case capacity < 60:
			return Battery50
		case capacity < 70:
			return Battery60
		case capacity < 80:
			return Battery70
		case capacity < 90:
			return Battery80
		case capacity < 100:
			return Battery90
		default:
			return Battery
		}
	}
}

// loadAvg
func loadAvg(fields uint, sep string) string {
	if content, err := parseFile(fdLoadavg); err != nil {
		return ""
	} else {
		splits := strings.Split(content, " ")
		return strings.Join(splits[:fields], sep)
	}
}

func Status(sep string) string {
	tab := "  "
	status := []string{tab}
	if la := loadAvg(4, " "); la != "" {
		status = append(status, la)
	}
	if lan := lanStatus(); lan != "" {
		status = append(status, lan)
	}
	if wifi := wifiStatus(); wifi != "" {
		status = append(status, wifi)
	}
	if pa := powerAdapterStatus(); pa {
		status = append(status, batteryCapacity())
	}
	status = append(status, time.Now().Format(DateFormat))
	status = append(status, tab)
	return strings.Join(status, sep)
}
