package play

import (
	"bufio"
	db2 "flashCardProject/db"
	"flashCardProject/internal"
	"log"
	"log/slog"
	"os"
)

type RequestQuestionID struct {
	ID int
}

func ChooseQuestion() {
	db, err := db2.SetUpDatabase()
	if err != nil {
		panic(err)
	}

	slog.Info("please Choose the index of your question you want ?")
	reader := bufio.NewReader(os.Stdin)
	questionID, _ := reader.ReadString('\n')
	questionID = questionID[:len(questionID)-1]

	var flashCard *internal.FlashCard
	if err := db.First(&flashCard, questionID).Error; err != nil {
		slog.Error("Failed to retrieve question:", err)
	}
	slog.Info("please write the Answer of this question")
	answer, _ := reader.ReadString('\n')
	answer = answer[:len(answer)-1]

	if answer == flashCard.Answer {
		flashCard.Status = "correct"
		slog.Info("your answer is correct")
		if err := db.Save(&flashCard).Error; err != nil {
			log.Fatalf("Error updating question: %v", err)
		}
	} else {
		flashCard.Status = "Not correct"
		slog.Info("your answer isn''t correct, please try again`")
		if err := db.Save(&flashCard).Error; err != nil {
			log.Fatalf("Error updating question: %v", err)
		}
		Play()
	}
}
