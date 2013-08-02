package gox12

import (
	"testing"
)

// AAA
func TestSegmentParseSegmentId(t *testing.T) {
    str2 := "TST&AA!1!1&BB!5"
    seg := NewSegment(str2, '&', '!', '^')
    expectedSegId := "TST"
	if seg.SegmentId != expectedSegId {
		t.Errorf("Didn't get expected result [%s], instead got [%s]", expectedSegId, seg.SegmentId)
	}
}
