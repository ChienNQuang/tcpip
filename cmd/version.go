package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "tcpip",
	Long:  `tcpip`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("tcpip")
	},
}
