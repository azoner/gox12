package gox12

import (
	"strings"
)

type Segment struct {
	SegmentId string
	Elements  []string
}

func MakeSegment(line string, elementTerm byte, subelementTerm byte, repTerm byte) (segment Segment) {
	fields := strings.Split(line, string(elementTerm))
	segment.SegmentId = fields[0]
	segment.Elements = fields[1:]
	return
}
