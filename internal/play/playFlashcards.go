package play

import (
	db2 "flashCardProject/db"
	"flashCardProject/internal"
	"flashCardProject/internal/play/displayTable"
	"log"
)

func Play() {
	db, err := db2.SetUpDatabase()
	var flashCards []internal.FlashCard
	flashCards, err = internal.GetAllFlashcardsData(db)
	if err != nil {
		log.Fatalln(err)
	}
	displayTable.ShowFlashCardsTable(flashCards)
	ChooseQuestion()
}
