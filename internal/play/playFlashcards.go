package play

import (
	db2 "flashCardProject/db"
	"flashCardProject/internal"
	"flashCardProject/internal/play/displayTable"
	"log"
)

var flashCards []internal.FlashCard

func Play() {
	db, err := db2.SetUpDatabase()
	if err != nil {
		log.Fatalln(err)
	}
	flashCards, err = internal.GetAllFlashcardsData(db)
	if err != nil {
		log.Fatalln(err)
	}
	displayTable.ShowFlashCardsTable(flashCards)
	ChooseQuestion()
}
