package main

import (
	"errors"
	"fmt"
	"log"
)

func division(x, y int) (int, error, error) {
	if y == 0 {
		return 0, nil, errors.New("Cannot divide by zero!")
	}
	if x%y != 0 {
		remainder := errors.New("There is a remainder!")
		return x / y, remainder, nil
	} else {
		return x / y, nil, nil
	}

}

func testDivide(x, y int) {
	result, rem, err := division(x, y)

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("The result is", result)
	}
	if rem != nil {
		fmt.Println(rem)
	}
}

func main() {
	testDivide(2, 2)
	testDivide(7, 13)

	// will generate log.Fatal()
	testDivide(15, 0)

	fmt.Println("Anything after log.Fatal() will not be executed!")
}
