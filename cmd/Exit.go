/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"log/slog"
	"os"
)

// ExitCmd represents the Exit command
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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// ExitCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// ExitCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
