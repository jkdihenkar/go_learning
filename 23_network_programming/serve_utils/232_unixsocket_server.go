package serve_utils

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func MainUnixSocketServerConcurrent(socket_path string) {
	ADDRESS := socket_path

	l, err := net.Listen("unix", ADDRESS)
	if err != nil {
		log.Panicf("Error - %s\n", err)
		return
	}
	defer l.Close()

	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, os.Interrupt, os.Kill, syscall.SIGTERM)

	go func(c chan os.Signal) {
		// Wait for a SIGINT or SIGKILL:
		sig := <-c
		log.Printf("Caught signal %s: shutting down.", sig)
		// Stop listening (and unlink the socket if unix type):
		l.Close()
	}(sigc)

	for {
		c, err := l.Accept()
		if err != nil {
			log.Printf("Unexpected error or server shutdown - %s\n", err)
			return
		}
		// HandleConnection(c) - will be single threaded
		go HandleConnection(c)
	}
}

/*
	[jd@jdpc go_learning]$ go run 23_network_programming/go_servers.go 232 ./go.sock
	Serving @
	Serving @
	Recieved connection close - EOF
	Recieved connection close - EOF

	nc -U ./go.sock # over multiple terminals
*/
