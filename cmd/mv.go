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

// mvCmd represents the mv command
var mvCmd = &cobra.Command{
	Use:   "mv [source] [destination]",
	Short: "Move files or directories with progress bar",
	Run: func(cmd *cobra.Command, args []string) {
		source := args[0]
		destination := args[1]
		showProgress, _ := cmd.Flags().GetBool("progress")

		err := lib.Copy(source, destination, showProgress)
		if err != nil {
			fmt.Printf("Error : %v\n", err)
			os.Exit(1)
		}

		err = os.RemoveAll(source)
		if err != nil {
			fmt.Printf("Error : %v\n", err)
			os.Exit(1)
		}
		fmt.Println("Move completed successfully")
	},
}

func init() {
	rootCmd.AddCommand(mvCmd)

	// Define flag for cp command
	mvCmd.Flags().BoolP("progress", "p", true, "Show progress indicator")
}
