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
	Short: "move files or directories with progress bar",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			fmt.Println("please provide both source and destination paths e.g(cpz mv <source> <destination>)")
			return
		}

		source := args[0]
		destination := args[1]
		showProgress, _ := cmd.Flags().GetBool("progress")
		excludeFilenames, _ := cmd.Flags().GetStringSlice("filter")

		err := lib.Copy(source, destination, showProgress, excludeFilenames, nil)
		if err != nil {
			fmt.Printf("error : %v\n", err)
			os.Exit(1)
		}

		err = os.RemoveAll(source)
		if err != nil {
			fmt.Printf("error : %v\n", err)
			os.Exit(1)
		}
		fmt.Println("move completed successfully")
	},
}

func init() {
	rootCmd.AddCommand(mvCmd)

	// Define flag for cp command
	mvCmd.Flags().BoolP("progress", "p", true, "Show progress indicator")
	mvCmd.Flags().StringSliceP("filter", "f", nil, "Filter file name to exclude when copying")
}
