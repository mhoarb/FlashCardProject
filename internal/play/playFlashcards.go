package play

import (
	db2 "flashCardProject/db"
	"flashCardProject/internal"
	"flashCardProject/internal/play/displaytable"
	"log"
)

var flashCards []internal.FlashCard

// Play initializes the game by setting up the database connection,
// retrieving all flash cards data, displaying them in a table format,
// and allowing the user to choose and interact with the flash cards.
func Play() {
	db, err := db2.SetUpDatabase()
	if err != nil {
		log.Fatalln(err)
	}
	flashCards, err = internal.GetAllFlashcardsData(db)
	if err != nil {
		log.Fatalln(err)
	}
	displaytable.ShowFlashCardsTable(flashCards)
	ChooseQuestion()
}
