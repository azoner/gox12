package gox12

import (
	//"fmt"
	"testing"
)

func BenchmarkSegmentParseFormatIdentity(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fpath, _ := ParseX12Path("/2000A/2000B/2300/2400/SV2[421]01")
		_ = fpath.String()
	}
}

func BenchmarkRefDes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = ParseX12Path("N1[372]02-5")
	}
}

func BenchmarkRelativePath(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = ParseX12Path("1000A/1000B/TST[AA]02")
	}
}

func TestSegmentParseFormatIdentity(t *testing.T) {
	paths := [...]string{
		"/2000A/2000B/2300/2400/SV2",
		"/2000A/2000B/2300/2400/SV201",
		"/2000A/2000B/2300/2400/SV2[421]01",
	}
	for _, p := range paths {
		//fmt.Println(p)
		fpath, err := ParseX12Path(p)
		if err != nil {
			t.Errorf("Didn't get a value for [%s]", p)
		}
		//fmt.Println(fpath)
		actual := fpath.String()
		if actual != p {
			t.Errorf("Didn't get expected result [%s], instead got [%s]", p, actual)
		}
	}
}

func TestRefDes(t *testing.T) {
	var tests = []struct {
		spath     string
		seg_id    string
		qual      string
		eleidx    int
		subeleidx int
	}{
		{"TST", "TST", "", 0, 0},
		{"TST02", "TST", "", 2, 0},
		{"TST03-2", "TST", "", 3, 2},
		{"TST[AA]02", "TST", "AA", 2, 0},
		{"TST[1B5]03-1", "TST", "1B5", 3, 1},
		{"03", "", "", 3, 0},
		{"03-2", "", "", 3, 2},
		{"N102", "N1", "", 2, 0},
		{"N102-5", "N1", "", 2, 5},
		{"N1[AZR]02", "N1", "AZR", 2, 0},
		{"N1[372]02-5", "N1", "372", 2, 5},
	}
	for _, tt := range tests {
		actual, err := ParseX12Path(tt.spath)
		if err != nil {
			t.Errorf("Didn't get a value for [%s]", actual)
		}
		if actual.SegmentId != tt.seg_id {
			t.Errorf("Didn't get expected result [%s], instead got [%s]", tt.seg_id, actual.SegmentId)
		}
		if actual.IdValue != tt.qual {
			t.Errorf("Didn't get expected result [%s], instead got [%s]", tt.qual, actual.IdValue)
		}
		if actual.ElementIdx != tt.eleidx {
			t.Errorf("Didn't get expected result [%s], instead got [%s]", tt.eleidx, actual.ElementIdx)
		}
		if actual.SubelementIdx != tt.subeleidx {
			t.Errorf("Didn't get expected result [%s], instead got [%s]", tt.subeleidx, actual.SubelementIdx)
		}
		if len(actual.Path) != 0 {
			t.Errorf("Path is not empty")
		}
		path := actual.String()
		if path != tt.spath {
			t.Errorf("Didn't get expected result [%s], instead got [%s]", tt.spath, path)
		}
	}
}

