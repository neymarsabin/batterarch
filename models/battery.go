package database

type BatteryDetails struct {
	ID            uint   `gorm:"primaryKey"`
	ModelName     string `gorm:"type:varchar(100)"`
	VoltageNow    string `gorm:"type:varchar(100)"`
	CapacityLevel string `gorm:"type:varchar(100)"`
	PowerNow      string `gorm:"type:varchar(100)"`
	EnergyNow     string `gorm:"type:varchar(100)"`
	Status        string `gorm:"type:varchar(100)"`
	CycleCount    string `gorm:"type:varchar(100)"`
	BatteryLevel  string `gorm:"type:varchar(100)"`
	RecordedAt    string `gorm:"type:varchar(100)"`
	SupplyType    string `gorm:"type:varchar(100)"`
}
