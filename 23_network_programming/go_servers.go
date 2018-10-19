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
		// channels and goroutines test example
		serve_utils.MainTcpServerConcurrent(option)
	default:
		fmt.Println("Not a valid code section number!!")
	}

}
