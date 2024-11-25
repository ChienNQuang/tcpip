package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	Port    uint
	Message string
)

var rootCmd = &cobra.Command{
	Use:   "tcpip",
	Short: "tcpip",
	Long:  `tcpip`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
