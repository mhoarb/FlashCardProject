package displayTable

import (
	"flashCardProject/internal"
	"github.com/olekukonko/tablewriter"
	"log/slog"
	"os"
)

func ShowFlashCardsTable(flashCards []internal.FlashCard) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"QUESTION", "STATUS: Correct/False/not Answer"})

	for _, cards := range flashCards {
		row := []string{
			cards.Question,
			cards.Status,
		}
		table.Append(row)
	}

	table.Render()
	slog.Info("%%%%%")
}
