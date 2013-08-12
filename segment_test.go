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

func TestSegmentParseComp01_1(t *testing.T) {
	var segtests = []struct {
		refdes   string
		expected string
	}{
		{"TST01-1", "AA"},
		{"TST01-2", "1"},
		{"TST01-3", "5"},
		{"TST02-1", "BB"},
		//{"TST03", ""},
	}
	segmentStr := "TST&AA!1!5&BB!5"
	seg := NewSegment(segmentStr, '&', '!', '^')
	for _, tt := range segtests {
		actual, found, err := seg.GetValue(tt.refdes)
		if err != nil {
			t.Errorf("Didn't get a value for [%s]", tt.refdes)
		}
		if !found {
			t.Errorf("Didn't get a value for [%s]", tt.refdes)
		}
		if actual != tt.expected {
			t.Errorf("Didn't get expected result [%s] for path [%s], instead got [%s]", tt.expected, tt.refdes, actual)
		}
	}
}

func BenchmarkSegmentParse(b *testing.B) {
	str2 := "TST&AA!1!1&BB!5"
	for i := 0; i < b.N; i++ {
		_ = NewSegment(str2, '&', '!', '^')
	}
}
