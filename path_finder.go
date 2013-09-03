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

// is the segment "matched"
type segMatcher func(seg Segment) bool

func segmentMatchBySegmentId(segmentId string) segMatcher {
	return func(seg Segment) bool {
		return seg.SegmentId == segmentId
	}
}

func segmentMatchIdByPath(segmentId string, x12path string, id_value string) segMatcher {
	return func(seg Segment) bool {
		v, found, _ := seg.GetValue(x12path)
		return seg.SegmentId == segmentId && found && v == id_value
	}
}

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
