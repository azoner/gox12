package gox12

import (
	"testing"
)

// AAA
func TestMatcherISA(t *testing.T) {
	finder := MakeMapFinder()
	str1 := "ISA&00&          &00&       "
	//str2 := "TST&AA!1!1&BB!5"
	seg := NewSegment(str1, '&', '!', '^')
	//expectedSegId := ""
	path1, _, _ := finder("", seg)
	expectedPath := "/ISA_LOOP/ISA"
	if expectedPath != path1 {
		t.Errorf("Didn't get expected result [%s], instead got [%s]", expectedPath, path1)
	}
}
