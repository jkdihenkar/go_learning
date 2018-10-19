package serve_utils

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func MainSimpleTCPClient(port string) {

	REMOTE := "127.0.0.1:" + port

	// connect to this socket
	conn, _ := net.Dial("tcp4", REMOTE)

	for {
		// read in input from stdin
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Text to send: ")
		text, _ := reader.ReadString('\n')
		// send to socket
		fmt.Fprintf(conn, text+"\n")
		// listen for reply
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			log.Fatalf("Error - %s", err)
		}
		fmt.Print("Message from server: " + message)
	}
}

/*

# start server
[jd@jdpc go_learning]$ go run 23_network_programming/go_servers.go 231 2033
Serving 127.0.0.1:54382

# other terminal client
[jd@jdpc go_learning]$ go run 23_network_programming/go_servers.go 236 2033
Text to send: hello
Message from server: 159
Text to send: test
Message from server: 643
Text to send: SHUTDOWN
2018/10/19 18:34:59 Error - EOF
exit status 1

*/
