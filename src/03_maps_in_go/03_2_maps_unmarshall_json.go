package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	// Struct fieldnames must start with uppercase letter for JSON to see their value.

	type personalDetail struct {
		FirstName       string `json:"firstname"`
		LastName        string `json:"lastname,omitempty"`
		FavouriteNumber int    `json:"favnum"`
		IsActive        bool   `json:"active"`
	}

	jsonString := `{"firstname":"Jay","lastname":"Dihenkar","favnum":6,"active":true}`

	var sampleObject personalDetail
	err := json.Unmarshal([]byte(jsonString), &sampleObject)

	fmt.Printf("Parsed JSON as MAP :: %+v\n", sampleObject)

	if err != nil {
		fmt.Println(err)
	}

	var sampleObjectEmpty []byte
	sampleObjectEmpty, err = json.Marshal(personalDetail{FirstName: "Hay", LastName: "", FavouriteNumber: 17, IsActive: false})
	fmt.Printf("Parsed output Empty Fields example :: %s\n", string(sampleObjectEmpty))

	if err != nil {
		fmt.Println(err)
	}

	// Reading unstructured data
	var unstrucData map[string]interface{}

	unstrucDataJSONStr := `{"hello": "world", "number": 3, "boole": false}`

	err = json.Unmarshal([]byte(unstrucDataJSONStr), &unstrucData)

	fmt.Printf("Parsed unstruc JSON as MAP :: %+v\n", unstrucData)
	fmt.Printf("What is hello? :: %s\n", unstrucData["hello"])

	if err != nil {
		fmt.Println(err)
	}

}

/* SAMPLE OUTPUT

[jay@localhost go_learning]$ go run src/03_maps_in_go/03_2_maps_unmarshall_json.go
Parsed JSON as MAP :: {FirstName:Jay LastName:Dihenkar FavouriteNumber:6 IsActive:true}
Parsed output :: {"firstname":"Hay","favnum":17,"active":false}
Parsed unstruc JSON as MAP :: map[boole:false hello:world number:3]
What is hello? :: world

*/
