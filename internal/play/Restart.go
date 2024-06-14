package play

import (
	db2 "flashCardProject/db"
	"log"
	"log/slog"
)

// Restart resets the game by updating the status of all flash cards to NULL in the database.
// It initializes a database connection, performs an update operation on the 'status' field
// of the 'flash_cards' table, logs success or failure, and handles errors gracefully.
func Restart() {
	db, err := db2.SetUpDatabase()
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}
	_ = db.Exec("UPDATE flash_cards SET status = NULL")
	if err != nil {
		slog.Error("There is a problem in resetting the game")
	}
	slog.Info("The game was reset successfully")
}
