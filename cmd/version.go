package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use: "version",
	Short: "Print the version number of dadjoke",
	Long: "All software has versions, even ones that print out dad jokes.",
	Run: func(cmd *cobra.Command, args[]string) {
		fmt.Println("dadjokes version: ugh")
	},
}