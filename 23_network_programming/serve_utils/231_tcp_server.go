package serve_utils

import (
	"fmt"
	"math/rand"
	"net"
	"time"
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
	rand.Seed(time.Now().Unix())

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
