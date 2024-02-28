package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	database "github.com/neymarsabin/batterarch/models"
	"gorm.io/gorm"
)

const sysFile = "/sys/class/power_supply/BAT0/uevent"
const logPath = "/home/neymarsabin/.config/batterarch/access.log"

type BatteryInfo map[string]interface{}

func main() {
	db := database.InitializeDb()
	data := readSysFile()
	batteryDetailsTemp := strings.Split(string(data), "\n")
	batteryDetails := GetKeyValueDataFrom(batteryDetailsTemp)
	batteryDetails.SaveToDatabase(db)
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
	savedData := db.Create(&database.BatteryDetails{
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
	fmt.Println("Saved data: ", savedData)
	var record database.BatteryDetails
	db.First(&record, 1)
}
