package main

import "fmt"

/*
	map data type
	struct for map string -> struct
	inline declare maps
	range on maps demo
	value reciever method of struct
	pointer reciever method of struct
*/

type personalDetail struct {
	firstName, lastName string
	favouriteNumber     int
	isActive            bool
}

// value reciever method
func (p personalDetail) toStringRep() string {
	return fmt.Sprintf("Name is %s %s. FavNum=%d. Active=%t", p.firstName, p.lastName, p.favouriteNumber, p.isActive)
}

// pointer reciever method
func (p *personalDetail) makeInActive() {
	p.isActive = false
}

// pointer reciever method
func (p *personalDetail) ChangeFavNumber(favnum int) {
	p.favouriteNumber = favnum
}

func main() {

	persDataAutoDeclare := map[string]personalDetail{"jay": personalDetail{"Jay", "Dihenkar", 6, true}}

	for k, v := range persDataAutoDeclare {
		fmt.Println(k, v.toStringRep())
	}

	mydetails := persDataAutoDeclare["jay"]
	mydetails.makeInActive()

	mydetails.ChangeFavNumber(99)

	fmt.Println(mydetails.toStringRep())

}

/*

[jd@jdpc go_learning]$ go run 06_structs_and_methods/06_structs_and_methods.go
jay Name is Jay Dihenkar. FavNum=6. Active=true
Name is Jay Dihenkar. FavNum=99. Active=false

*/
