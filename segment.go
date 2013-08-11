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

func NewSegment(line string, elementTerm byte, subelementTerm byte, repTerm byte) Segment {
	fields := strings.Split(line, string(elementTerm))
    segmentId := fields[0]
	comps := make([][]string, len(fields)-1)
	for i, v := range fields[1:] {
		c := strings.Split(v, string(subelementTerm))
		t := make([]string, cap(c))
		copy(t, c)
		comps[i] = t
	}
	return Segment{segmentId, comps}
}

func splitComposite(f2 string, term string) (ret []string) {
	ret = strings.Split(f2, term)
	return
}

func (s *Segment) GetValue(x12path string) (val string, err error) {
	return
}
