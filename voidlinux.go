package main

import (
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

const (
	path_power   = "/sys/class/power_supply"
	path_network = "/sys/class/net"
	path_proc    = "/proc"

	ep_battery      = "/BAT0"
	ep_poweradapter = "/AC"
	ep_wifi         = "/wlp3s0"

	fd_battery_capacity    = path_power + ep_battery + "/capacity"    // 0 - 100 || Full
	fd_poweradapter_status = path_power + ep_poweradapter + "/online" // 1
	fd_wifi_status         = path_network + ep_wifi + "/operstate"    // up
	fd_loadavg             = path_proc + "/loadavg"
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
	if content, err := parseFile(fd_poweradapter_status); err != nil {
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
	if content, err := parseFile(fd_wifi_status); err != nil {
		return ""
	} else {
		switch content {
		case "up":
			return Wifi_On
		case "down":
			return Wifi_Off
		default:
			return Wifi_Off
		}
	}
}

// batteryCapacity returns the current battery capacity in a range between 0-100, otherwise an error
func batteryCapacity() string {
	content, err := parseFile(fd_battery_capacity)
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
			return Battery_Alert
		case capacity < 20:
			return Battery_Charging
		case capacity < 30:
			return Battery_Charging_20
		case capacity < 40:
			return Battery_Charging_30
		case capacity < 60:
			return Battery_Charging_40
		case capacity < 80:
			return Battery_Charging_60
		case capacity < 90:
			return Battery_Charging_80
		default:
			return Battery_Charging
		}
	} else {
		switch {
		case capacity < 10:
			return Battery_Alert
		case capacity < 20:
			return Battery_10
		case capacity < 30:
			return Battery_20
		case capacity < 40:
			return Battery_30
		case capacity < 50:
			return Battery_40
		case capacity < 60:
			return Battery_50
		case capacity < 70:
			return Battery_60
		case capacity < 80:
			return Battery_70
		case capacity < 90:
			return Battery_80
		case capacity < 100:
			return Battery_90
		default:
			return Battery
		}
	}
}

// loadAvg
func loadAvg(fields uint, sep string) string {
	if content, err := parseFile(fd_loadavg); err != nil {
		return ""
	} else {
		splits := strings.Split(content, " ")
		return strings.Join(splits[:fields], sep)
	}
}

func Status(sep string) string {
	tab := "  "
	status := []string{
		tab,
		loadAvg(4, " "),
		wifiStatus(),
		batteryCapacity(),
		time.Now().Format(DateFormat),
		tab,
	}
	return strings.Join(status, sep)
}
