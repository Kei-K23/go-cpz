/*
Copyright © 2024 Kei-K23 <arkar.dev.kei@gmail.com>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/Kei-K23/cpz/internal/lib"
	"github.com/spf13/cobra"
)

// cpCmd represents the cp command
var cpCmd = &cobra.Command{
	Use:   "cp [source] [destination]",
	Short: "Copy files or directories with progress bar",
	Run: func(cmd *cobra.Command, args []string) {
		source := args[0]
		destination := args[1]
		showProgress, _ := cmd.Flags().GetBool("progress")

		err := lib.Copy(source, destination, showProgress)
		if err != nil {
			fmt.Printf("Error : %v\n", err)
			os.Exit(1)
		}
		fmt.Println("Copy completed successfully")
	},
}

func init() {
	rootCmd.AddCommand(cpCmd)

	// Define flag for cp command
	cpCmd.Flags().BoolP("progress", "p", true, "Show progress indicator")
}
