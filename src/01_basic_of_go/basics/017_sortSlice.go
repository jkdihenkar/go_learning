package basics

import (
	"fmt"
	"sort"
)

type personalScore struct {
	name                  string
	plusPoint, minusPoint int
}

func MainSortSliceDemo() {
	mySlice := make([]personalScore, 0)

	mySlice = append(mySlice, personalScore{name: "JayD", plusPoint: 30, minusPoint: 5})
	mySlice = append(mySlice, personalScore{name: "Linux", plusPoint: 80, minusPoint: 25})
	mySlice = append(mySlice, personalScore{name: "Mac", plusPoint: 80, minusPoint: 35})
	mySlice = append(mySlice, personalScore{name: "Windows", plusPoint: 80, minusPoint: 95})

	fmt.Println("Original slice - ")
	fmt.Println(mySlice)

	fmt.Println("Sort Ascending by points - ")
	sort.Slice(mySlice, func(i, j int) bool {
		return mySlice[i].plusPoint-mySlice[i].minusPoint < mySlice[j].plusPoint-mySlice[j].minusPoint
	})
	fmt.Println(mySlice)

	fmt.Println("Sort Desc by points - ")
	sort.Slice(mySlice, func(i, j int) bool {
		return mySlice[i].plusPoint-mySlice[i].minusPoint > mySlice[j].plusPoint-mySlice[j].minusPoint
	})
	fmt.Println(mySlice)
}

/*

[jd@jdpc go_learning]$ go run 01_basic_of_go/01_basic_of_go.go 017
Original slice -
[{JayD 30 5} {Linux 80 25} {Mac 80 35} {Windows 80 95}]
Sort Ascending by points -
[{Windows 80 95} {JayD 30 5} {Mac 80 35} {Linux 80 25}]
Sort Desc by points -
[{Linux 80 25} {Mac 80 35} {JayD 30 5} {Windows 80 95}]


*/
