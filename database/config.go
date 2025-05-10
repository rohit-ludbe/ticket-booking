package database

import (
	"fmt"
	"log"
	"os"

	"github.com/rohit-ludbe/ticket-booking-server-v1/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Dbinstance struct {
	Db *gorm.DB
}

var DB Dbinstance

func ConnectDb() {

	dsn := fmt.Sprintf("host=localhost user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Kolkata", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
		os.Exit(2)
	}

	db.Logger = logger.Default.LogMode(logger.Info)

	// auto migration
	log.Println("running migration")
	db.AutoMigrate(&models.User{}, &models.Event{}, &models.Ticket{})

	DB = Dbinstance{
		Db: db,
	}

}
