package play

import (
	db2 "flashCardProject/db"
	"flashCardProject/internal"
	"gorm.io/gorm"
	"log"
)

var status string

const (
	correct     = "correct"
	notCorrect  = "false"
	notAnswered = "not Answered"
)

func SetCardStatus() *gorm.DB {
	db, err := db2.SetUpDatabase()
	if err != nil {
		panic(err)
	}

	result := db.Model(&internal.FlashCard{}).Where("1 = 1").Updates(map[string]interface{}{"Status": notAnswered})
	if result.Error != nil {
		log.Fatalf("failed to update flashCard status: %v", result.Error)
	}
	return db
}
