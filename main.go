package main

import (
	"encoding/json"
	"fmt"
	"os"
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
	return batteryInfo
}

func (b BatteryInfo) SaveToDatabase(db *gorm.DB) {
	db.Create(&database.BatteryDetails{
		ModelName:     b["POWER_SUPPLY_MODEL_NAME"].(string),
		VoltageNow:    b["POWER_SUPPLY_VOLTAGE_NOW"].(string),
		CapacityLevel: b["POWER_SUPPLY_CAPACITY_LEVEL"].(string),
		PowerNow:      b["POWER_SUPPLY_POWER_NOW"].(string),
		EnergyNow:     b["POWER_SUPPLY_ENERGY_NOW"].(string),
		Status:        b["POWER_SUPPLY_STATUS"].(string),
		CycleCount:    b["POWER_SUPPLY_CYCLE_COUNT"].(string),
		BatteryLevel:  b["POWER_SUPPLY_CAPACITY"].(string),
		RecordedAt:    b["recordedAt"].(time.Time).String(),
		SupplyType:    b["POWER_SUPPLY_TYPE"].(string),
	})
	var record database.BatteryDetails
	db.First(&record, 1)
}
