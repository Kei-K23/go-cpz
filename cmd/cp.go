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
	Short: "copy files or directories with progress bar",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			fmt.Println("please provide both source and destination paths e.g(cpz cp <source> <destination>)")
			return
		}

		source := args[0]
		destination := args[1]
		showProgress, _ := cmd.Flags().GetBool("progress")
		excludeFilenames, _ := cmd.Flags().GetStringSlice("filter")
		excludeExtensions, _ := cmd.Flags().GetStringSlice("extensions")
		excludeRegexs, _ := cmd.Flags().GetStringSlice("regex")

		err := lib.Copy(source, destination, showProgress, excludeFilenames, excludeExtensions, excludeRegexs)

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
	cpCmd.Flags().BoolP("progress", "p", false, "Show progress indicator")
	cpCmd.Flags().StringSliceP("filter", "f", nil, "Filter files to exclude when copying with file names")
	cpCmd.Flags().StringSliceP("extensions", "e", nil, "Filter files to exclude when copying with extensions")
	cpCmd.Flags().StringSliceP("regex", "r", nil, "Filter files to exclude when copying with regular expression")
}
