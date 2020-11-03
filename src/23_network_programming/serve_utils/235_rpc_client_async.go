package serve_utils

import (
	"fmt"
	"log"
	"net/rpc"
)

func MainRPCHttpClientAsync(port string) {

	PORT := ":" + port

	client, err := rpc.DialHTTP("tcp4", PORT)
	if err != nil {
		log.Fatal("dialing:", err)
	}

	// Synchronous call
	args1 := &Args{100, 23}

	reply2 := new(Quotient)

	// Test multiply
	divCall := client.Go("Arith.Divide", args1, reply2, nil)
	reply_channel := <-divCall.Done

	fmt.Println(*(reply_channel.Reply.(*Quotient)))

}

/*
	[jd@jdpc go_learning]$ go run 23_network_programming/go_servers.go 235 2033
	{4 8}
*/
