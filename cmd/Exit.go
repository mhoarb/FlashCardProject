package cmd

import (
	"github.com/spf13/cobra"
	"log/slog"
	"os"
)

var ExitCmd = &cobra.Command{
	Use:   "Exit",
	Short: "Exit And Finish thw program",

	Run: func(cmd *cobra.Command, args []string) {
		slog.Info("Exiting the application...")
		os.Exit(0)
	},
}

func init() {
	rootCmd.AddCommand(ExitCmd)

}
