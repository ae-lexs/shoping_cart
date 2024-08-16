package database

import (
	"fmt"
	"log"
	"os"

	"github.com/ae-lexs/vinyl_store/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	DBInstance *gorm.DB
	err        error
)

// Connects to the database and sets the DBInstance variable.
func SetUpDabatabse() {
	dsn := fmt.Sprintf(
		"host=postgres user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Shanghai",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		panic("Failed to connect to the database.")
	}

	log.Println("Connected to the Dabase")

	db.Logger = logger.Default.LogMode(logger.Info)

	db.AutoMigrate(&model.Album{})

	log.Println("Models migrated")

	DBInstance = db
}
