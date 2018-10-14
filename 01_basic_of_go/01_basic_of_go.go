package main

/*

importing custom package
use of scanln
use of switch case

*/

import (
	"fmt"
	"os"

	"github.com/jkdihenkar/go_learning/01_basic_of_go/basics"
)

func main() {
	arguements := os.Args
	codenum := ""

	if len(arguements) < 2 {
		fmt.Print("Enter the basic go codenum to run: ")
		fmt.Scanln(&codenum)
	} else {
		codenum = arguements[1]
	}

	switch codenum {
	case "011":
		basics.MainHelloGo()
	case "012":
		basics.MainCmdLineArgs()
	default:
		fmt.Println("No such demo created yet!")
	}
}

/*

[jd@jdpc go_learning]$ go run 01_basic_of_go/01_basic_of_go.go
Enter the basic go codenum to run: 011
Hello world....
36
j(0)|a(1)|y(2)|D(3)|
3_4_5_6_7_
15 123

[jd@jdpc go_learning]$ go run 01_basic_of_go/01_basic_of_go.go 012 12 14 16
0 /tmp/go-build344579483/b001/exe/01_basic_of_go
1 012
2 12
3 14
4 16


*/
