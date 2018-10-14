package goroutines_examples

import (
	"fmt"
	"strconv"
	"sync"
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
func genNumbersT2(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

// Step 4 square the numbers in channel
func squareNumsT2(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			time.Sleep(1 * time.Second)
			out <- n * n
		}
		close(out)
	}()
	return out
}

func merge(cs []<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	// Start an output goroutine for each input channel in cs.  output
	// copies values from c to out until c is closed, then calls wg.Done.
	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}
	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	// Start a goroutine to close out once all the output goroutines are
	// done.  This must start after the wg.Add call.
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func MainChannelsAndPipelinesEnhance(option string) {

	worker_count, _ := strconv.Atoi(option)
	fmt.Println("Option is considered num of workers -> ", worker_count)

	var workers []<-chan int

	c1 := genNumbersT2(9, 10, 11, 12, 13, 2, 3, 4, 5, 6, 7, 8)

	for i := 0; i < worker_count; i++ {
		workers = append(workers, squareNumsT2(c1))
	}

	for n := range merge(workers) {
		fmt.Print(n, " ")
	}
}

/*

[jd@jdpc go_learning]$ time go run 21_goroutines/21_goroutines.go 213 1
Option is considered num of workers ->  1
81 100 121 144 169 4 9 16 25 36 49 64
real	0m12.252s
user	0m0.244s
sys	0m0.102s
[jd@jdpc go_learning]$ time go run 21_goroutines/21_goroutines.go 213 4
Option is considered num of workers ->  4
144 81 121 100 16 4 169 9 49 64 36 25
real	0m3.147s
user	0m0.165s
sys	0m0.052s
[jd@jdpc go_learning]$ time go run 21_goroutines/21_goroutines.go 213 8
Option is considered num of workers ->  8
9 121 100 16 169 144 4 81 25 64 36 49
real	0m2.207s
user	0m0.209s
sys	0m0.079s


*/
