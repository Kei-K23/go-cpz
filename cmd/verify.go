/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/Kei-K23/cpz/internal/lib"
	"github.com/spf13/cobra"
)

// verifyCmd represents the verify command
var verifyCmd = &cobra.Command{
	Use:   "verify [source] [destination]",
	Short: "Verify and identify between source and destination file or directory",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			fmt.Println("please provide both source and destination paths e.g(cpz verify <source> <destination>)")
			return
		}

		source := args[0]
		destination := args[1]

		err := lib.Verify(source, destination)
		if err != nil {
			fmt.Printf("Verification failed: %v\n", err)
		} else {
			fmt.Println("Source and destination are identical.")
		}
	},
}

func init() {
	rootCmd.AddCommand(verifyCmd)
}
