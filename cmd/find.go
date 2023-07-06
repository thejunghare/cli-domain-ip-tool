package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// findCmd represents the find command
var findCmd = &cobra.Command{
	Use:   "find",
	Short: "Find IP address of domain",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("find called")
	},
}

var filePath string

func init() {
	RootCmd.AddCommand(findCmd)
	RootCmd.Flags().StringVarP(&filePath, "file", "f", "", "File path")
}
