/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	db2 "flashCardProject/db"
	"flashCardProject/internal"
	"github.com/spf13/cobra"
	"log"
)

// ListingFlashCardCmd represents the ListingFlashCard command
var ListingFlashCardCmd = &cobra.Command{
	Use:   "ListingFlashCard",
	Short: "list all Flashcards",

	Run: func(cmd *cobra.Command, args []string) {
		db, err := db2.SetUpDatabase()
		if err != nil {
			panic(err)
		}
		var flashCards []internal.FlashCard
		flashCards, err = internal.GetAllFlashcardsData(db)
		if err != nil {
			log.Fatalln(err)
		}
		internal.DisplayTable(flashCards)
	},
}

func init() {
	rootCmd.AddCommand(ListingFlashCardCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// ListingFlashCardCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// ListingFlashCardCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
