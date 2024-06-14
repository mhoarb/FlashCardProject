package internal

import (
	database "flashCardProject/db"
	"gorm.io/gorm"
	"log"
	"log/slog"
)

// FlashCardCount represents the current count of flash cards.
var FlashCardCount = 0

// NewFlashCard creates and saves a new flash card with the given question and answer to the database.
// It initializes a database connection, performs auto migration for FlashCard model,
// increments the FlashCardCount, and logs operations. Returns the database connection.
//
// Parameters:
//   - question: The question text for the new flash card.
//   - answer: The answer text for the new flash card.
//
// Returns:
//   - *gorm.DB: Pointer to the GORM DB instance after creating the flash card.
func NewFlashCard(question, answer string) *gorm.DB {
	db, err := database.SetUpDatabase()
	if err != nil {
		log.Fatalf("failed to setup database: %v", err)
	}
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
