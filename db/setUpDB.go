package db

import (
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

// LoadEnv loads environment variables from a .env file located in the current directory.
// It returns an error if the .env file cannot be loaded or parsed successfully.
func LoadEnv() error {
	return godotenv.Load()
}

// SetUpDatabase initializes and returns a connection to the database using GORM.
// It connects to the MySQL database using the provided DSN (Data Source Name).
// If the connection is successful, it returns a pointer to the GORM DB object.
// If there is an error during the connection process, it returns a nil pointer and the error.
//
// Returns:
// - *gorm.DB: Pointer to the GORM DB object if the connection is successful.
// - error: Error object if there is an issue with the database connection.
func SetUpDatabase() (*gorm.DB, error) {
	dsn := os.Getenv("DSN")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
