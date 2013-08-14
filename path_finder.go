package gox12

import (
)

//type X12PathFinder interface {
//	FindNext(x12Path string, segment Segment) (newpath string, found bool, err error)
//}

type PathFinder func(string, Segment) (string, bool, error)

type EmptyPath struct {
    Path string
}

func (e *EmptyPath) Run2 PathFinder {
    return "", true, nil
}

//func MakeFinder() func() {
//    f := func(x12Path string, segment Segment) {
//    }
//    return f
//}

//func makeFinderFunction() (func(string, Segment) string, bool, error) {
//    return 
//}
