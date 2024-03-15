package database

import (
	"gorm.io/gorm"
)

type BatteryDetails struct {
	ID             uint   `gorm:"primaryKey"`
	ModelName      string `gorm:"type:varchar(100)"`
	VoltageNow     string `gorm:"type:varchar(100)"`
	CapacityLevel  string `gorm:"type:varchar(100)"`
	PowerNow       string `gorm:"type:varchar(100)"`
	EnergyNow      string `gorm:"type:varchar(100)"`
	Status         string `gorm:"type:varchar(100)"`
	CycleCount     string `gorm:"type:varchar(100)"`
	BatteryLevel   int
	RecordedAt     string `gorm:"type:varchar(100)"`
	SupplyType     string `gorm:"type:varchar(100)"`
	RecordedAtUnix int64  `gorm:"type:bigint"`
}

func GetAllRecords(db *gorm.DB) *[]BatteryDetails {
	var records []BatteryDetails
	db.Where("recorded_at_unix is not null").Order("battery_level desc").Find(&records)

	return &records
}

func GetLastBatteryFullCharge(db *gorm.DB) *BatteryDetails {
	var record BatteryDetails
	db.Where("recorded_at_unix is not null").Where("status == 'Charging'").Order("battery_level desc").First(&record)

	return &record
}
