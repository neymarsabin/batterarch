package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
	database "github.com/neymarsabin/batterarch/models"
	"gorm.io/gorm"
)

func InitializeServer(db *gorm.DB) {
	r := gin.Default()
	batteryDetails := database.GetAllRecords(db)
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"data":    batteryDetails,
			"message": "All records fetched successfully...",
		})
	})
	r.Run(":42069")
	fmt.Println("Server started at http://localhost:42069/")
}
