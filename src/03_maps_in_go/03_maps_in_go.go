package main

import (
	"encoding/json"
	"fmt"
)

/*
	map data type
	struct for map string -> struct
	inline declare maps
	range on maps demo
	len of map and delete in map

	map as json
*/

func main() {

	// Struct fieldnames must start with uppercase letter for JSON to see their value.

	type personalDetail struct {
		FirstName       string `json:"firstname"`
		LastName        string `json:"lastname"`
		FavouriteNumber int    `json:"favnum"`
		IsActive        bool   `json:"active"`
	}

	personalDetailsMap := make(map[string]personalDetail)

	personalDetailsMap["jkdihenkar@gmail.com"] = personalDetail{"Jay", "Dihenkar", 6, true}

	fmt.Println("Fetch personalDetails from Map - ", personalDetailsMap["jkdihenkar@gmail.com"])
	fmt.Println("Print only fav number - ", personalDetailsMap["jkdihenkar@gmail.com"].FavouriteNumber)

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

	// Print value formatting +%v
	fmt.Printf("Values for persDataAutoDeclare using printf perV :: %+v \n", persDataAutoDeclare)

	jsonString, err := json.Marshal(persDataAutoDeclare)
	//jsonString is a byteArray

	fmt.Println("JSON encoded output for MAP is # ", string(jsonString))
	fmt.Println("Error encountered while JSON encoding :: ", err)

}

/* OUTPUT
[jay@localhost go_learning]$ go run src/03_maps_in_go/03_maps_in_go.go
Fetch personalDetails from Map -  {Jay Dihenkar 6 true}
Print only fav number -  6
Latest map -  map[garb:{GarbFN GarbLN 99 false} jkdihenkar@gmail.com:{Jay Dihenkar 6 true}]  | Length= 2
jkdihenkar@gmail.com {Jay Dihenkar 6 true}
garb {GarbFN GarbLN 99 false}
Map after delete -  map[jkdihenkar@gmail.com:{Jay Dihenkar 6 true}]  | Length= 1
Autodeclare map -  map[jay:{Jay Dihenkar 6 true}]  | Length= 1
Values for persDataAutoDeclare using printf perV :: map[jay:{FirstName:Jay LastName:Dihenkar FavouriteNumber:6 IsActive:true}]
JSON encoded output for MAP is #  {"jay":{"firstname":"Jay","lastname":"Dihenkar","favnum":6,"active":true}}
Error encountered while JSON encoding ::  <nil>

[jay@localhost go_learning]$ go run src/03_maps_in_go/03_maps_in_go.go  | grep 'JSON encoded output' | cut -d'#' -f2 | jq .
{
  "jay": {
    "firstname": "Jay",
    "lastname": "Dihenkar",
    "favnum": 6,
    "active": true
  }
}

*/
