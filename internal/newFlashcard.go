package internal

import (
	database "flashCardProject/db"
	"gorm.io/gorm"
	"log"
	"log/slog"
)

var FlashCardCount = 0

func NewFlashCard(question, answer string) *gorm.DB {
	db, err := database.SetUpDatabase()
	err = db.AutoMigrate(&FlashCard{})
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	flashCard := &FlashCard{Question: question, Answer: answer}

	FlashCardCount++
	result := db.Create(flashCard)
	if result.Error != nil {
		log.Fatalf("failed to create flashCard: %v", result.Error)
	}

	slog.Info("adding a flashcard")

	return db

}
