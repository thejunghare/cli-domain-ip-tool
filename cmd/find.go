package cmd

import (
	"bufio"
	"fmt"
	"net"
	"os"

	"github.com/spf13/cobra"
)

var filePath string

// findCmd represents the find command
var findCmd = &cobra.Command{
	Use:   "find",
	Short: "Find IP address of domain",
	Run: func(cmd *cobra.Command, args []string) {
		// open the file
		file, err := os.Open(filePath)
		checkerr(err)
		// fmt.Println("File found")

		// Read the file
		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanLines)

		// Print content of file
		// domain -> ip
		for scanner.Scan(){
			ip, _ := net.LookupIP(scanner.Text())
			// fmt.Println(scanner.Text())
			fmt.Println(ip)
		}

		// checkerr(err)

		// close the file
		defer file.Close()
	},
}

func checkerr(err error) {
	if err != nil {
		fmt.Print(err)
	}
}

func init() {
	findCmd.Flags().StringVarP(&filePath, "file", "f", "", "File path")
	// findCmd.MarkFlagRequired("file") // Mark the file flag as required
	RootCmd.AddCommand(findCmd)
}
