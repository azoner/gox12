package gox12

import (
	"fmt"
	"log"
	"os"
	"strings"
	"testing"
)

// AAA
func TestArbitraryDelimiters(t *testing.T) {
	str1 := "ISA&00&          &00&          &ZZ&ZZ000          &ZZ&ZZ001          &030828&1128&U&00401&000010121&0&T&!+\n"
	str1 += "GS&HC&ZZ000&ZZ001&20030828&1128&17&X&004010X098+\n"
	str1 += "ST&837&11280001+\n"
	str1 += "TST&AA!1!1&BB!5+\n"
	str1 += "SE&3&11280001+\n"
	str1 += "GE&1&17+\n"
	str1 += "IEA&1&000010121+\n"
	inFile := strings.NewReader(str1)
	raw, err := NewRawX12FileReader(inFile)
	if err != nil {
		fmt.Println(err)
	}
	//expectedDelimeters := Delimiters{'+', '&', '!', 0}
	expectedSegTerm := '+'
	// segmentTerm
	// elementTerm
	// subelementTerm
	// repetitionTerm
	if raw.segmentTerm != byte(expectedSegTerm) {
		t.Errorf("Didn't get expected result [%s], instead got [%s]", expectedSegTerm, raw.segmentTerm)
	}
	for s := range raw.GetSegments() {
		fmt.Println(s)
	}
}

func testParse834(t *testing.T) {
	inFilename := "test834.txt"
	//inFile *os.File
	//inFile io.Reader
	inFile, err := os.Open(inFilename)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer inFile.Close()
	raw, err := NewRawX12FileReader(inFile)
	for seg := range raw.GetSegments() {
		fmt.Println(seg)
	}
}
