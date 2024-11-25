package cmd

import (
	"net"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(tcpclientCmd)
	tcpclientCmd.Flags().UintVarP(&Port, "port", "p", 3000, "port")
	tcpclientCmd.Flags().StringVarP(&Message, "message", "m", "hello", "message")
}

var tcpclientCmd = &cobra.Command{
	Use:   "tcpclient",
	Short: "tcpclient",
	Long:  `tcpclient`,
	Run: func(cmd *cobra.Command, args []string) {
		conn, err := net.DialTCP("tcp", nil, &net.TCPAddr{
			IP:   net.IPv4(127, 0, 0, 1),
			Port: int(Port),
		})

		if err != nil {
			panic(err)
		}

		conn.Write([]byte(Message))

		conn.Close()
	},
}
