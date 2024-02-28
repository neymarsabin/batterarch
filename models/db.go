package database

import (
	"fmt"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const dbFile = "batterarch.db"

func InitializeDb() *gorm.DB {
	// Create database directory if it doesn't exist: $HOME/.config/batterarch/
	homeDir := getHomeDirectory()
	databasePath := fmt.Sprintf("%s/.config/batterarch/", homeDir)
	if _, err := os.Stat(databasePath); os.IsNotExist(err) {
		os.MkdirAll(databasePath, 0700)
	}

	fullDatabasePath := fmt.Sprintf("%s%s", databasePath, dbFile)
	db, err := gorm.Open(sqlite.Open(fullDatabasePath), &gorm.Config{})
	if err != nil {
		fmt.Println("Error while initializing database: ", err)
		os.Exit(1)
	}

	db.AutoMigrate(&BatteryDetails{})
	return db
}

func getHomeDirectory() string {
	home, _ := os.UserHomeDir()
	return home
}
