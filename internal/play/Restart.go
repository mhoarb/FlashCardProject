package play

import (
	db2 "flashCardProject/db"
	"log"
	"log/slog"
)

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
