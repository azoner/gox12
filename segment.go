package gox12

import (
	"strings"
)

type Segment struct {
	SegmentId string
	Elements  []string
}

func MakeSegment(line string, delim Delimiters) (segment Segment) {
	fields := strings.Split(line, string(delim.ElementTerm))
	segment.SegmentId = fields[0]
	segment.Elements = fields[1:]
	return
}
