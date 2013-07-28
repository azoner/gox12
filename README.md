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
  "github.com/azoner/gox12"
  "log"
	"os"
  "fmt"
)

func main() {
	file, err := os.Open("testx12.txt")
	if err != nil {
		log.Fatal(err)
    os.Exit(1)
	}
  defer file.Close()
  ch := make(chan RawSegment)
  go ReadSegmentLines(file, ch)
  for row := range ch {
    fmt.Println(row)
  }
}
```
