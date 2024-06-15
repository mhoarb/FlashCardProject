package db

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
	"sync"
)

var (
	db   *gorm.DB
	once sync.Once
)

// SetUpDatabase initializes and returns a connection to the database using GORM.
// It connects to the MySQL database using the provided DSN (Data Source Name).
// If the connection is successful, it returns a pointer to the GORM DB object.
// If there is an error during the connection process, it returns a nil pointer and the error.
//
// Returns:
// - *gorm.DB: Pointer to the GORM DB object if the connection is successful.
// - error: Error object if there is an issue with the database connection.
func SetUpDatabase() (*gorm.DB, error) {
	var err error
	if err = godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	dsn := os.Getenv("DB_DSN")

	once.Do(func() {
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {

			fmt.Println("Error connecting to the database:", err)
		}
	})
	if err != nil {
		return nil, err
	}
	return db, nil
}