func stringSliceEquals(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestRefDesMatchingNone(t *testing.T) {
	tpath := "/ISA_LOOP/GS_LOOP/ST_LOOP/DETAIL/2000A/2000B/2300/2400"
	actual, err := ParseX12Path(tpath)
	if err != nil {
		t.Errorf("Didn't get a value for [%s], error is [%s]", tpath, err.Error())
	}
	apath := actual.String()
	if apath != tpath {
		t.Errorf("Didn't get expected result [%s], instead got [%s]", tpath, apath)
	}
	if actual.SegmentId != "" {
		t.Errorf("Didn't get expected result [%s], instead got [%s]", "", actual.SegmentId)
	}
}

func TestRefDesMatchingOk(t *testing.T) {
	tpath := "/ISA_LOOP/GS_LOOP/ST_LOOP/DETAIL/2000A/2000B/2300/2400/SV2"
	actual, err := ParseX12Path(tpath)
	if err != nil {
		t.Errorf("Didn't get a value for [%s]", tpath)
	}
	apath := actual.String()
	if apath != tpath {
		t.Errorf("Didn't get expected result [%s], instead got [%s]", tpath, apath)
	}
	if actual.SegmentId != "SV2" {
		t.Errorf("Didn't get expected result [%s], instead got [%s]", "SV2", actual.SegmentId)
	}
}

func TestBadPaths(t *testing.T) {
	var tests = []struct {
		spath string
	}{
		{"AAA/03"},
		{"BB/CC/03-2"},
		{"/AAA/03"},
		{"/BB/CC/03-2"},
	}
	for _, tt := range tests {
		actual, err := ParseX12Path(tt.spath)
		if err == nil {
			t.Errorf("Path should be invalid Didn't get a value for [%s]", actual)
		}
	}
}

func TestX12PathGeneralNoSegment(t *testing.T) {
	var tests = []struct {
		spath   string
		x12Path X12Path
	}{
		{"ISA_LOOP/GS_LOOP", X12Path{Path: "ISA_LOOP/GS_LOOP"}},
		{"GS_LOOP", X12Path{Path: "GS_LOOP"}},
		{"ST_LOOP/DETAIL/2000", X12Path{Path: "ST_LOOP/DETAIL/2000"}},
		{"GS_LOOP/ST_LOOP/DETAIL/2000A", X12Path{Path: "GS_LOOP/ST_LOOP/DETAIL/2000A"}},
		{"DETAIL/2000A/2000B", X12Path{Path: "DETAIL/2000A/2000B"}},
		{"2000A/2000B/2300", X12Path{Path: "2000A/2000B/2300"}},
		{"2000B/2300/2400", X12Path{Path: "2000B/2300/2400"}},
		{"ST_LOOP/HEADER", X12Path{Path: "ST_LOOP/HEADER"}},
		{"ISA_LOOP/GS_LOOP/ST_LOOP/HEADER/1000A", X12Path{Path: "ISA_LOOP/GS_LOOP/ST_LOOP/HEADER/1000A"}},
		{"GS_LOOP/ST_LOOP/HEADER/1000B", X12Path{Path: "GS_LOOP/ST_LOOP/HEADER/1000B"}},
		{"/ISA_LOOP/GS_LOOP", X12Path{Path: "/ISA_LOOP/GS_LOOP"}},
		{"/ISA_LOOP/GS_LOOP/ST_LOOP/DETAIL/2000", X12Path{Path: "/ISA_LOOP/GS_LOOP/ST_LOOP/DETAIL/2000"}},
		{"/ISA_LOOP/GS_LOOP/ST_LOOP/DETAIL/2000A", X12Path{Path: "/ISA_LOOP/GS_LOOP/ST_LOOP/DETAIL/2000A"}},
		{"/ISA_LOOP/GS_LOOP/ST_LOOP/DETAIL/2000A/2000B", X12Path{Path: "/ISA_LOOP/GS_LOOP/ST_LOOP/DETAIL/2000A/2000B"}},
		{"/ISA_LOOP/GS_LOOP/ST_LOOP/DETAIL/2000A/2000B/2300", X12Path{Path: "/ISA_LOOP/GS_LOOP/ST_LOOP/DETAIL/2000A/2000B/2300"}},
		{"/ISA_LOOP/GS_LOOP/ST_LOOP/DETAIL/2000A/2000B/2300/2400", X12Path{Path: "/ISA_LOOP/GS_LOOP/ST_LOOP/DETAIL/2000A/2000B/2300/2400"}},
		{"/ISA_LOOP/GS_LOOP/ST_LOOP/HEADER", X12Path{Path: "/ISA_LOOP/GS_LOOP/ST_LOOP/HEADER"}},
		{"/ISA_LOOP/GS_LOOP/ST_LOOP/HEADER/1000A", X12Path{Path: "/ISA_LOOP/GS_LOOP/ST_LOOP/HEADER/1000A"}},
		{"/ISA_LOOP/GS_LOOP/ST_LOOP/HEADER/1000B", X12Path{Path: "/ISA_LOOP/GS_LOOP/ST_LOOP/HEADER/1000B"}},
	}
	for _, tt := range tests {
		actual, err := ParseX12Path(tt.spath)
		if err != nil {
			t.Errorf("Didn't get a value for [%s]", tt.spath)
		}
		if actual.IsAbs() != tt.x12Path.IsAbs() {
			t.Errorf("[%s] was not relative", tt.spath)
		}
		if actual.SegmentId != "" {
			t.Errorf("Didn't get expected result [%s], instead got [%s]", "", actual.SegmentId)
		}
		if actual.IdValue != "" {
			t.Errorf("Didn't get expected result [%s], instead got [%s]", "", actual.IdValue)
		}
		if actual.ElementIdx != 0 {
			t.Errorf("Didn't get expected result [%s], instead got [%s]", 0, actual.ElementIdx)
		}
		if actual.SubelementIdx != 0 {
			t.Errorf("Didn't get expected result [%s], instead got [%s]", 0, actual.SubelementIdx)
		}
		if actual.Path != tt.x12Path.Path {
			t.Errorf("Path: Didn't get expected result [%s], instead got [%s]", tt.x12Path.Path, actual.Path)
		}
		path := actual.String()
		if path != tt.x12Path.String() {
			t.Errorf("String Didn't get expected result [%s], instead got [%s]", tt.x12Path.String(), path)
		}
	}
}

func TestX12PathGeneral(t *testing.T) {
	var tests = []struct {
		spath   string
		x12Path X12Path
	}{
		{"/AAA/TST", X12Path{"/AAA", "TST", "", 0, 0}},
		{"/B1000/TST02", X12Path{"/B1000", "TST", "", 2, 0}},
		{"/1000B/TST03-2", X12Path{"/1000B", "TST", "", 3, 2}},
		{"/1000A/1000B/TST[AA]02", X12Path{"/1000A/1000B", "TST", "AA", 2, 0}},
		{"/AA/BB/CC/TST[1B5]03-1", X12Path{"/AA/BB/CC", "TST", "1B5", 3, 1}},
		{"/DDD/E1000/N102", X12Path{"/DDD/E1000", "N1", "", 2, 0}},
		{"/E1000/D322/N102-5", X12Path{"/E1000/D322", "N1", "", 2, 5}},
		{"/BB/CC/N1[AZR]02", X12Path{"/BB/CC", "N1", "AZR", 2, 0}},
		{"/BB/CC/N1[372]02-5", X12Path{"/BB/CC", "N1", "372", 2, 5}},
		{"AAA/TST", X12Path{"AAA", "TST", "", 0, 0}},
		{"B1000/TST02", X12Path{"B1000", "TST", "", 2, 0}},
		{"1000B/TST03-2", X12Path{"1000B", "TST", "", 3, 2}},
		{"1000A/1000B/TST[AA]02", X12Path{"1000A/1000B", "TST", "AA", 2, 0}},
		{"AA/BB/CC/TST[1B5]03-1", X12Path{"AA/BB/CC", "TST", "1B5", 3, 1}},
		{"DDD/E1000/N102", X12Path{"DDD/E1000", "N1", "", 2, 0}},
		{"E1000/D322/N102-5", X12Path{"E1000/D322", "N1", "", 2, 5}},
		{"BB/CC/N1[AZR]02", X12Path{"BB/CC", "N1", "AZR", 2, 0}},
		{"BB/CC/N1[372]02-5", X12Path{"BB/CC", "N1", "372", 2, 5}},
	}
	for _, tt := range tests {
		actual, err := ParseX12Path(tt.spath)
		if err != nil {
			t.Errorf("Didn't get a value for [%s]", tt.spath)
		}
		if actual.IsAbs() != tt.x12Path.IsAbs() {
			t.Errorf("[%s] was not relative", tt.spath)
		}
		if actual.SegmentId != tt.x12Path.SegmentId {
			t.Errorf("Didn't get expected result [%s], instead got [%s]", tt.x12Path.SegmentId, actual.SegmentId)
		}
		if actual.IdValue != tt.x12Path.IdValue {
			t.Errorf("Didn't get expected result [%s], instead got [%s]", tt.x12Path.IdValue, actual.IdValue)
		}
		if actual.ElementIdx != tt.x12Path.ElementIdx {
			t.Errorf("Didn't get expected result [%s], instead got [%s]", tt.x12Path.ElementIdx, actual.ElementIdx)
		}
		if actual.SubelementIdx != tt.x12Path.SubelementIdx {
			t.Errorf("Didn't get expected result [%s], instead got [%s]", tt.x12Path.SubelementIdx, actual.SubelementIdx)
		}
		if actual.Path != tt.x12Path.Path {
			t.Errorf("Path: Didn't get expected result [%s], instead got [%s]", tt.x12Path.Path, actual.Path)
		}
		path := actual.String()
		if path != tt.x12Path.String() {
			t.Errorf("String Didn't get expected result [%s], instead got [%s]", tt.x12Path.String(), path)
		}
	}
}

/*

class AbsolutePath(unittest.TestCase):
    def test_plain_loops(self):
        paths = [
        ]
        for spath in paths:
            plist = spath.split("/")[1:]
            rd = pyx12.path.X12Path(spath)
            self.assertEqual(rd.loop_list, plist,
                             "%s: %s != %s" % (spath, rd.loop_list, plist))


class Equality(unittest.TestCase):
    def test_equal1(self):
        p1 = pyx12.path.X12Path("/ISA_LOOP/GS_LOOP/ST_LOOP/DETAIL/2000A")
        p2 = pyx12.path.X12Path("/ISA_LOOP/GS_LOOP/ST_LOOP/DETAIL/2000A")
        self.assertEqual(p1, p2)
        self.assertEqual(p1.format(), p2.format())
        self.assertEqual(p1.__hash__(), p2.__hash__())

    def test_equal2(self):
        p1 = pyx12.path.X12Path("/ISA_LOOP/GS_LOOP/ST_LOOP/DETAIL/2000A")
        p2 = pyx12.path.X12Path("/ISA_LOOP/GS_LOOP/ST_LOOP/DETAIL/")
        p2.loop_list.append("2000A")
        self.assertEqual(p1, p2)
        self.assertEqual(p1.format(), p2.format())
        self.assertEqual(p1.__hash__(), p2.__hash__())

    def test_equal3(self):
        p1 = pyx12.path.X12Path("/AA/BB/CC/TST[1B5]03-1")
        p2 = pyx12.path.X12Path("/AA/BB/CC/AAA[1B5]03-1")
        p2.seg_id = "TST"
        self.assertEqual(p1, p2)
        self.assertEqual(p1.format(), p2.format())
        self.assertEqual(p1.__hash__(), p2.__hash__())

    def test_equal4(self):
        p1 = pyx12.path.X12Path("1000B/TST03-2")
        p2 = pyx12.path.X12Path("1000B/TST04-2")
        p2.ele_idx = 3
        self.assertEqual(p1, p2)
        self.assertEqual(p1.format(), p2.format())
        self.assertEqual(p1.__hash__(), p2.__hash__())


class NonEquality(unittest.TestCase):
    def test_nequal1(self):
        p1 = pyx12.path.X12Path("/ISA_LOOP/GS_LOOP/ST_LOOP/DETAIL/2000A")
        p2 = pyx12.path.X12Path("ISA_LOOP/GS_LOOP/ST_LOOP/DETAIL/2000A")
        self.assertNotEqual(p1, p2)
        self.assertNotEqual(p1.format(), p2.format())
        self.assertNotEqual(p1.__hash__(), p2.__hash__())

    def test_nequal2(self):
        p1 = pyx12.path.X12Path("/ISA_LOOP/GS_LOOP/ST_LOOP/DETAIL/2000A")
        p2 = pyx12.path.X12Path("/ISA_LOOP/GS_LOOP/ST_LOOP/DETAIL/")
        self.assertNotEqual(p1, p2)
        self.assertNotEqual(p1.format(), p2.format())
        self.assertNotEqual(p1.__hash__(), p2.__hash__())

    def test_nequal3(self):
        p1 = pyx12.path.X12Path("/AA/BB/CC/TST[1B5]03-1")
        p2 = pyx12.path.X12Path("/AA/BB/CC/AAA[1B5]03-1")
        self.assertNotEqual(p1, p2)
        self.assertNotEqual(p1.format(), p2.format())
        self.assertNotEqual(p1.__hash__(), p2.__hash__())

    def test_nequal4(self):
        p1 = pyx12.path.X12Path("1000B/TST03-2")
        p2 = pyx12.path.X12Path("1000B/TST04-2")
        self.assertNotEqual(p1, p2)
        self.assertNotEqual(p1.format(), p2.format())
        self.assertNotEqual(p1.__hash__(), p2.__hash__())


class Empty(unittest.TestCase):
    def test_not_empty_1(self):
        p1 = "1000B/TST03-2"
        self.assertFalse(pyx12.path.X12Path(
            p1).empty(), "Path "%s" is not empty" % (p1))

    def test_not_empty_2(self):
        p1 = "/AA/BB/CC/AAA[1B5]03"
        self.assertFalse(pyx12.path.X12Path(
            p1).empty(), "Path "%s" is not empty" % (p1))

    def test_not_empty_3(self):
        p1 = "GS_LOOP/ST_LOOP/DETAIL/2000A"
        self.assertFalse(pyx12.path.X12Path(
            p1).empty(), "Path "%s" is not empty" % (p1))

    def test_not_empty_4(self):
        p1 = "/ISA_LOOP/GS_LOOP/ST_LOOP/DETAIL/2000A"
        self.assertFalse(pyx12.path.X12Path(
            p1).empty(), "Path "%s" is not empty" % (p1))

    def test_not_empty_5(self):
        p1 = "/"
        self.assertFalse(pyx12.path.X12Path(
            p1).empty(), "Path "%s" is not empty" % (p1))

    def test_not_empty_6(self):
        p1 = "/ISA_LOOP/GS_LOOP/ST_LOOP/DETAIL/2000A/"
        self.assertFalse(pyx12.path.X12Path(
            p1).empty(), "Path "%s" is not empty" % (p1))

    def test_empty_1(self):
        p1 = ""
        a = pyx12.path.X12Path(p1)
        self.assertTrue(pyx12.path.X12Path(
            p1).empty(), "Path "%s" is empty" % (p1))
*/
