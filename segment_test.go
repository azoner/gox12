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
		{"TST01-1", ""},
	}
	segmentStr := "TST&AA!1!1&BB!5"
	seg := NewSegment(segmentStr, '&', '!', '^')
	for _, tt := range segtests {
		actual, err := seg.GetValue(tt.refdes)
		if err != nil {
			t.Errorf("Didn't get a value for [%s]", tt.refdes)
		}
		if actual != tt.expected {
			t.Errorf("Didn't get expected result [%s], instead got [%s]", tt.expected, actual)
		}
	}
}

func BenchmarkSegmentParse(b *testing.B) {
	str2 := "TST&AA!1!1&BB!5"
	for i := 0; i < b.N; i++ {
		_ = NewSegment(str2, '&', '!', '^')
	}
}
