package main

import (
	"fmt"
	"os"

	"github.com/jkdihenkar/go_learning/23_network_programming/serve_utils"
)

func main() {
	var codenum string
	var option string

	if len(os.Args) < 3 {
		fmt.Print("Enter choice of codenum to run: ")
		fmt.Scanln(&codenum)
		fmt.Print("Enter additional option: ")
		fmt.Scanln(&option)
	} else {
		codenum = os.Args[1]
		option = os.Args[2]
	}

	switch codenum {
	case "231":
		// concurrent tcp server
		serve_utils.MainTcpServerConcurrent(option)
	case "232":
		// concurrent unix socket server
		serve_utils.MainUnixSocketServerConcurrent(option)
	case "233":
		// concurrent rpc server
		serve_utils.MainRPCHttpServer(option)
	case "234":
		// rpc client sync call
		serve_utils.MainRPCHttpClientSync(option)
	case "235":
		// rpc client sync call
		serve_utils.MainRPCHttpClientAsync(option)
	case "236":
		// tcp client
		serve_utils.MainSimpleTCPClient(option)
	case "237":
		// concurrent UDP server
		serve_utils.MainUDPServerConcurrent(option)
	case "238":
		// simple udp client
		serve_utils.MainSimpleUDPClient(option)
	default:
		fmt.Println("Not a valid code section number!!")
	}

}

/*
	go run 23_network_programming/go_servers.go 231 2230
		# 231 code section tcp server
		# 2230 = port
		Test concurrency by running nc in two seperate terminals -
		nc localhost 2230

	go run 23_network_programming/go_servers.go 232 ./go.sock
		nc -U ./go.sock

*/
