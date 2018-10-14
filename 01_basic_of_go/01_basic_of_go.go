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
	case "013":
		basics.MainArraySlicesDemo()
	case "014":
		basics.MainRandomNumbers()
	case "015":
		basics.MainOcuurancesDemo()
	case "016":
		basics.MainFindReplaceDemo()
	case "017":
		basics.MainSortSliceDemo()
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

[jd@jdpc go_learning]$ go run 01_basic_of_go/01_basic_of_go.go 013
Initially...
[1 2 4 -4]  | len =  4
[[1 2 3] [4 5 6] [7 8 9]]  | len =  3
[[[1 2] [3 4]] [[5 6] [7 8]]]  | len =  2
After change...
[7 2 4 -4]
[[1 2 3] [4 5 15] [7 8 9]]
[[[1 2] [3 -1]] [[5 6] [7 8]]]
Range 3D example...
(0 [[1 2] [3 -1]]) (1 [[5 6] [7 8]])
Manipulate by array defination -
[1 2 4 -4]
Manipulate by slice defination -
[1 2 99 -4]
Slice Allocations and cap -
Before = Capacity: 4, Length: 4
1 2 99 -4
After = Capacity: 8, Length: 5
1 2 99 -4 -199


[jd@jdpc go_learning]$ go run 01_basic_of_go/01_basic_of_go.go 014
781 163 664 155 153 365 822 152 914 654
[jd@jdpc go_learning]$ go run 01_basic_of_go/01_basic_of_go.go 014
923 352 823 325 438 997 564 811 421 957
[jd@jdpc go_learning]$ go run 01_basic_of_go/01_basic_of_go.go 014
807 822 843 923 101 616 541 919 580 772

*/
