/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// PLAYCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// PLAYCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
