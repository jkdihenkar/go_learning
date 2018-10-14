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

/*

[jd@jdpc go_learning]$ go run 10_error_handling/10_error_handling.go
2018/10/14 14:27:37 Running test with 2/2 ...
2018/10/14 14:27:37 The result is 1
2018/10/14 14:27:37 Running test with 7/13 ...
2018/10/14 14:27:37 The result is 0
2018/10/14 14:27:37 There is a remainder!
2018/10/14 14:27:37 Running test with 15/0 ...
2018/10/14 14:27:37 Cannot divide by zero!
panic: Cannot divide by zero!

goroutine 1 [running]:
log.Panic(0xc42004ff18, 0x1, 0x1)
	/usr/lib/golang/src/log/log.go:326 +0xc0
main.testDivide(0xf, 0x0)
	learning_go/src/github.com/jkdihenkar/go_learning/10_error_handling/10_error_handling.go:36 +0xaf
main.main()
	learning_go/src/github.com/jkdihenkar/go_learning/10_error_handling/10_error_handling.go:55 +0xe1
exit status 2


*/
