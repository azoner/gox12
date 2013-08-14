package main

import (
	"fmt"
	"github.com/azoner/gox12"
)

type Lookup func(string, gox12.Segment) (string, bool, error)

func MakeMapFinder() Lookup {
	var hardMap = map[string]string{
		"ISA": "/ISA_LOOP/ISA",
		"IEA": "/ISA_LOOP/IEA",
		"GS":  "/ISA_LOOP/GS_LOOP/GS",
		"GE":  "/ISA_LOOP/GS_LOOP/GE",
		"ST":  "/ISA_LOOP/GS_LOOP/ST_LOOP/ST",
		"SE":  "/ISA_LOOP/GS_LOOP/ST_LOOP/SE",
	}
	return func(rawpath string, s gox12.Segment) (string, bool, error) {
		segId := s.SegmentId
		p, ok := hardMap[segId]
		if ok {
			return p, ok, nil
		}
		return "", false, nil
	}
}

//func findPath(rawpath string, seg Segment) (newpath string, ok bool, err error) {
//}

func main() {
	//var gg string
	//gg = "adas;ldka;sldk;"
	finders := make([]Lookup, 0)
	f := MakeMapFinder()
	finders = append(finders, f)
	segs := [...]string{
		"ISA",
		"ST",
		"AAA",
	}
	for _, s := range segs {
		seg := gox12.Segment{SegmentId: s}
		for _, f2 := range finders {
			res, ok, _ := f2(s, seg)
			if ok {
				fmt.Println(res)
				break
			}
		}
		fmt.Println("Not found")
	}
	return
}

/*
func testmap() {
	s := []int{1, 2, 3}
	f := func(x int) { return x * 2 }
	for i, _ := range s {
		s[i] = f(s[i])
	}
}
*/
