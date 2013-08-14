package main

import (
	"encoding/json"
	"fmt"
)

type Animal struct {
	Name  string
	Order string
}

func main() {
	var jsonBlob = []byte(`[
		{"Name": "Platypus", "Order": "Monotremata"},
		{"Name": "Quoll",    "Order": "Dasyuromorphia"}
	]`)
	var json2 = []byte(`{
		"Maps": {
			"DETAIL/2000/2100A/DMG": [
			{
				"newpath": "DETAIL/2000/2100G/NM1",
				"segid": "NM1",
				"f": "seg"
			}
			],
			"DETAIL/2000/2100A/LUI": [
			{
				"newpath": "DETAIL/2000/2100G/NM1",
				"segid": "NM1",
				"f": "seg"
			}
			]
		}
	}`)
	var animals []Animal
	err := json.Unmarshal(jsonBlob, &animals)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("%+v", animals)
	fmt.Println(json2)
}
