package main

import (
	"fmt"
)

type Lookup func(string) (string, bool, error)

func MakeMapFinder() Lookup {
    var pathk = map[string]string{
        "ISA": "/ISA_LOOP",
        "GS": "GS_LOOP",
        "ST": "ST_LOOP",
    }
	return func(name string) (string, bool, error) {
        p, ok := pathk[name]
		return p, ok, nil
	}
}

func main() {
	//var gg string
	//gg = "adas;ldka;sldk;"
	f := MakeMapFinder()
    segs := [...]string{
        "ISA",
        "ST",
        "AAA",
    }
    for _, s := range segs {
        res, ok, _ := f(s)
        if ok {
	        fmt.Println(res)
        } else {
	        fmt.Println("Not found")
        }
    }
}
