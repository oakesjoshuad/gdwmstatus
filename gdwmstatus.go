package main

import (
	"os/exec"
	"time"
)

const (
	// Date Format
	DateFormat = "Mon Jan 2 15:04 MST 2006"
	// Wireless nf-mdi-wifi
	Wifi_On  = "\ufaa8"
	Wifi_Off = "\ufaa9"
	// Battery nf-mdi-battery
	Battery              = "\uf578"
	Battery_10           = "\uf579"
	Battery_20           = "\uf57a"
	Battery_30           = "\uf57b"
	Battery_40           = "\uf57c"
	Battery_50           = "\uf57d"
	Battery_60           = "\uf57e"
	Battery_70           = "\uf57f"
	Battery_80           = "\uf580"
	Battery_90           = "\uf581"
	Battery_Alert        = "\uf582"
	Battery_Charging     = "\uf583"
	Battery_Charging_100 = "\uf584"
	Battery_Charging_20  = "\uf585"
	Battery_Charging_30  = "\uf586"
	Battery_Charging_40  = "\uf587"
	Battery_Charging_60  = "\uf588"
	Battery_Charging_80  = "\uf589"
	Battery_Charging_90  = "\uf58a"
)

func main() {
	duration := 5 * time.Second
	for ; ; time.Sleep(duration) {
		if err := exec.Command("xsetroot", "-name", Status(" ")).Run(); err != nil {
			panic(err)
		}
	}
}
