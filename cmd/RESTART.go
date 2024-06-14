// Package cmd provides the commands for the CLI application.
package cmd

import (
	"flashCardProject/internal/play"
	"github.com/spf13/cobra"
)

// RESTARTCmd represents the RESTART command
var RESTARTCmd = &cobra.Command{
	Use:   "RESTART",
	Short: "Reset the game",

	Run: func(cmd *cobra.Command, args []string) {
		play.Restart()
	},
}

func init() {
	rootCmd.AddCommand(RESTARTCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// RESTARTCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// RESTARTCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
