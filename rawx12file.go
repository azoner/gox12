package gox12

import (
	"bufio"
	//"bytes"
	//"fmt"
	"io"
	"log"
	"strings"
)

func NewRawX12FileReader(inFile io.Reader) (*rawX12FileReader, error) {
	const isaLength = 106
	r := new(rawX12FileReader)
	r.reader = bufio.NewReader(inFile)
	//buffer := bytes.NewBuffer(make([]byte, 0))
	first, err := r.reader.Peek(isaLength)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	isa := string(first)
	segTerm, eleTerm, subeleTerm, repTerm := getDelimiters(isa)
	r.segmentTerm = segTerm
	r.elementTerm = eleTerm
	r.subelementTerm = subeleTerm
	r.repetitionTerm = repTerm
	return r, nil
}

func (r *rawX12FileReader) GetSegments() <-chan RawSegment {
	ch := make(chan RawSegment)
	ct := 0
	go func() {
		for {
			row, err := r.reader.ReadString(r.segmentTerm)
			if err == io.EOF {
				break
			} else if err != nil {
				panic(err)
			}
			if strings.HasSuffix(row, string(r.segmentTerm)) {
				row = row[:len(row)-1]
			}
			row = strings.Trim(row, "\r\n")
			mySeg := MakeSegment(row, r.elementTerm, r.subelementTerm, r.repetitionTerm)
			ct++
			seg := RawSegment{
				mySeg,
				ct,
			}
			ch <- seg
		}
		close(ch)
	}()
	return ch
}

// X12 line reader and part delimiters
type rawX12FileReader struct {
	reader         *bufio.Reader
	segmentTerm    byte
	elementTerm    byte
	subelementTerm byte
	repetitionTerm byte
}

type RawSegment struct {
	Segment   Segment
	LineCount int
}

// Get the X12 delimiters specified in the ISA segment
func getDelimiters(isa string) (segTerm byte, eleTerm byte, subeleTerm byte, repTerm byte) {
	segTerm = isa[len(isa)-1]
	eleTerm = isa[3]
	subeleTerm = isa[len(isa)-2]
	if isa[84:89] == "005010" {
		repTerm = isa[82]
	}
	return
}
