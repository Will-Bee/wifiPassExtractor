package main

import (
	"os/exec"
	"strconv"
	"strings"
)

func Getpass(SSID string) string {
	// execute the command

	wifiCounter++

	password := ""

	SSID = strings.Replace(SSID, "\r", "", -1)

	out, err := exec.Command("cmd", "/C", "netsh wlan show profile "+strconv.Quote(SSID)+" key=clear").Output()
	if err != nil {
		// fmt.Println(err)
		errorCounter++
		return cRed + "[ENCODING ERROR]" + reset
	}

	out_Array := strings.Split(string(out), "\n")

	for i := 0; i < len(out_Array); i++ {
		if strings.Contains(out_Array[i], "    Key Content            : ") {
			password = out_Array[i]
			password = strings.Replace(password, "    Key Content            : ", "", -1)
		}
	}

	successCounter++

	if password == "" {
		return cGreen + "[NO PASS]" + reset
	}

	return cCyan + password + reset
}
