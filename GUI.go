package main

import (
	"fmt"
	"strings"
)

var cRed string = "\033[31m"
var cGreen string = "\033[32m"
var cYellow string = "\033[33m"
var cBlue string = "\033[34m"
var cPurple string = "\033[35m"
var cCyan string = "\033[36m"
var cWhite string = "\033[37m"
var reset string = "\033[0m"

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
}
