package main

// create a basic app

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var errorCounter int = 0
var wifiCounter int = 0
var successCounter int = 0

var filename_str string = ""

var colours bool = false

var cRed string = "\033[31m"
var cGreen string = "\033[32m"
var cYellow string = "\033[33m"
var cBlue string = "\033[34m"
var cPurple string = "\033[35m"
var cCyan string = "\033[36m"
var cWhite string = "\033[37m"
var reset string = "\033[0m"

func main() {

	args := os.Args[1:]

	if len(args) > 0 {

		if "-c" == args[0] {
			colours = true
		}
	}

	if colours == false {
		cRed = ""
		cGreen = ""
		cYellow = ""
		cBlue = ""
		cPurple = ""
		cCyan = ""
		cWhite = ""
		reset = ""
	}

	SSID_Raw := getSSID_Raw()
	// fmt.Println(SSID_Raw)

	SSID_RawArray := getSSID_rawArray(SSID_Raw)

	SSID_Array := []string{}

	passwd := ""

	for i := 0; i < len(SSID_RawArray); i++ {
		SSID_Array = append(SSID_Array, getSSID(SSID_RawArray[i]))
	}

	Welcome()

	for i := 0; i < len(SSID_Array); i++ {
		passwd = Getpass(SSID_Array[i])
		// passwd = "PASSTEST"
		PrintWifiDetails(SSID_Array[i], passwd)

	}

	Outro() // :D
}

//
// The SSID part
//

// Returns a SSID name
func getSSID(SSID_Raw string) string {
	// get the SSID from the raw data

	SSID := ""

	if strings.Contains(SSID_Raw, "All User Profile") {
		SSID = SSID_Raw
		SSID = strings.Replace(SSID, "    All User Profile     : ", "", -1)
	}

	return SSID
}

// Returns array of lines that contains ssid name
func getSSID_rawArray(SSID_Raw string) []string {
	// get the SSID from the raw data

	// split the string into an array line by line
	SSID_Raw_Array := strings.Split(SSID_Raw, "\n")

	SSID_Array := []string{}

	for i := 0; i < len(SSID_Raw_Array); i++ {
		if strings.Contains(SSID_Raw_Array[i], "All User Profile") {
			SSID_Array = append(SSID_Array, SSID_Raw_Array[i])
			// fmt.Println(SSID_Raw_Array[i])
		}

	}
	return SSID_Array
}

// Returns a raw string
func getSSID_Raw() string {
	cmd := exec.Command("cmd", "/C", "netsh wlan show profiles")
	out, err := cmd.Output()

	if err != nil {
		fmt.Println(err)
	}

	return string(out)
}
