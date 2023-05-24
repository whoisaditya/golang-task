package initializers

import (
	"os"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// var DB *gorm.DB

func ConnectToDB(){
	var err error
	dsn := os.Getenv("DB_URL")
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	// Use DB
	_ = DB

	if err != nil {
		panic("Failed to connect DB")
	} else {
		fmt.Println("Connected to DB")
	}
}