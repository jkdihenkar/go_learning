package main

// #cgo CFLAGS: -g -Wall
// #include <stdlib.h>
// #include "greeter.h"
import "C"
import (
	"fmt"
	"unsafe"

	"github.com/jkdihenkar/go_learning/12_calling_c_from_go/unsafe_example"
)

func main() {
	unsafe_example.MainUnsafeDemo()
	normalCGOInterfacing()
	structCGOInterfacing()
}

func structCGOInterfacing() {
	name := C.CString("Gopher")
	defer C.free(unsafe.Pointer(name))

	year := C.int(2018)

	g := C.struct_Greetee{
		name: name,
		year: year,
	}

	ptr := C.malloc(C.sizeof_char * 1024)
	defer C.free(unsafe.Pointer(ptr))

	size := C.greet_struct(&g, (*C.char)(ptr))

	b := C.GoBytes(ptr, size)
	fmt.Println(string(b))
}

func normalCGOInterfacing() {
	name := C.CString("Gopher")
	defer C.free(unsafe.Pointer(name))

	year := C.int(2018)

	ptr := C.malloc(C.sizeof_char * 1024)
	defer C.free(unsafe.Pointer(ptr))

	size := C.greet(name, year, (*C.char)(ptr))

	b := C.GoBytes(ptr, size)
	fmt.Println(string(b))
}

/*

[jd@jdpc go_learning]$ cd 12_calling_c_from_go/
[jd@jdpc 12_calling_c_from_go]$
[jd@jdpc 12_calling_c_from_go]$
[jd@jdpc 12_calling_c_from_go]$ gcc -c greeter.c
[jd@jdpc 12_calling_c_from_go]$
[jd@jdpc 12_calling_c_from_go]$ go build
[jd@jdpc 12_calling_c_from_go]$ go install
[jd@jdpc 12_calling_c_from_go]$
[jd@jdpc 12_calling_c_from_go]$ $GOPATH/bin/12_calling_c_from_go
Safe print ->  456888777994
Unsafe print ->  1622244618
Greetings, Gopher from 2018! We come in peace :)
Greetings, Gopher from 2018! We come in peace :)

*/
