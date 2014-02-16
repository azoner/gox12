package gox12

import ()

// Given the current location, find the path of the new segment
type X12PathFinder interface {
	FindNext(x12Path string, segment Segment) (foundPath string, found bool, err error)
}

// Hardcoded lookups for standard X12 structure wrappers
type HeaderPathFinder struct {
	hardMap map[string]string
}

func NewHeaderMapFinder() *HeaderPathFinder {
	f := new(HeaderPathFinder)
	f.hardMap = map[string]string{
		"ISA": "/ISA_LOOP/ISA",
		"IEA": "/ISA_LOOP/IEA",
		"GS":  "/ISA_LOOP/GS_LOOP/GS",
		"GE":  "/ISA_LOOP/GS_LOOP/GE",
		"ST":  "/ISA_LOOP/GS_LOOP/ST_LOOP/ST",
		"SE":  "/ISA_LOOP/GS_LOOP/ST_LOOP/SE",
	}
	return f
}

func (finder *HeaderPathFinder) FindNext(x12Path string, segment Segment) (foundPath string, found bool, err error) {
	segId := segment.SegmentId
	p, ok := finder.hardMap[segId]
	if ok {
		return p, ok, nil
	}
	return "", false, nil
}

type PathFinder func(string, Segment) (string, bool, error)

type EmptyPath struct {
	Path string
}

//func (e *EmptyPath) Run2 PathFinder {
//    return "", true, nil
//}

// this is the method signature
// need to close lookup maps
func findPath(rawpath string, seg Segment) (foundPath string, ok bool, err error) {
	return "", true, nil
}

// segMatcher is the function signature for segment matcher
// is the segment "matched"
type segMatcher func(seg Segment) bool

// segmentMatchBySegmentId matches a segment only by the segment ID
func segmentMatchBySegmentId(segmentId string) segMatcher {
	return func(seg Segment) bool {
		return seg.SegmentId == segmentId
	}
}

// segmentMatchIdByPath matches a segment by the segment ID and the ID value of the
// element at the x12path
func segmentMatchIdByPath(segmentId string, x12path string, id_value string) segMatcher {
	return func(seg Segment) bool {
		v, found, _ := seg.GetValue(x12path)
		return seg.SegmentId == segmentId && found && v == id_value
	}
}

// segmentMatchIdByPath matches a segment by the segment ID and one of the ID value of the
// element at the x12path
func segmentMatchIdListByPath(segmentId string, x12path string, id_list []string) segMatcher {
	return func(seg Segment) bool {
		v, found, _ := seg.GetValue(x12path)
		x := stringInSlice(v, id_list)
		return seg.SegmentId == segmentId && found && x
	}
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
