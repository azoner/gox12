package gox12

import (
	"bufio"
	//"bytes"
	"fmt"
	"io"
	"log"
	"strings"
)

func ReadSegmentLines(inFile io.Reader, ch chan RawSegment) {
	reader := bufio.NewReader(inFile)
	//buffer := bytes.NewBuffer(make([]byte, 0))
	first, err := reader.Peek(106)
	if err != nil {
		log.Fatal(err)
		return
	}
	isa := string(first)
	delim := getDelimiters(isa)
	fmt.Println(delim)
	ct := 0
	for {
		row, err := reader.ReadString(delim.SegmentTerm)
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
		row = strings.Trim(row, "~\r\n")
		mySeg := MakeSegment(row, delim)
		ct++
		seg := RawSegment{
			mySeg,
			delim,
			ct,
		}
		ch <- seg
	}
	close(ch)
}

type RawSegment struct {
	Segment     Segment
	Terminators Delimiters
	LineCount   int
}

type Delimiters struct {
	SegmentTerm    byte
	ElementTerm    byte
	SubelementTerm byte
	RepetitionTerm byte
}

func getDelimiters(isa string) Delimiters {
	d := Delimiters{
		isa[len(isa)-1],
		isa[3],
		isa[len(isa)-2],
		0,
	}
	if isa[84:89] == "005010" {
		d.SubelementTerm = isa[82]
	}
	return d
}
