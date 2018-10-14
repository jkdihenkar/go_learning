package main

import (
	"errors"
	"fmt"
	"log"
)

/*

error handling
logging
error logging
panic and fatal

*/

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
		// prints some additional information, useful for debugging purposes.
		log.Panic(err)
	} else {
		log.Printf("The result is %d", result)
	}
	if rem != nil {
		log.Print(rem)
	}
}

func main() {

	log.Printf("Running test with 2/2 ...")
	testDivide(2, 2)

	log.Printf("Running test with 7/13 ...")
	testDivide(7, 13)

	// will generate log.Fatal()
	log.Printf("Running test with 15/0 ...")
	testDivide(15, 0)

	fmt.Println("Anything after log.Fatal() will not be executed!")
}
