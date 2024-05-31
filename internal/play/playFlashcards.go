package play

import (
	"flashCardProject/internal"
	"flashCardProject/internal/play/displayTable"
	"log"
)

func Play() {
	db := SetCardStatus()
	var flashCards []internal.FlashCard
	flashCards, err := internal.GetAllFlashcardsData(db)
	if err != nil {
		log.Fatalln(err)
	}

	displayTable.ShowFlashCardsTable(flashCards)

}
