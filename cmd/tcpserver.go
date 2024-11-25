package cmd

import (
	"fmt"
	"net"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(tcpserverCmd)
	tcpserverCmd.Flags().UintVarP(&Port, "port", "p", 3000, "port")
}

var tcpserverCmd = &cobra.Command{
	Use:   "tcpserver",
	Short: "tcpserver",
	Long:  `tcpserver`,
	Run: func(cmd *cobra.Command, args []string) {
		l, err := net.ListenTCP("tcp", &net.TCPAddr{
			IP:   net.IPv4(127, 0, 0, 1),
			Port: int(Port),
		})
		if err != nil {
			panic(err)
		}
		fmt.Println("started TCP server on port", Port)
		for {
			conn, err := l.Accept()
			if err != nil {
				fmt.Errorf("accept error: %v", err)
			}
			go func() {
				defer conn.Close()
				for {
					buf := make([]byte, 1024)
					n, err := conn.Read(buf)
					if err != nil {
						fmt.Printf("read error: %v\n", err)
						break
					}
					// conn.Write(buf[:n])
					fmt.Println("message received:", string(buf[:n]))
				}
			}()
		}
	},
}
