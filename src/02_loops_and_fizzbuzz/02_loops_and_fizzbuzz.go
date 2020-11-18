package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
	accept input from prompt
	fizz = div by 3
	buzz = div by 5
	fizzbuzz = div by 15
*/

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter the limit n: ")

	numStr, _ := reader.ReadString('\n')
	fmt.Printf("Input is %s\n", numStr)

	n, e := strconv.ParseInt(strings.TrimSuffix(numStr, "\n"), 10, 32)
	fmt.Printf("DEBUG: Value of n is %d, err = %s\n", n, e)

	var i int64

	for i = 1; i <= n; i++ {
		if i%15 == 0 {
			fmt.Println(i, " = fizzbuzz")
		} else if i%5 == 0 {
			fmt.Println(i, " = buzz")
		} else if i%3 == 0 {
			fmt.Println(i, " = fizz")
		}
	}

	// there is another way to do this with expression less switch-case
	fmt.Println("Output of fizzbuzz using switch-case :: ")

	for i = 1; i <= n; i++ {
		switch {
		case i%15 == 0:
			fmt.Println(i, " = fizzbuzz")
			break
		case i%5 == 0:
			fmt.Println(i, " = buzz")
			break
		case i%3 == 0:
			fmt.Println(i, " = fizz")
			break
		}
	}

	// switch t := t.(type) ==> Type Switch with case bool, case int, case *int and so on...
	// switch cases can be combined as case 1,2,3,4: Expr... for multiple values same expr

}

/* SAMPLE OUTPUT -
[jay@localhost go_learning]$ go run src/02_loops_and_fizzbuzz/02_loops_and_fizzbuzz.go
Enter the limit n: 18
Input is 18

DEBUG: Value of n is 18, err = %!s(<nil>)
3  = fizz
5  = buzz
6  = fizz
9  = fizz
10  = buzz
12  = fizz
15  = fizzbuzz
18  = fizz

[jay@localhost go_learning]$ go run src/02_loops_and_fizzbuzz/02_loops_and_fizzbuzz.go
Enter the limit n: qw
Input is qw

DEBUG: Value of n is 0, err = strconv.ParseInt: parsing "qw": invalid syntax
*/
