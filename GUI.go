package main

import (
	"fmt"
	"strings"
)

func Welcome() {
	fmt.Println(cYellow + "\nWifi passwords saved in this computer \n" + reset)
	fmt.Println(cYellow + "SSID                               |   Password\n" + reset)

}

func PrintWifiDetails(SSID string, password string) {
	SSID = strings.Replace(SSID, "\r", "", -1)

	SSIDLength := len(SSID)

	SSIDLength = 35 - SSIDLength

	if SSIDLength < 0 {
		SSIDLength = 0
	}

	tempSpace := strings.Repeat(" ", SSIDLength)

	// dataOut := ""

	if password == cGreen+"[NO PASS]"+reset {
		cSSID = cGreen
	} else if password == cRed+"[ENCODING ERROR]"+reset {
		cSSID = cRed
		failedSSID = append(failedSSID, SSID)
	} else {
		cSSID = cCyan
	}

	fmt.Println(cSSID + SSID + tempSpace + "|   " + password)
}

func Outro() {
	fmt.Println("\n\n"+cGreen+"SUCCESS: ", reset, successCounter, "/", wifiCounter)
	fmt.Println(cRed+"ERRORS:  ", reset, errorCounter, cYellow, "\n\nFailed SSIDs:\n", reset)

	for i := 0; i < len(failedSSID); i++ {
		fmt.Println(cRed + failedSSID[i] + reset)
	}

	fmt.Scanln()
	fmt.Println("exiting...")
}
