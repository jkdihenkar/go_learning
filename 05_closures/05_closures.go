package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
	basic adder closure for sum upto n
*/

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter the limit to sum upto: ")

	numStr, _ := reader.ReadString('\n')
	fmt.Printf("Input is %s\n", numStr)

	n, e := strconv.ParseInt(strings.TrimSuffix(numStr, "\n"), 10, 32)
	fmt.Printf("DEBUG: Value of n is %d, err = %s\n", n, e)

	sum := adder()
	finalResult := 0

	for i := 1; i <= int(n); i++ {
		finalResult = sum(i)
	}

	fmt.Println(finalResult)

}

/*

[jd@jdpc go_learning]$ go run 05_closures/05_closures.go
Enter the limit to sum upto: 100
Input is 100

DEBUG: Value of n is 100, err = %!s(<nil>)
5050


*/
