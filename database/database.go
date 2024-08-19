package database

import (
	"fmt"
	"log"
	"os"

	"github.com/ae-lexs/vinyl_store/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	err error
)

// Connects to the database and returns the database instance.
func SetUp() *gorm.DB {
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

	log.Println("Connected to the Database")

	db.AutoMigrate(&entity.Album{})

	log.Println("Models migrated")

	return db
}
