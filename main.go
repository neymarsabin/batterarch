package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	database "github.com/neymarsabin/batterarch/models"
	server "github.com/neymarsabin/batterarch/tcp"
	"github.com/neymarsabin/batterarch/visual"
	"gorm.io/gorm"
)

const sysFile = "/sys/class/power_supply/BAT0/uevent"

type BatteryInfo map[string]interface{}

func main() {
	db := database.InitializeDb()
	data := readSysFile()

	// take arugments from OS
	args := os.Args[1:]

	// save battery related details into Database
	batteryInfo := GetKeyValueDataFrom(strings.Split(string(data), "\n"))
	batteryInfo.SaveToDatabase(db)

	if len(args) > 0 {
		if args[0] == "server" || args[0] == "s" {
			server.InitializeServer(db)
		}

		if args[0] == "json" || args[0] == "j" {
			records := database.GetAllRecords(db)
			val, err := json.MarshalIndent(records, "", "  ")

			if err != nil {
				fmt.Println("Error while marshalling data: ", err)
				os.Exit(1)
			}

			fmt.Println(string(val))
		}

		if args[0] == "graph" || args[0] == "g" {
			records := database.GetAllRecords(db)
			lastChargedBatteryDetails := database.GetLastBatteryFullCharge(db)
			lastChargedTimeUnix := time.Unix(lastChargedBatteryDetails.RecordedAtUnix, 0)
			fmt.Println("Last charged time Unix: ", lastChargedTimeUnix.Format(time.RFC3339))
			fmt.Println("Last Charged Battery Level: ", lastChargedBatteryDetails.BatteryLevel)
			visual.GenerateGraph(*records)
		}
	}
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
	batteryInfo["recordedAtUnix"] = time.Now().Unix()
	newBatteryLevel, err := strconv.Atoi(batteryInfo["POWER_SUPPLY_CAPACITY"].(string))
	if err != nil {
		fmt.Println("Error while converting string to int: ", err)
		os.Exit(1)
	}
	batteryInfo["POWER_SUPPLY_CAPACITY"] = newBatteryLevel
	return batteryInfo
}

func (b BatteryInfo) SaveToDatabase(db *gorm.DB) {
	db.Create(&database.BatteryDetails{
		ModelName:      b["POWER_SUPPLY_MODEL_NAME"].(string),
		VoltageNow:     b["POWER_SUPPLY_VOLTAGE_NOW"].(string),
		CapacityLevel:  b["POWER_SUPPLY_CAPACITY_LEVEL"].(string),
		PowerNow:       b["POWER_SUPPLY_POWER_NOW"].(string),
		EnergyNow:      b["POWER_SUPPLY_ENERGY_NOW"].(string),
		Status:         b["POWER_SUPPLY_STATUS"].(string),
		CycleCount:     b["POWER_SUPPLY_CYCLE_COUNT"].(string),
		RecordedAt:     b["recordedAt"].(time.Time).String(),
		BatteryLevel:   b["POWER_SUPPLY_CAPACITY"].(int),
		SupplyType:     b["POWER_SUPPLY_TYPE"].(string),
		RecordedAtUnix: b["recordedAtUnix"].(int64),
	})
	var record database.BatteryDetails
	db.First(&record, 1)
}
