package internal

import (
	"github.com/olekukonko/tablewriter"
	"os"
)

// DisplayTable renders a table to stdout displaying flash cards' questions and answers.
// It uses the tablewriter package to format and print the table.
//
// Parameters:
//   - flashCards: A slice of FlashCard structs containing questions and answers to display.
func DisplayTable(flashCards []FlashCard) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"QUESTION", "ANSWER"})

	for _, cards := range flashCards {
		row := []string{
			cards.Question,
			cards.Answer,
		}
		table.Append(row)
	}

	table.Render()

}
