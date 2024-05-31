package cmd

import (
	"flashCardProject/internal/play"
	"github.com/spf13/cobra"
)

// PLAYCmd represents the PLAY command
var PLAYCmd = &cobra.Command{
	Use:   "PLAY",
	Short: "playing flashcards and answer the question",

	Run: func(cmd *cobra.Command, args []string) {
		play.Play()
	},
}

func init() {
	rootCmd.AddCommand(PLAYCmd)

}
