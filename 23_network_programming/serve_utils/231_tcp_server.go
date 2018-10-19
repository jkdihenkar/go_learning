package serve_utils

import (
	"fmt"
	"net"
)

func MainTcpServerConcurrent(listen_port string) {
	PORT := ":" + listen_port
	// even try tcp6
	l, err := net.Listen("tcp4", PORT)
	if err != nil {
		fmt.Printf("Error - %s", err)
		return
	}
	defer l.Close()

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		// HandleConnection(c) - will be single threaded
		go HandleConnection(c)
	}
}

/*
	go run 23_network_programming/go_servers.go 231 2230
		# 231 code section tcp server
		# 2230 = port
		Test concurrency by running nc in two seperate terminals -
		nc localhost 2230
*/
