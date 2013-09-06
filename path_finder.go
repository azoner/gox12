package gox12

import ()

//type X12PathFinder interface {
//	FindNext(x12Path string, segment Segment) (newpath string, found bool, err error)
//}

type PathFinder func(string, Segment) (string, bool, error)

type EmptyPath struct {
	Path string
}

// map[start_x12path] []struct {match_func func(seg) bool, newpath string}

//testLookups := map[string]

//func (e *EmptyPath) Run2 PathFinder {
//    return "", true, nil
//}

//func MakeFinder() func() {
//    f := func(x12Path string, segment Segment) {
//    }
//    return f
//}

//func makeFinderFunction() (func(string, Segment) string, bool, error) {
//    return
//}

// this is the method signature
// need to close lookup maps
func findPath(rawpath string, seg Segment) (newpath string, ok bool, err error) {
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
