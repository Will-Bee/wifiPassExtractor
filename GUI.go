package main

import (
	"fmt"
	"strings"
)

func Welcome() {
	fmt.Println(cGreen + "Wifi passwords saved in this computer" + reset)

}

func PrintWifiDetails(SSID string, password string) {
	SSID = strings.Replace(SSID, "\r", "", -1)
	fmt.Println(cCyan + "SSID: " + cYellow + SSID)
	fmt.Println(cCyan + "Password: " + reset + password)
	fmt.Println("--------------------")
}

func Outro() {
	fmt.Println(cGreen+"SUCCESS: ", reset, successCounter, "/", wifiCounter)
	fmt.Println(cRed+"ERRORS: ", reset, errorCounter)
	fmt.Scanln()
	fmt.Println("exiting...")
}
