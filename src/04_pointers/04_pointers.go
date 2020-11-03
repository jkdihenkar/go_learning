package main

import (
	"fmt"
	"reflect"
)

/*
	basic pointers
*/

func main() {
	a := 5
	b := &a

	fmt.Println("Values = ", a, *b)
	fmt.Println("Address = ", b)
	fmt.Println("Type of pointer is - ", reflect.TypeOf(b))

	*b = 13
	fmt.Println("Changed value using pointer = ", a, *b)

	b = nil
	fmt.Println("Values = ", a, *b)

}

/* OUTPUT -

[jd@jdpc go_learning]$ go run 04_pointers/04_pointers.go
Values =  5 5
Address =  0xc420018138
Type of pointer is -  *int
Changed value using pointer =  13 13
panic: runtime error: invalid memory address or nil pointer dereference
[signal SIGSEGV: segmentation violation code=0x1 addr=0x0 pc=0x482348]

goroutine 1 [running]:
main.main()
	jkdihenkar/go_learning/04_pointers/04_pointers.go:24 +0x368
exit status 2

*/
