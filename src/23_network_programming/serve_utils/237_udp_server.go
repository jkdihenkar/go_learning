package serve_utils

import (
	"fmt"
	"net"
)

func MainUDPServerConcurrent(listen_port string) {
	PORT := ":" + listen_port
	// even try tcp6
	l, err := net.ListenPacket("udp", PORT)
	if err != nil {
		fmt.Printf("Error - %s", err)
		return
	}
	defer l.Close()

	for {
		buf := make([]byte, 1024)
		n, addr, err := l.ReadFrom(buf)
		if err != nil {
			fmt.Println(err)
			return
		}
		go EchoMe(addr, n, buf)
	}
}

/*

# start server -
[jd@jdpc go_learning]$ go run 23_network_programming/go_servers.go 237 2033
[::1]:37878
5
hello

# to test -
[jd@jdpc go_learning]$ echo -n "hello" >/dev/udp/localhost/2033

*/
