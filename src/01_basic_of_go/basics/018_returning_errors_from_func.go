package basics

import (
	"fmt"
)

func TestMultipleReturns() {

	fmt.Println(divTwoNos(2, 5))
	fmt.Println(divTwoNos(9, 4))
	fmt.Println(divTwoNos(10999, 0))

}

func divTwoNos(a, b int) (float64, error) {
	// Smaller case begin letter in d in div implies unexported function
	if b == 0 {
		return 0, fmt.Errorf("divisor can never be zero")
	}
	return float64(a) / float64(b), nil
}
