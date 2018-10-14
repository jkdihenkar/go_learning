package unsafe_example

import (
	"fmt"
	"unsafe"
)

func MainUnsafeDemo() {
	var myNumber int64 = 456888777994

	var myNumberSafePtr = &myNumber
	var myNumberUnsafePtr = (*int32)(unsafe.Pointer(myNumberSafePtr))

	fmt.Println("Safe print -> ", *myNumberSafePtr)
	fmt.Println("Unsafe print -> ", *myNumberUnsafePtr)
}
