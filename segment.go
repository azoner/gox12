package gox12

import (
	"fmt"
	"strings"
)

type Segment struct {
	SegmentId  string
	Composites [][]string
}

func NewSegment(line string, elementTerm byte, subelementTerm byte, repTerm byte) Segment {
	fields := strings.Split(line, string(elementTerm))
	segmentId := fields[0]
	comps := make([][]string, len(fields)-1)
	for i, v := range fields[1:] {
		c := strings.Split(v, string(subelementTerm))
		comps[i] = c
	}
	return Segment{segmentId, comps}
}

// Acts like golang maps, if not found, returns default value with found==false
// X12 Path indices are 1-based
func (s *Segment) GetValue(x12path string) (val string, found bool, err error) {
	var xpath *X12Path
	if xpath, err = ParseX12Path(x12path); err != nil {
		return "", false, err
	}
	if xpath.SegmentId != "" && s.SegmentId != xpath.SegmentId {
		return "", false, fmt.Errorf("Looking for Segment ID[%s], mine is [%s]", xpath.SegmentId, s.SegmentId)
	}
	if xpath.ElementIdx == 0 {
		return "", false, fmt.Errorf("No element index specified for [%s]", x12path)
	}
	myEleIdx := xpath.ElementIdx - 1
	var mySubeleIdx int
	if xpath.SubelementIdx == 0 {
		if myEleIdx < len(s.Composites) && len(s.Composites[myEleIdx]) > 1 {
			return "", false, fmt.Errorf("This is a composite but no sub-element index was specified for [%s]", x12path)
		}
		mySubeleIdx = 0
	} else {
		mySubeleIdx = xpath.SubelementIdx - 1
	}
	if myEleIdx < len(s.Composites) {
		if mySubeleIdx < len(s.Composites[myEleIdx]) {
			return s.Composites[myEleIdx][mySubeleIdx], true, nil
		}
	}
	return "", false, nil
}

type ElementValue struct {
	X12Path X12Path
	Value   string
}

func (s *Segment) GetAllValues() <-chan ElementValue {
	ch := make(chan ElementValue)
	go func() {
		for i, comp := range s.Composites {
			for j, elem := range comp {
				x12path := X12Path{SegmentId: s.SegmentId, ElementIdx: i + 1, SubelementIdx: j + 1}
				ev := ElementValue{x12path, elem}
				//ch <- new(ElementValue{new(X12Path{SegmentId: s.SegmentId, ElementIdx: i+1, SubelementIdx: j+1}), elem})
				ch <- ev
			}
		}
		close(ch)
	}()
	return ch
}

//func splitComposite(f2 string, term string) (ret []string) {
//	ret = strings.Split(f2, term)
//	return
//}

//type Composite []string
