package main

import (
	"fmt"
	"os"

	"github.com/jkdihenkar/go_learning/21_goroutines/goroutines_examples"
)

func main() {
	var codenum string
	var option string

	if len(os.Args) < 3 {
		fmt.Print("Enter choice of codenum to run: ")
		fmt.Scanln(&codenum)
		fmt.Print("Enter choice of mode c/s: ")
		fmt.Scanln(&option)
	} else {
		codenum = os.Args[1]
		option = os.Args[2]
	}

	switch codenum {
	case "211":
		// channels and goroutines test example
		goroutines_examples.MainKeygenerator(option)
	case "212":
		goroutines_examples.MainChannelsAndPipelines(option)
	case "213":
		goroutines_examples.MainChannelsAndPipelinesEnhance(option)
	default:
		fmt.Println("Not a valid code section number!!")
	}

}
