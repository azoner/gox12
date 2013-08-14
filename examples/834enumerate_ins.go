package main

import (
	"fmt"
	"github.com/azoner/gox12"
	"log"
	"os"
)

func main() {
	inFilename := "test834.txt"
	inFile, err := os.Open(inFilename)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer inFile.Close()
	raw, err := gox12.NewRawX12FileReader(inFile)
	if err != nil {
		fmt.Println(err)
	}
	for rs := range raw.GetSegments() {
		if rs.Segment.SegmentId == "INS" {
			fmt.Println(rs)
			v, _, _ := rs.Segment.GetValue("INS01")
			fmt.Println(v)
			for v := range rs.Segment.GetAllValues() {
				fmt.Println(v.X12Path, v.Value)
			}
			fmt.Println()
		}
	}
}
