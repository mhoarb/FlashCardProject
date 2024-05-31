package internal

import (
	"gorm.io/gorm"
)

func GetAllFlashcardsData(db *gorm.DB) ([]FlashCard, error) {
	var flashCards []FlashCard
	result := db.Find(&flashCards)
	return flashCards, result.Error
}
