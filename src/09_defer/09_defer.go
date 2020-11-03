package main

import "fmt"

/*
	deferred functions are executed in Last In First Out (LIFO)
		order after the return of the surrounding function
*/

func a1() {
	/*
		The for loop in a1() defers a single function call
			that uses the current value of the i variable.
		As a result, all numbers are printed in reverse order
			because the last used value of i is 3.
	*/
	for i := 1; i <= 3; i++ {
		defer fmt.Print(i, " ")
	}
}

func a2() {
	/*
		the function body is evaluated after the for loop ends while it is
			still referencing the local i variable, which at that time was
			equal to 4 for all evaluations of the body.
	*/
	for i := 1; i <= 3; i++ {
		defer func() { fmt.Print(i, " ") }()
	}
}

func a3() {
	/*
		current value of i is passed as an argument to the deferred function,
			due to the (i) code at the end of the a3() function definition.
			So, each time the deferred function is executed,
			it has a different i value to process.
	*/
	for i := 1; i <= 3; i++ {
		defer func(n int) { fmt.Print(n, " ") }(i)
	}
}

func main() {
	a1()
	fmt.Println()
	a2()
	fmt.Println()
	a3()
	fmt.Println()
}

/*

[jd@jdpc go_learning]$ go run 09_defer/09_defer.go
3 2 1
4 4 4
3 2 1

*/
