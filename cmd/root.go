package cmd

import "github.com/spf13/cobra"

var RootCmd = &cobra.Command{
	Use:   "domain-ip-tool",
	Short: "A CLI tool to find IP addresses from domains",
}
