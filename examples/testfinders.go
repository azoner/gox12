package main

import (
	"fmt"
	"gox12"
)

func main() {
	finder := gox12.NewFirstMatchPathFinder(gox12.NewHeaderMapFinder())

	segs := [...]string{
		"ISA",
		"ST",
		"AAA",
	}
	for _, s := range segs {
		seg := gox12.Segment{SegmentId: s}
		res, ok, _ := finder.FindNext(s, seg)
		if ok {
			fmt.Println(res)

		} else {
			fmt.Println("Not found " + s)
		}
	}
	return
}
