package basics

import (
	"fmt"
	"unsafe"
)

/*
	print hello
	strings and types
	array
	loops
	range
*/

func MainHelloGo() {
	fmt.Println("Hello world....")
	mainForTypes()
}

func mainForTypes() {

	// ints
	t := 36
	fmt.Printf("%d\n", t)

	// string
	name := "jayD"
	for index, c := range name {
		fmt.Printf("%c(%d)|", c, index)
	}
	fmt.Println()

	// array of ints
	arr := []int{3, 4, 5, 6, 7}
	for _, v := range arr {
		fmt.Printf("%d_", v)
	}
	fmt.Println()

	// testing struct, func
	fmt.Println(testInts().second, testInts().fourth)

	// Inferred INT types (int32 for 32bit systems and int64 for 64bit systems)
	// Let's test out

	testAutoInferredType := 3778
	fmt.Printf("Inferred type of := 3778 is Type :: %T, Size :: %d[bytes] \n", testAutoInferredType, unsafe.Sizeof(testAutoInferredType))

	/*

		similarly uint => uint64 for 64bit
			float => float64
			complex => complex128 (r float64, i float64)

	*/

	sampleComplex := 2.37 + 99.42i
	fmt.Printf("Inferred type of := 2.37 + 99.42i is Type :: %T, Size :: %d[bytes] \n", sampleComplex, unsafe.Sizeof(sampleComplex))
}

type bunchOfNumbers struct {
	first  int16
	second int32
	third  int64
	fourth int8
}

func testInts() bunchOfNumbers {
	var (
		k int16 = 45
		m int32 = 15
		o int64 = 19990000
		p int8  = 123
	)

	return bunchOfNumbers{k, m, o, p}
}
