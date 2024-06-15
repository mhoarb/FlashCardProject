package internal

import (
	"gorm.io/gorm"
)

// GetAllFlashcardsData retrieves all flash cards data from the database using the provided GORM DB connection.
// It returns a slice of FlashCard structs and any error encountered during the database query.
//
// Parameters:
//   - db: Pointer to a GORM DB instance initialized with a database connection.
//
// Returns:
//   - []FlashCard: A slice of FlashCard structs representing all flash cards in the database.
//   - error: An error if the database query fails.

func GetAllFlashcardsData(db *gorm.DB) ([]FlashCard, error) {
	var flashCards []FlashCard
	result := db.Find(&flashCards)
	return flashCards, result.Error
}
