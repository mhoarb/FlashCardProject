// Package cmd provides the commands for the CLI application.
package cmd

import (
	"flashCardProject/internal/play/displayTable"
	"github.com/spf13/cobra"
)

// StatisticsCmd represents the Statistics command
var StatisticsCmd = &cobra.Command{
	Use:   "Statistics",
	Short: "statistics are displayed here",
	Long: "The following statistics are displayed here:\n" +
		"1- The total number of questions. \n" +
		"2-% of questions that have answers\n" +
		"3- The percentage of questions that have the correct answer",
	Run: func(cmd *cobra.Command, args []string) {
		displayTable.GameStatistics()
	},
}

func init() {
	rootCmd.AddCommand(StatisticsCmd)

}
