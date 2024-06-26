package displaytable

import (
	db2 "flashCardProject/db"
	"flashCardProject/internal"
	"fmt"
)

// Variables to track statistics related to flash cards.
var (
	AllQuestion      int64
	CorrectAnswer    int64
	NotCorrectAnswer int64
)

// GameStatistics calculates and prints various statistics related to flash cards stored in the database.
// It retrieves data from the database, including the total number of questions, the number of questions
// answered correctly and incorrectly, and calculates percentages for answered questions and correct answers.
// Prints the results to the standard output.
func GameStatistics() {
	db, err := db2.SetUpDatabase()
	if err != nil {
		panic(err)
	}
	var flashCard *internal.FlashCard
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

// Progress calculates the progress of correct answers among all questions
// and displays it as a percentage in a progress bar.
// It retrieves data from the database to count total questions and correct answers,
// then uses the calculated percentage to update the progress bar display.
func Progress() {
	db, err := db2.SetUpDatabase()
	if err != nil {
		panic(err)
	}
	var flashCard *internal.FlashCard
	if err := db.Model(&flashCard).Count(&AllQuestion).Error; err != nil {
		return
	}
	if err := db.Model(&flashCard).Where("status = ?", "correct").Count(&CorrectAnswer).Error; err != nil {
		fmt.Println("ds")
	}
	PercentageOfQuestionsWithCorrectAnswers := float64(CorrectAnswer) / float64(AllQuestion) * 100
	printProgressBar(PercentageOfQuestionsWithCorrectAnswers)
}

func printProgressBar(percentage float64) {
	barWidth := 50 // Width of the progress bar
	pos := int(percentage / 100 * float64(barWidth))
	fmt.Println()
	fmt.Println()

	fmt.Print("[")
	for i := 0; i < barWidth; i++ {
		if i < pos {
			fmt.Print("=")
		} else {
			fmt.Print(" ")
		}
	}
	fmt.Printf("] %.2f%%\n\n\n", percentage)
}
