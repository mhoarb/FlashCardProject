package play

import (
	"bufio"
	db2 "flashCardProject/db"
	"flashCardProject/internal"
	"log"
	"log/slog"
	"os"
	"strconv"
)

// ChooseQuestion allows the user to select a flash card by its ID, answer it, and update its status accordingly.
// It interacts with the user via stdin for input, retrieves the flash card from the database based on the chosen ID,
// prompts the user to provide an answer, checks the correctness of the answer, updates the flash card's status,
// and recursively calls the Play function for further interaction.
func ChooseQuestion() {
	db, err := db2.SetUpDatabase()
	if err != nil {
		panic(err)
	}
	slog.Info("please Choose the index of your question you want ?")
	reader := bufio.NewReader(os.Stdin)
	questionIDStr, err := reader.ReadString('\n')
	if err != nil {
		slog.Error("input error", "Failed to read input:", err)
	}
	questionIDStr = questionIDStr[:len(questionIDStr)-1]
	questionID, err := strconv.Atoi(questionIDStr) // Convert to int
	if err != nil {
		slog.Error("invalid Question ID", "Invalid question ID format:", err)
		return
	}
	var flashCard internal.FlashCard
	if err := db.First(&flashCard, questionID).Error; err != nil {
		slog.Error("error", "Failed to retrieve question:", err)
	}

	if flashCard.Status == "correct" {

		slog.Error("this Question is already correct and you answer this")
		slog.Info("please choose another question")
		Play()
	} else {

		slog.Info("please write the Answer of this question")
		answer, _ := reader.ReadString('\n')
		answer = answer[:len(answer)-1]

		if answer == flashCard.Answer {
			flashCard.Status = "correct"
			slog.Info("your answer is correct")
			if err := db.Save(&flashCard).Error; err != nil {
				log.Fatalf("Error updating question: %v", err)
			}
			Play()
		} else {
			flashCard.Status = "Not correct"
			slog.Info("your answer isn''t correct, please try again`")
			if err := db.Save(&flashCard).Error; err != nil {
				log.Fatalf("Error updating question: %v", err)
			}
			Play()
		}
	}

}
