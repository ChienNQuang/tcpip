package cmd

import (
	"fmt"
	"io"
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
		// accept loop
		for {
			conn, err := l.Accept()
			if err != nil {
				fmt.Errorf("accept error: %v", err)
				continue
			}
			go func() {
				defer conn.Close()
				// read loop
				for {
					buf := make([]byte, 1024)
					n, err := conn.Read(buf)
					if err != nil {
						if err == io.EOF {
							fmt.Println("client", conn.RemoteAddr().String(), "disconnected")
							break
						}
						fmt.Printf("read error: %v\n", err)
						break
					}

					message := string(buf[:n])
					fmt.Printf("Received message from (%s): %s\n", conn.RemoteAddr().String(), message)
					conn.Write([]byte(fmt.Sprintf("We have received your message: %s", message)))
				}
			}()
		}
	},
}
