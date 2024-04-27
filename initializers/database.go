package initializers

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error

	fmt.Println("Connecting to database...")
	dsn := os.Getenv("DB_URI") 
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	} else {
		fmt.Println("Connected to database💡")
	}
}