package basics

import (
	"fmt"
	"os"
	"regexp"
)

func MainFindReplaceDemo() {

	var s [3]string
	s[0] = "1 b 3"
	s[1] = "11 a B 14 1 1"
	s[2] = "b 2 -3 B -5"

	parse, err := regexp.Compile("[bB]")

	if err != nil {
		fmt.Printf("Error compiling RE: %s\n", err)
		os.Exit(-1)
	}

	for i := 0; i < len(s); i++ {
		temp := parse.ReplaceAllString(s[i], "C")
		fmt.Println(temp)
	}
}

/*

[jd@jdpc go_learning]$ go run 01_basic_of_go/01_basic_of_go.go 016
1 C 3
11 a C 14 1 1
C 2 -3 C -5

*/
