package internal

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"log/slog"
)

const (
	dsn = "mhoarb:Mohammad2218@tcp(127.0.0.1:3306)/FlashCard?charset=utf8mb4&parseTime=True&loc=Local"
)

var db *gorm.DB

func NewFlashCard(question, answer string) *gorm.DB {
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	err = db.AutoMigrate(&FlashCard{})
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	flashCard := &FlashCard{Question: question, Answer: answer}

	result := db.Create(flashCard)
	if result.Error != nil {
		log.Fatalf("failed to create flashCard: %v", result.Error)
	}

	slog.Info("adding a flashcard")

	return db

}
