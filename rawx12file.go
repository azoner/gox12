package gox12

import (
	"bufio"
	//"bytes"
	"fmt"
	"io"
	"log"
	"strings"
)

const isa_length = 106

type rawX12FileReader struct {
    reader *Reader
	SegmentTerm    byte
	ElementTerm    byte
	SubelementTerm byte
	RepetitionTerm byte
}

func NewRawX12FileReader(inFile io.Reader) *rawX12FileReader, error {
    r := new(rawX12FileReader)
	r.reader := bufio.NewReader(inFile)
	//buffer := bytes.NewBuffer(make([]byte, 0))
	first, err := reader.Peek(isa_length)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	isa := string(first)
	delim := getDelimiters(isa)
    r.SegmentTerm = delim.SegmentTerm
    r.ElementTerm = delim.ElementTerm
    r.SubelementTerm = delim.SubelementTerm
    r.RepetitionTerm = delim.RepetitionTerm
    return r, nil
}

func (r *rawX12FileReader) Iter() <-chan RawSegment) {
    ch := make(chan RawSegment)
	ct := 0
	for {
		row, err := r.reader.ReadString(delim.SegmentTerm)
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
	return ch
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
