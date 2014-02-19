gox12
===

golang X12 non-validating parser

[![Build Status](https://travis-ci.org/azoner/gox12.png)](https://travis-ci.org/azoner/gox12)
<!--- [![Coverage Status](https://coveralls.io/repos/azoner/gox12/badge.png?branch=pathfinder)](https://coveralls.io/r/azoner/gox12?branch=pathfinder) -->
Installation
------------

  go get github.com/azoner/gox12


Example
-----

```go
package main

import (
        "fmt"
        "os"
        "log"
        "github.com/azoner/gox12"
)

func main() {
        inFilename := "x12file.txt"
        inFile, err := os.Open(inFilename)
        if err != nil {
                log.Fatal(err)
                os.Exit(1)
        }
        defer inFile.Close()
        raw, err := gox12.NewRawX12FileReader(inFile)
        if err != nil {
                fmt.Println(err)
        }
        for rs := range raw.GetSegments() {
                if rs.Segment.SegmentId == "INS" {
                        fmt.Println(rs)
                        v, _, _ := rs.Segment.GetValue("INS01")
                        fmt.Println(v)
                        for v := range rs.Segment.GetAllValues() {
                                fmt.Println(v.X12Path, v.Value)
                        }
                        fmt.Println()
                }
        }
}

```
