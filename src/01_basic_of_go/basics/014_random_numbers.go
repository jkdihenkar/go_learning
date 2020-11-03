package basics

import (
	"fmt"
	"math/rand"
	"time"
)

/*
	random numbers in GO
*/

func MainRandomNumbers() {
	minLimit := 0
	maxLimit := 999
	numbers := 10

	rand.Seed(time.Now().Unix())

	for i := 1; i <= numbers; i++ {
		fmt.Print(rand.Intn(maxLimit-minLimit)+minLimit, " ")
	}

	fmt.Println()

}

/*

[jd@jdpc go_learning]$ go run 01_basic_of_go/01_basic_of_go.go 014
781 163 664 155 153 365 822 152 914 654
[jd@jdpc go_learning]$ go run 01_basic_of_go/01_basic_of_go.go 014
923 352 823 325 438 997 564 811 421 957
[jd@jdpc go_learning]$ go run 01_basic_of_go/01_basic_of_go.go 014
807 822 843 923 101 616 541 919 580 772

*/
