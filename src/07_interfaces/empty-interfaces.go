package main

import (
	"fmt"
)

// How do you check a variable type at runtime

func do(i interface{}) {
	// empty interfaces take everything as input
	// as all the types implements a nothing interface{}
	// interface{} has no defined methods
	switch v := i.(type) {
	case int:
		fmt.Printf("KNOWN INT => %v is of type %T || %v, %T\n", i, i, v, v)
	case string:
		fmt.Printf("KNOWN STRING => %v is of type %T || %v, %T\n", i, i, v, v)
	default:
		fmt.Printf("UNKNOWN => %v is of type %T || %v, %T\n", i, i, v, v)
	}
}

func main() {
	do(42)
	do("something")
	do(7.443)
}
