package gox12

import (
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
		t.Errorf("NewRawX12FileReader failed")
	}
	//expectedDelimeters := Delimiters{'+', '&', '!', 0}
	expectedSegTerm := '+'
	expectedElementTerm := '&'
	expectedSubelementTerm := '!'
	expectedRepetitionTerm := 0
	if raw.segmentTerm != byte(expectedSegTerm) {
		t.Errorf("Didn't get expected result [%c], instead got [%c]", expectedSegTerm, raw.segmentTerm)
	}
	if raw.elementTerm != byte(expectedElementTerm) {
		t.Errorf("Didn't get expected result [%c], instead got [%c]", expectedElementTerm, raw.elementTerm)
	}
	if raw.subelementTerm != byte(expectedSubelementTerm) {
		t.Errorf("Didn't get expected result [%c], instead got [%c]", expectedSubelementTerm, raw.subelementTerm)
	}
	if raw.repetitionTerm != byte(expectedRepetitionTerm) {
		t.Errorf("Didn't get expected result [%c], instead got [%c]", expectedRepetitionTerm, raw.repetitionTerm)
	}
	expectedCount := 7
	ct := 0
	for _ = range raw.GetSegments() {
		ct += 1
	}
	if ct != expectedCount {
		t.Errorf("Didn't get expected segment count of %d, instead got %d", expectedCount, ct)
	}
}

func TestArbitraryDelimiters5010(t *testing.T) {
	str1 := "ISA&00&          &00&          &ZZ&ZZ000          &ZZ&ZZ001          &030828&1128&^&00501&000010121&0&T&!+\n"
	str1 += "GS&HC&ZZ000&ZZ001&20030828&1128&17&X&005010X223+\n"
	str1 += "ST&837&11280001+\n"
	str1 += "TST&AA!1!1&BB!5+\n"
	str1 += "SE&3&11280001+\n"
	str1 += "GE&1&17+\n"
	str1 += "IEA&1&000010121+\n"
	inFile := strings.NewReader(str1)
	raw, err := NewRawX12FileReader(inFile)
	if err != nil {
		t.Errorf("NewRawX12FileReader failed")
	}
	//expectedDelimeters := Delimiters{'+', '&', '!', 0}
	expectedSegTerm := '+'
	expectedElementTerm := '&'
	expectedSubelementTerm := '!'
	expectedRepetitionTerm := '^'
	if raw.segmentTerm != byte(expectedSegTerm) {
		t.Errorf("Didn't get expected result [%c], instead got [%c]", expectedSegTerm, raw.segmentTerm)
	}
	if raw.elementTerm != byte(expectedElementTerm) {
		t.Errorf("Didn't get expected result [%c], instead got [%c]", expectedElementTerm, raw.elementTerm)
	}
	if raw.subelementTerm != byte(expectedSubelementTerm) {
		t.Errorf("Didn't get expected result [%c], instead got [%c]", expectedSubelementTerm, raw.subelementTerm)
	}
	if raw.repetitionTerm != byte(expectedRepetitionTerm) {
		t.Errorf("Didn't get expected result [%c], instead got [%c]", expectedRepetitionTerm, raw.repetitionTerm)
	}
	expectedCount := 7
	ct := 0
	for _ = range raw.GetSegments() {
		ct += 1
	}
	if ct != expectedCount {
		t.Errorf("Didn't get expected segment count of %d, instead got %d", expectedCount, ct)
	}
}

//str1 = strings.Replace(str1, "&", "\x1C", -1)

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
	ct := 0
	for _ = range raw.GetSegments() {
		ct += 1
	}
	if ct != 7 {
		t.Errorf("Didn't get expected segment count of %d, instead got %d", 7, ct)
	}
}
