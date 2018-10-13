package main

import "fmt"

/*
	map data type
	struct for map string -> struct
	inline declare maps
	range on maps demo
	len of map and delete in map
*/

func main() {

	type personalDetail struct {
		firstName       string
		lastName        string
		favouriteNumber int
		isActive        bool
	}

	personalDetailsMap := make(map[string]personalDetail)

	personalDetailsMap["jkdihenkar@gmail.com"] = personalDetail{"Jay", "Dihenkar", 6, true}

	fmt.Println("Fetch personalDetails from Map - ", personalDetailsMap["jkdihenkar@gmail.com"])
	fmt.Println("Print only fav number - ", personalDetailsMap["jkdihenkar@gmail.com"].favouriteNumber)

	personalDetailsMap["garb"] = personalDetail{"GarbFN", "GarbLN", 99, false}

	fmt.Println("Latest map - ", personalDetailsMap, " | Length=", len(personalDetailsMap))

	// test range on maps
	for k, v := range personalDetailsMap {
		fmt.Println(k, v)
	}

	delete(personalDetailsMap, "garb")

	fmt.Println("Map after delete - ", personalDetailsMap, " | Length=", len(personalDetailsMap))

	persDataAutoDeclare := map[string]personalDetail{"jay": personalDetail{"Jay", "Dihenkar", 6, true}}
	fmt.Println("Autodeclare map - ", persDataAutoDeclare, " | Length=", len(persDataAutoDeclare))

}

/* OUTPUT
[jd@jdpc go_learning]$ go run 03_maps_in_go/03_maps_in_go.go
Fetch personalDetails from Map -  {Jay Dihenkar 6 true}
Print only fav number -  6
Latest map -  map[jkdihenkar@gmail.com:{Jay Dihenkar 6 true} garb:{GarbFN GarbLN 99 false}]  | Length= 2
jkdihenkar@gmail.com {Jay Dihenkar 6 true}
garb {GarbFN GarbLN 99 false}
Map after delete -  map[jkdihenkar@gmail.com:{Jay Dihenkar 6 true}]  | Length= 1
Autodeclare map -  map[jay:{Jay Dihenkar 6 true}]  | Length= 1
*/
