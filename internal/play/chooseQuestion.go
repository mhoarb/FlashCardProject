package play

import (
	"bufio"
	db2 "flashCardProject/db"
	"flashCardProject/internal"
	"fmt"
	"log"
	"log/slog"
	"os"
)

var AllQuestion int64
var flashCard *internal.FlashCard
var CorrectAnswer int64
var NotCorrectAnswer int64

func ChooseQuestion() {
	db, err := db2.SetUpDatabase()
	if err != nil {
		panic(err)
	}

	slog.Info("please Choose the index of your question you want ?")
	reader := bufio.NewReader(os.Stdin)
	questionID, _ := reader.ReadString('\n')
	questionID = questionID[:len(questionID)-1]

	if err := db.First(&flashCard, questionID).Error; err != nil {
		slog.Error("Failed to retrieve question:", err)
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
func GameStatistics() {
	db, err := db2.SetUpDatabase()
	if err != nil {
		panic(err)
	}

	if err := db.Model(&flashCard).Count(&AllQuestion).Error; err != nil {
		return
	}
	fmt.Println("The total number of questions is equal to:", AllQuestion)

	if err := db.Model(&flashCard).Where("status = ?", "correct").Count(&CorrectAnswer).Error; err != nil {
		fmt.Println("ds")
	}
	if err := db.Model(&flashCard).Where("status = ?", "Not correct").Count(&NotCorrectAnswer).Error; err != nil {
		return
	}
	percentageOfQuestionsAnswered := float64(CorrectAnswer+NotCorrectAnswer) / float64(AllQuestion) * 100
	fmt.Println("The Percentage of questions answered is equal to:", percentageOfQuestionsAnswered, "%")
	PercentageOfQuestionsWithCorrectAnswers := float64(CorrectAnswer) / float64(AllQuestion) * 100
	fmt.Println("The Percentage of questions with correct answers is equal to:", PercentageOfQuestionsWithCorrectAnswers, "%")

}
