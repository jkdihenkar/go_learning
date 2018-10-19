package serve_utils

import (
	"fmt"
	"log"
	"net/rpc"
)

func MainRPCHttpClientSync(port string) {

	PORT := ":" + port

	client, err := rpc.DialHTTP("tcp4", PORT)
	if err != nil {
		log.Fatal("dialing:", err)
	}

	// Synchronous call
	args1 := &Args{7, 8}
	args2 := &Args{100, 23}

	var reply int
	reply2 := new(Quotient)

	// Test multiply
	err = client.Call("Arith.Multiply", args1, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("Arith: %d*%d=%d\n", args1.A, args1.B, reply)

	// Test divide
	err = client.Call("Arith.Divide", args2, &reply2)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("Arith: %d/%d=%d(rem=%d)\n", args2.A, args2.B, reply2.Quo, reply2.Rem)

}

/*
	[jd@jdpc go_learning]$ go run 23_network_programming/go_servers.go 234 2033
	Arith: 7*8=56
	Arith: 100/23=4(rem=8)
*/
