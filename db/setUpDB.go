package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	dsn = "mhoarb:Mohammad2218@tcp(127.0.0.1:3306)/FlashCard?charset=utf8mb4&parseTime=True&loc=Local"
)

func SetUpDatabase() (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
