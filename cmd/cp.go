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
		if len(args) < 2 {
			fmt.Println("please provide both source and destination paths e.g(cpz cp <source> <destination>)")
			return
		}

		source := args[0]
		destination := args[1]
		showProgress, _ := cmd.Flags().GetBool("progress")

		err := lib.Copy(source, destination, showProgress)
		if err != nil {
			fmt.Printf("error : %v\n", err)
			os.Exit(1)
		}
		fmt.Println("copy completed successfully")
	},
}

func init() {
	rootCmd.AddCommand(cpCmd)

	// Define flag for cp command
	cpCmd.Flags().BoolP("progress", "p", true, "Show progress indicator")
}
