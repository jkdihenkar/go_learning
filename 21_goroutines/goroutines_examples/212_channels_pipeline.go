package goroutines_examples

import (
	"fmt"
	"time"
)

/*
	Pipeline ->
		generate n numbers
		add 3 to it
		subtract 1 from it
		square it

*/

// Step 1 list of numbers to int channel
func genNumbers(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

// Step 2 add 3 to the numbers in channel
func addThree(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n + 3
		}
		close(out)
	}()
	return out
}

// Step 3 subtract 1 from the numbers in channel
func subOne(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			// this operation is slow and expensive
			time.Sleep(1 * time.Second)
			out <- n - 1
		}
		close(out)
	}()
	return out
}

// Step 4 square the numbers in channel
func squareNums(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func MainChannelsAndPipelines(option string) {

	fmt.Println("Ignoring option here ...", option)

	for n := range squareNums(
		subOne(
			addThree(
				genNumbers(9, 10, 11, 12, 13, 2, 3, 4, 5, 6, 7, 8)))) {
		fmt.Print(n, " ")
	}
}

/*

[jd@jdpc go_learning]$ time go run 21_goroutines/21_goroutines.go 212 c
Ignoring option here ... c
121 144 169 196 225 16 25 36 49 64 81 100
real	0m12.177s
user	0m0.181s
sys	0m0.062s

# This could have been faster with more consumers in channel

*/
