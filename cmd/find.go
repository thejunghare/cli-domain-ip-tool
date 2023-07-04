/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// findCmd represents the find command
var findCmd = &cobra.Command{
	Use:   "find",
	Short: "find ip for domain",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("find called")

		if filePath == "" {
			fmt.Print("Please provide vaild file path")
			return
		}

		file, err := os.Open(filePath)
		if err != nil {
			fmt.Print("Failed to open th file")
			return
		}

		defer file.Close()
	},
}

var filePath string

func init() {
	/* defining a Flag that accepts an file path */
	rootCmd.Flags().StringVarP(&filePath, "file", "f", "", "File path")
	rootCmd.AddCommand(findCmd)
}
