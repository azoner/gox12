package gox12

import ()

type PathFinder func(string, Segment) (string, bool, error)

func MakeMapFinder() PathFinder {
	var hardMap = map[string]string{
		"ISA": "/ISA_LOOP/ISA",
		"IEA": "/ISA_LOOP/IEA",
		"GS":  "/ISA_LOOP/GS_LOOP/GS",
		"GE":  "/ISA_LOOP/GS_LOOP/GE",
		"ST":  "/ISA_LOOP/GS_LOOP/ST_LOOP/ST",
		"SE":  "/ISA_LOOP/GS_LOOP/ST_LOOP/SE",
	}
	return func(rawpath string, s Segment) (string, bool, error) {
		segId := s.SegmentId
		p, ok := hardMap[segId]
		if ok {
			return p, ok, nil
		}
		return "", false, nil
	}
}
