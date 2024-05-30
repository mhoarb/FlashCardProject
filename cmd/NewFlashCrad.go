/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"flashCardProject/internal"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

// NewFlashCardCmd represents the NewFlashCrad command
var NewFlashCardCmd = &cobra.Command{
	Use:   "NewFlashCard",
	Short: "Creating new flashCard easily",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("please enter your question:")
		question, _ := reader.ReadString('\n')
		question = question[:len(question)-1] // حذف کاراکتر newline

		fmt.Println("please enter your answer :")
		answer, _ := reader.ReadString('\n')
		answer = answer[:len(answer)-1]
		internal.NewFlashCard(question, answer)
	},

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// NewFlashCardCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// NewFlashCardCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func init() {
	rootCmd.AddCommand(NewFlashCardCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// RESTARTCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// RESTARTCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
