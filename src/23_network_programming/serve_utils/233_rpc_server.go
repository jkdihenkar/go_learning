package serve_utils

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
)

func MainRPCHttpServer(port string) {
	PORT := ":" + port
	arith := new(Arith)
	rpc.Register(arith)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp4", PORT)
	if e != nil {
		log.Fatal("listen error:", e)
	}
	http.Serve(l, nil)
}
