package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

const sysFile = "/sys/class/power_supply/BAT0/uevent"

type BatteryInfo map[string]interface{}

func main() {
	data := readSysFile()
	temp := strings.Split(string(data), "\n")
	batteryDetails := GetKeyValueDataFrom(temp)
	batteryDetails.SaveToJsonFile()
}

func readSysFile() []byte {
	data, err := os.ReadFile(sysFile)
	if err != nil {
		os.Exit(1)
	}

	return data
}

func GetKeyValueDataFrom(data []string) BatteryInfo {
	batteryInfo := make(BatteryInfo)
	for _, line := range data {
		if line != "" {
			pair := strings.Split(line, "=")
			batteryInfo[pair[0]] = pair[1]
		}
	}
	batteryInfo["recordedAt"] = time.Now()
	return batteryInfo
}

func (b BatteryInfo) SaveToJsonFile() {
	jsonDetails, err := json.Marshal(b)
	if err != nil {
		fmt.Println("Error while converting json details: ", err)
		os.Exit(1)
	}
	fmt.Println(string(jsonDetails))
}
