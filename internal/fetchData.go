package internal

import (
	"gorm.io/gorm"
)

func FetchData(db *gorm.DB) ([]FlashCard, error) {
	var flashCards []FlashCard
	resault := db.Find(flashCards)
	if resault.Error != nil {
		return nil, resault.Error
	}
	return flashCards, nil
}
