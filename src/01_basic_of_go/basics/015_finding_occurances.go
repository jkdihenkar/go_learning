package basics

import (
	"fmt"
	"strings"
)

func testStringFields() {
	myString := "Finding the number of occurrences"
	fmt.Println(strings.Fields(myString))
}

func MainOcuurancesDemo() {

	fmt.Println("Testing output of string.Fields ...")
	testStringFields()

	var s [3]string
	s[0] = "1 b 3 1 a a b"
	s[1] = "11 a 1 1 1 1 a a"
	s[2] = "-1 b 1 -4 a 1"

	counts := make(map[string]int)

	for i := 0; i < len(s); i++ {
		data := strings.Fields(s[i])
		for _, word := range data {
			_, ok := counts[word]
			if ok {
				counts[word] = counts[word] + 1
			} else {
				counts[word] = 1
			}
		}
	}

	for key, _ := range counts {

		fmt.Printf("%s -> %d \n", key, counts[key])
	}
}

/*

[jd@jdpc go_learning]$ go run 01_basic_of_go/01_basic_of_go.go 015 | sort -t\  -k3,3 -rn
1 -> 8
a -> 6
b -> 3
-4 -> 1
3 -> 1
11 -> 1
-1 -> 1
Testing output of string.Fields ...
[Finding the number of occurrences]

*/
