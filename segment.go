package gox12

import (
	"strings"
    //"fmt"
)

type Segment struct {
	SegmentId string
	//Composites []Composite
    Composites [][]string
}

//type Composite []string

func NewSegment(line string, elementTerm byte, subelementTerm byte, repTerm byte) (segment Segment) {
	fields := strings.Split(line, string(elementTerm))
    segment.SegmentId = fields[0]
    comps := make([][]string, len(fields)-1)
    for i, v := range fields[1:] {
        c := strings.Split(v, string(subelementTerm))
        comps[i] = c
    }
    segment.Composites = comps
    return
}


func splitComposite(f2 string, term string) (ret []string) {
    ret = strings.Split(f2, term)
    return
}
