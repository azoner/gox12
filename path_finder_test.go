package gox12

import (
	"testing"
)

// AAA
func TestMatcherISA(t *testing.T) {
	finder := NewHeaderMapFinder()
	//finder := MakeMapFinder()
	str1 := "ISA&00&          &00&       "
	seg := NewSegment(str1, '&', '!', '^')
	foundPath, ok, _ := finder.FindNext("", seg)
	if !ok {
		t.Errorf("Should have found path for [%s], ok was false", foundPath)
	}
	expectedPath := "/ISA_LOOP/ISA"
	if expectedPath != foundPath {
		t.Errorf("Didn't get expected result [%s], instead got [%s]", expectedPath, foundPath)
	}
}

func TestMatcherST(t *testing.T) {
	finder := NewHeaderMapFinder()
	str1 := "ST*001*AFD"
	seg := NewSegment(str1, '*', '!', '^')
	foundPath, ok, _ := finder.FindNext("", seg)
	if !ok {
		t.Errorf("Should have found path for [%s], ok was false", foundPath)
	}
	expectedPath := "/ISA_LOOP/GS_LOOP/ST_LOOP/ST"
	if expectedPath != foundPath {
		t.Errorf("Didn't get expected result [%s], instead got [%s]", expectedPath, foundPath)
	}
}

func TestMatcherGE(t *testing.T) {
	finder := NewHeaderMapFinder()
	str1 := "GE*001*0002"
	seg := NewSegment(str1, '*', '!', '^')
	foundPath, ok, _ := finder.FindNext("", seg)
	if !ok {
		t.Errorf("Should have found path for [%s], ok was false", foundPath)
	}
	expectedPath := "/ISA_LOOP/GS_LOOP/GE"
	if expectedPath != foundPath {
		t.Errorf("Didn't get expected result [%s], instead got [%s]", expectedPath, foundPath)
	}
}

func TestMatcherMissing(t *testing.T) {
	finder := NewHeaderMapFinder()
	str1 := "AAA*001*002"
	seg := NewSegment(str1, '*', '!', '^')
	foundPath, ok, _ := finder.FindNext("", seg)
	if ok {
		t.Errorf("Should not have found path for [%s], ok was true", foundPath)
	}
	expectedPath := ""
	if expectedPath != foundPath {
		t.Errorf("Didn't get expected result [%s], instead got [%s]", expectedPath, foundPath)
	}
}

func TestMatcherAgg(t *testing.T) {
		var tests = []struct {
			isFound bool
		spath     string
		seg_id    string
	}{
		{true, "/ISA_LOOP/ISA", "ISA"},
		{true, "/ISA_LOOP/GS_LOOP/ST_LOOP/ST", "ST"},
		{false, "", "AAA"},
		}

	finder := NewFirstMatchPathFinder(NewHeaderMapFinder())
	curPath := ""
	for _, tt := range tests {
		seg := Segment{SegmentId: tt.seg_id}
		res, ok, err := finder.FindNext(curPath, seg)
		if err != nil {
			t.Errorf("Error for [%s]", tt.seg_id)
		}
		if ok != tt.isFound {
			t.Errorf("Didn't get expected result [%s], instead got [%s]", tt.isFound, ok)
		}
		if res != tt.spath {
			t.Errorf("Didn't get expected result [%s], instead got [%s]", tt.spath, res)
		}
		if ok {
			curPath = res

		} else {
			curPath = ""
		}
	}
}