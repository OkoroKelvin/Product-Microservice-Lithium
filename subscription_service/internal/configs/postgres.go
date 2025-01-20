package configs

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetUpDatabaseConnection() *gorm.DB {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", dbHost, dbUser, dbPassword, dbName, dbPort)
	db, dbErr := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if dbErr != nil {
		time.Sleep(5 * time.Second)

		log.Printf("Failed to connect to database, retrying...: %v\n", dbErr)
		return SetUpDatabaseConnection()
	}

	// RUN MIGRATIONS HERE
	// db.AutoMigrate(models.Subscription{})

	log.Println("Connected To PostgresDB!")
	return db
}

func CloseDatabaseConnection(db *gorm.DB) {
	dbInstance, err := db.DB()
	if err != nil {
		panic(fmt.Sprintf("Error Closing Database: %v", err))
	}
	log.Println("Closing Database Connection")
	dbInstance.Close()
}
