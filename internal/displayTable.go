package internal

import (
	"github.com/olekukonko/tablewriter"
	"os"
)

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
