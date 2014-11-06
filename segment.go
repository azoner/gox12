package gox12

import (
	"bytes"
	"fmt"
	"strings"
)

type Segment struct {
	SegmentId  string
	Composites []Composite
}

type Composite []string

type ElementValue struct {
	X12Path X12Path
	Value   string
}

func NewSegment(line string, elementTerm byte, subelementTerm byte, repTerm byte) Segment {
	fields := strings.Split(line, string(elementTerm))
	segmentId := fields[0]
	comps := make([]Composite, len(fields)-1)
	for i, v := range fields[1:] {
		c := strings.Split(v, string(subelementTerm))
		comps[i] = c
	}
	return Segment{segmentId, comps}
}

// GetValue returns the string value of the simple element at the x12path
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

func (s *Segment) SetValue(x12path, val string) (err error) {
	var xpath *X12Path
	if xpath, err = ParseX12Path(x12path); err != nil {
		return err
	}
	if xpath.SegmentId != "" && s.SegmentId != xpath.SegmentId {
		return fmt.Errorf("Looking for Segment ID[%s], mine is [%s]", xpath.SegmentId, s.SegmentId)
	}
	if xpath.ElementIdx == 0 {
		return fmt.Errorf("No element index specified for [%s]", x12path)
	}
	myEleIdx := xpath.ElementIdx - 1
	var mySubeleIdx int
	if xpath.SubelementIdx == 0 {
		if myEleIdx < len(s.Composites) && len(s.Composites[myEleIdx]) > 1 {
			return fmt.Errorf("This is a composite but no sub-element index was specified for [%s]", x12path)
		}
		mySubeleIdx = 0
	} else {
		mySubeleIdx = xpath.SubelementIdx - 1
	}
	if myEleIdx < len(s.Composites) {
		if mySubeleIdx < len(s.Composites[myEleIdx]) {
			s.Composites[myEleIdx][mySubeleIdx] = val
		}
	}
	return nil
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

// Default formatting
func (s *Segment) String() string {
	return s.Format('*', ':', '^')
}

func (s *Segment) Format(elementTerm byte, subelementTerm byte, repTerm byte) string {
	var buf bytes.Buffer
	buf.WriteString(s.SegmentId)
	for _, comp := range s.Composites {
		buf.WriteByte(elementTerm)
		buf.WriteString(formatComposite(comp, subelementTerm, repTerm))
	}
	return buf.String()
}

func formatComposite(c Composite, subelementTerm byte, repTerm byte) string {
	return strings.Join(c, string(subelementTerm))
}

//func splitComposite(f2 string, term string) (ret []string) {
//	ret = strings.Split(f2, term)
//	return
//}
