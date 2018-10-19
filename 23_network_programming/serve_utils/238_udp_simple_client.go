package serve_utils

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func MainSimpleUDPClient(port string) {

	REMOTE := "127.0.0.1:" + port

	// connect to this socket
	conn, _ := net.Dial("udp", REMOTE)

	for {
		// read in input from stdin
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Don't care - keep writing : ")
		text, _ := reader.ReadString('\n')
		// send to socket
		fmt.Fprintf(conn, text)
	}
}

/*

# start server
[jd@jdpc go_learning]$ go run 23_network_programming/go_servers.go 237 2033
127.0.0.1:40426
5
test

127.0.0.1:40426
4
123


# other terminal client
[jd@jdpc go_learning]$ go run 23_network_programming/go_servers.go 238 2033
Don't care - keep writing : test
Don't care - keep writing : 123
Don't care - keep writing : ^Csignal: interrupt


*/
