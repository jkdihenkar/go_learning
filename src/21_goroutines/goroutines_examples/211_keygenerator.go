package goroutines_examples

import (
	"fmt"
	"math/rand"
	"time"
)

func generateKeyConcurrent(channel chan int) {
	fmt.Println("Generating key")
	// Super-secret algorithm!
	keys := []int{3, 5, 7, 11}
	key := keys[rand.Intn(len(keys))]
	// It's kinda slow!
	time.Sleep(3 * time.Second)
	fmt.Print("Done generating -> ")
	channel <- key
}

func MainConcurrentKeygen() {
	rand.Seed(time.Now().Unix())
	channel := make(chan int)
	// Call generateKey 3 times.
	for i := 0; i < 3; i++ {
		go generateKeyConcurrent(channel)
	}

	for i := 0; i < 3; i++ {
		fmt.Println(<-channel)
	}

	fmt.Println("All done!")
}

func generateKeySequential() int {
	fmt.Println("Generating key")
	// Super-secret algorithm!
	keys := []int{3, 5, 7, 11}
	key := keys[rand.Intn(len(keys))]
	// It's kinda slow!
	time.Sleep(3 * time.Second)
	fmt.Print("Done generating -> ")
	return key
}

func MainSequentialKeygen() {
	rand.Seed(time.Now().Unix())
	// Call generateKey 3 times.
	for i := 0; i < 3; i++ {
		fmt.Println(generateKeySequential())
	}
	fmt.Println("All done!")
}

func MainKeygenerator(option string) {

	switch option {
	case "c":
		MainConcurrentKeygen()
	case "s":
		MainSequentialKeygen()
	default:
		fmt.Println("Only input allowed are c and s!!")
	}
}

/*

[jd@jdpc go_learning]$ time go run 21_goroutines/21_goroutines.go 211 c
Generating key
Generating key
Generating key
Done generating -> 7
Done generating -> 11
Done generating -> 3
All done!

real	0m3.149s
user	0m0.167s
sys	0m0.045s
[jd@jdpc go_learning]$ time go run 21_goroutines/21_goroutines.go 211 s
Generating key
Done generating -> 5
Generating key
Done generating -> 5
Generating key
Done generating -> 5
All done!

real	0m9.198s
user	0m0.188s
sys	0m0.072s

*/
