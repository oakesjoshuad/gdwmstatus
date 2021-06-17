package main

import (
	"os/exec"
	"time"
)

const (
	// Date Format
	DateFormat = "Mon Jan 2 15:04 MST 2006"
	// Wireless nf-mdi-wifi
	WifiConnected    = "\ufaa8"
	WifiDisconnected = "\ufaa9"
	// LAN nf-mdi-lan
	LANConnected    = "\uf817"
	LANDisconnected = "\uf818"
	// Battery nf-mdi-battery
	Battery            = "\uf578"
	Battery10          = "\uf579"
	Battery20          = "\uf57a"
	Battery30          = "\uf57b"
	Battery40          = "\uf57c"
	Battery50          = "\uf57d"
	Battery60          = "\uf57e"
	Battery70          = "\uf57f"
	Battery80          = "\uf580"
	Battery90          = "\uf581"
	BatteryAlert       = "\uf582"
	BatteryCharging    = "\uf583"
	BatteryCharging100 = "\uf584"
	BatteryCharging20  = "\uf585"
	BatteryCharging30  = "\uf586"
	BatteryCharging40  = "\uf587"
	BatteryCharging60  = "\uf588"
	BatteryCharging80  = "\uf589"
	BatteryCharging90  = "\uf58a"
)

func main() {
	duration := 5 * time.Second
	for ; ; time.Sleep(duration) {
		if err := exec.Command("xsetroot", "-name", Status(" ")).Run(); err != nil {
			panic(err)
		}
	}
}
