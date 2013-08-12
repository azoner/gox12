gox12
===

golang X12 non-validating parser

[![Build Status](https://travis-ci.org/azoner/gox12.png)](https://travis-ci.org/azoner/gox12)

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
                fmt.Println(rs)
        }
}

```
