package main

import (
	"fmt"
	"regexp"
)

func main() {
	myWordValidator, _ := regexp.Compile("\\[\\w+]")
	myTextOnlyValidator, _ := regexp.Compile("\\[[a-z]+]")

	myCases := []string{"[foo]", "[foo=bar]", "[fufu]", "[32fu]"}
	fmt.Println("Set of test cases = ", myCases)

	fmt.Println("Matching cases against: ", myWordValidator)
	for _, tcase := range myCases {
		match := myWordValidator.MatchString(tcase)
		fmt.Println("Match for ", tcase, " is ... ", match)
	}

	fmt.Println("Matching cases against: ", myTextOnlyValidator)
	for _, tcase := range myCases {
		match := myTextOnlyValidator.MatchString(tcase)
		fmt.Println("Match for ", tcase, " is ... ", match)
	}
}

/*

[jd@jdpc go_learning]$ go run 11_regex/11_regex.go
Set of test cases =  [[foo] [foo=bar] [fufu] [32fu]]
Matching cases against:  \[\w+]
Match for  [foo]  is ...  true
Match for  [foo=bar]  is ...  false
Match for  [fufu]  is ...  true
Match for  [32fu]  is ...  true
Matching cases against:  \[[a-z]+]
Match for  [foo]  is ...  true
Match for  [foo=bar]  is ...  false
Match for  [fufu]  is ...  true
Match for  [32fu]  is ...  false


*/
