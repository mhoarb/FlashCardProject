package internal

import (
	"github.com/olekukonko/tablewriter"
	"os"
)

func DisplayTable() {
	table := tablewriter.NewWriter(os.Stdout)

	table.SetHeader([]string{"QUESTION", "ANSWER"})

	table.Render()

}
