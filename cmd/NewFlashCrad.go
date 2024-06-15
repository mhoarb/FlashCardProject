// Package cmd provides the commands for the CLI application.
package cmd

import (
	"bufio"
	"flashCardProject/internal"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

// NewFlashCardCmd represents the NewFlashCard command
var NewFlashCardCmd = &cobra.Command{
	Use:   "NewFlashCard",
	Short: "Creating new flashCard easily",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("please enter your question:")
		question, _ := reader.ReadString('\n')
		question = question[:len(question)-1]

		fmt.Println("please enter your answer :")
		answer, _ := reader.ReadString('\n')
		answer = answer[:len(answer)-1]
		internal.NewFlashCard(question, answer)
	},
}

func init() {
	rootCmd.AddCommand(NewFlashCardCmd)
}
