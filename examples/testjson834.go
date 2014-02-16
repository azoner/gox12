package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type PathTarget struct {
	FoundPath string
	Segid     string
	F         string `json:"type"`
}

func main() {
	file, e := ioutil.ReadFile("./834.used.json")
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}
	/*	var json2 = []byte(`{
		  "DETAIL/2000/2100A/DMG": [
		    {
		      "Newpath": "DETAIL/2000/2100G/NM1",
		      "Segid": "NM1",
		      "F": "seg"
		    }
		  ],
		  "DETAIL/2000/2100A/LUI": [
		    {
		      "Newpath": "DETAIL/2000/2100G/NM1",
		      "Segid": "NM1",
		      "F": "seg"
		    }
		  ]
			}`)
	*/
	var targets map[string][]PathTarget
	err := json.Unmarshal(file, &targets)
	if err != nil {
		fmt.Println("error:", err)
	}
	for k, v := range targets {
		fmt.Println(k)
		for _, p := range v {
			//fmt.Print(p["newpath"])
			//fmt.Print(p["segid"])
			fmt.Printf("\t%+v\n", p)
		}
	}
}
