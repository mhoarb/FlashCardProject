package internal

import (
	"fmt"
	"gorm.io/gorm"
)

func GetAllFlashcardsData(db *gorm.DB) ([]FlashCard, error) {
	var flashCards []FlashCard
	result := db.Find(&flashCards)
	fmt.Println(flashCards)
	return flashCards, result.Error
}
