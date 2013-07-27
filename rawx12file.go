package main

import (
    "os"
    "bufio"
    //"bytes"
    "io"
    "fmt"
    "strings"
    "log"
)

func main() {
    inFilename := "test834.txt"
    //inFile *os.File
    //inFile io.Reader
    inFile, err := os.Open(inFilename)
    if err != nil {
        log.Fatal(err)
        os.Exit(1)
    }
    defer inFile.Close()
    ch := make(chan string)
    go ReadSegmentLines(inFile, ch)
    for row := range ch {
        fmt.Println(row)
    }
}

func ReadSegmentLines(inFile io.Reader, ch chan string) {
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
    for {
        row, err := reader.ReadString(delim.SegmentTerm)
        if err == io.EOF {
            break
        } else if err != nil {
            panic(err)
        }
        row = strings.Trim(row, "~\r\n")
        ch <- row
    }
    close(ch)
}

type Delimiters struct {
    SegmentTerm byte
    ElementTerm byte
    SubelementTerm byte
    RetitionTerm byte
}

func getDelimiters(first string) Delimiters {
    d := Delimiters{
        first[len(first)-1],
        first[3],
        first[len(first)-2],
        0,
    }
    if first[84:89] == "005010" {
        d.SubelementTerm = first[82]
    }
    return d
}
