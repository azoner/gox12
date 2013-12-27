package gox12

import (
	"testing"
)

// AAA
func TestMatcherISA(t *testing.T) {
	finder := MakeMapFinder()
	str1 := "ISA&00&          &00&       "
	seg := NewSegment(str1, '&', '!', '^')
	foundPath, _, _ := finder("", seg)
	expectedPath := "/ISA_LOOP/ISA"
	if expectedPath != foundPath {
		t.Errorf("Didn't get expected result [%s], instead got [%s]", expectedPath, foundPath)
	}
}
