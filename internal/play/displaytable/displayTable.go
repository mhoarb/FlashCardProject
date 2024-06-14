package displaytable

import (
	"flashCardProject/internal"
	"github.com/olekukonko/tablewriter"
	"os"
	"strconv"
)

// ShowFlashCardsTable displays a table of flash cards with their index, question, and status.
// It uses the tablewriter package to format and print the table to the standard output.
// After displaying the table, it calls the Progress function to update the game progress.
//
// Parameters:
//   - flashCards: A slice of FlashCard structs containing the data to display in the table.
func ShowFlashCardsTable(flashCards []internal.FlashCard) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"INDEX", "QUESTION", "STATUS: Correct/False/not Answer"})

	for _, cards := range flashCards {
		row := []string{
			strconv.Itoa(int(cards.ID)),
			cards.Question,
			cards.Status,
		}
		table.Append(row)
	}

	table.Render()
	Progress()
}
