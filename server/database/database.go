package database

import (
	"fmt"

	"github.com/gookit/config/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DBConn *gorm.DB

func InitDatabase() {

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s",
		config.String("db_host"),
		config.String("db_user"),
		config.String("db_pass"),
		config.String("db_name"),
		config.String("db_port", "5432"),
		config.String("time_zone"),
	)

	var err error
	DBConn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
	fmt.Println("Database Connected")
}
