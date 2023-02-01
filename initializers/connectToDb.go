package initializers

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDb() {
	var err error
	URL_DB := "host=localhost user=farhan_binar password=farhan_322 dbname=go_api port=5432  sslmode=disable"
	dsn := URL_DB
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed connect to db")
	}
}
