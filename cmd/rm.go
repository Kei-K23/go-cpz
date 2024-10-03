/*
Copyright © 2024 Kei-K23 <arkar.dev.kei@gmail.com>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/Kei-K23/trashbox"
	"github.com/spf13/cobra"
)

// rmCmd represents the rm command
var rmCmd = &cobra.Command{
	Use:   "rm [file]",
	Short: "Remove file or directory with the feature of moving to Recycle bin or Trash box",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("please provide file name to remove")
			return
		}

		file := args[0]
		err := trashbox.MoveToTrashMacOS(file)
		if err != nil {
			fmt.Printf("error : %v\n", err)
			os.Exit(1)
		}

		fmt.Println("remove to trashbox completed successfully")
	},
}

func init() {
	rootCmd.AddCommand(rmCmd)
}
