package basics

import (
	"fmt"
)

/*
	1D, 2D, 3D arrays
	slices
	pass by value
	limitations of array
	slices are dynamic
	pass by reference
*/

func MainArraySlicesDemo() {
	arraysAndDimensions()

	oneD := [4]int{1, 2, 4, -4}
	fmt.Println("Manipulate by array defination - ")
	changeArr(oneD)
	fmt.Println(oneD)

	fmt.Println("Manipulate by slice defination - ")
	change(oneD[0:4])
	fmt.Println(oneD)

	fmt.Println("Slice Allocations and cap - ")
	printAppendSlice(oneD[0:4])
}

func changeArr(x [4]int) {
	x[2] = 99
}

func change(x []int) {
	x[2] = 99
}

func printAppendSlice(x []int) {
	fmt.Printf("Before = Capacity: %d, Length: %d\n", cap(x), len(x))
	for _, number := range x {

		fmt.Printf("%d ", number)
	}
	fmt.Println()

	aSlice := append(x, -199)
	fmt.Printf("After = Capacity: %d, Length: %d\n", cap(aSlice), len(aSlice))
	for _, number := range aSlice {

		fmt.Printf("%d ", number)
	}
	fmt.Println()

}

func arraysAndDimensions() {
	oneD := [4]int{1, 2, 4, -4}
	twoD := [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	threeD := [2][2][2]int{{{1, 2}, {3, 4}}, {{5, 6}, {7, 8}}}

	fmt.Println("Initially...")
	fmt.Println(oneD, " | len = ", len(oneD))
	fmt.Println(twoD, " | len = ", len(twoD))
	fmt.Println(threeD, " | len = ", len(threeD))

	oneD[0] = 7
	twoD[1][2] = 15
	threeD[0][1][1] = -1

	fmt.Println("After change...")
	fmt.Println(oneD)
	fmt.Println(twoD)
	fmt.Println(threeD)

	fmt.Println("Range 3D example...")
	for idx, val := range threeD {
		fmt.Printf("(%d %d) ", idx, val)
	}
	fmt.Println()
}

/*

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

*/
