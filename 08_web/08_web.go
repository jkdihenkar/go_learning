package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world\n")
}

func test(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Testing more on routes\n")
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/test", test)
	fmt.Println("Server starting ...")
	http.ListenAndServe("127.0.0.1:9988", nil)
}

/*

[jd@jdpc go_learning]$ go run 08_web/08_web.go
Server starting ...
^Csignal: interrupt


[jd@jdpc go_learning]$ curl 127.0.0.1:9988/
Hello world
[jd@jdpc go_learning]$ curl 127.0.0.1:9988/test
Testing more on routes

*/
