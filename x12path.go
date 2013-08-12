/*
Parses an x12 path

*/

package gox12

import (
	"bytes"
	"errors"
	"fmt"
	"path"
	"regexp"
	"strconv"
	"strings"
)

var refdesRegexp = regexp.MustCompile("^(?P<seg_id>[A-Z][A-Z0-9]{1,2})?(\\[(?P<id_val>[A-Z0-9]+)\\])?(?P<ele_idx>[0-9]{2})?(-(?P<subele_idx>[0-9]+))?$")

// An X12 path is comprised of a path of loop identifiers, a segment
// identifier, and element position, and a composite position.
//
// The last loop id might be a segment id.
// /LOOP_1/LOOP_2
// /LOOP_1/LOOP_2/SEG
// /LOOP_1/LOOP_2/SEG02
// /LOOP_1/LOOP_2/SEG[424]02-1
// LOOP_2/SEG02
// SEG[434]02-1
// 02-1
// 02
type X12Path struct {
	Path          string // no leading slash indicates a relative path
	SegmentId     string
	IdValue       string
	ElementIdx    int
	SubelementIdx int
}

// Maybe s is of the form t c u.
// If so, return t, c u (or t, u if cutc == true).
// If not, return s, "".
func split(s string, c string, cutc bool) (string, string) {
	i := strings.Index(s, c)
	if i < 0 {
		return s, ""
	}
	if cutc {
		return s[0:i], s[i+len(c):]
	}
	return s[0:i], s[i:]
}

// Parse parses an X12 Path string into component parts,
// The last part of the may may be a segment identifier
func ParseX12Path(rawpath string) (x12path *X12Path, err error) {
	if rawpath == "" {
		err = errors.New("empty x12path")
		return nil, err
	}
	// get last part of path
	// try to parse as a segment, rest is path
	// if that fails, it is all the path with no segment
	x12path = new(X12Path)
	// set struct values...
	basepath, refdes := path.Split(rawpath)
	var seg_id string
	var id_val string
	var ele_idx int
	var subele_idx int
	if seg_id, id_val, ele_idx, subele_idx, err = parseRefDes(refdes); err != nil {
		// not a segment
		x12path.Path = rawpath
		return nil, err
	}
	if basepath != "" && basepath[len(basepath)-1] == '/' {
		x12path.Path = basepath[:len(basepath)-1]
	}
	x12path.SegmentId = seg_id
	x12path.IdValue = id_val
	x12path.ElementIdx = ele_idx
	x12path.SubelementIdx = subele_idx

	/*
		if x12path.SegmentId == "" && x12path.IdValue != "" {
			err = fmt.Errorf("Path '%s' is invalid. Must specify a segment identifier with a qualifier", path_str)
			return
		}
		if x12path.SegmentId == "" && (x12path.ElementIdx != 0 || x12path.SubelementIdx != 0) && len(x12path.Loops) > 0 {
			err = fmt.Errorf("Path '%s' is invalid. Must specify a segment identifier", path_str)
			return
		} */
	return x12path, nil
}

func (x12path *X12Path) IsAbs() bool {
	return path.IsAbs(x12path.Path)
}

func getSubeleIdx(refdes string) (rest string, idx int, err error) {
	var part string
	var num uint64
	i := strings.LastIndex(refdes, "-")
	if i != -1 {
		rest, part = refdes[0:i], refdes[i+1:]
		if num, err = strconv.ParseUint(part, 10, 8); err != nil {
			return "", 0, errors.New("Subelement index must be in range [1,99]")
		}
		idx = int(num)
		if 1 < idx || idx > 99 {
			return "", 0, errors.New("Subelement index must be in range [1,99]")
		}
		return rest, idx, nil
	}
	return rest, 0, nil
}

//func getEleIdx(refdes string) (rest string, idx int, err error) {
//    part string
//    rest, part = refdes[0:i], refdes[i+1:]
//        if idx, err = strconv.ParseUint(part, 10, 8); err != nil {
//            return nil, nil, Error("Subelement index must be in range [1,99]")
//        }
//        x := int(idx)
//        if 1 < x || x > 99 {
//            return nil, nil, Error("Subelement index must be in range [1,99]")
//        }
//        return rest, x, nil
//    return rest, nil, nil
//}

/*
func parseRefDes(refdes string) (seg_id, id_val string, ele_idx, subele_idx int, err error) {
    //  failure 1 - idx not int, depend not satisfied
    // failure 2 - is not a refdes
    var rest, myint string
	if refdes == "" {
		err = errors.New("empty refdes")
		return
	}
    if rest, subele_idx, err = getSubeleIdx(refdes); err != nil {
        goto Error2
    }
*/

func parseRefDes(refdes string) (seg_id, id_val string, ele_idx, subele_idx int, err error) {
	//  failure 1 - idx not int, depend not satisfied
	// failure 2 - is not a refdes
	if refdes == "" {
		err = errors.New("empty refdes")
		return
	}
	match := refdesRegexp.FindStringSubmatch(refdes)
	if match == nil {
		// no segment component
		err = errors.New("Not a refdes")
		return
	}
	for i, name := range refdesRegexp.SubexpNames() {
		// Ignore the whole regexp match and unnamed groups
		if i == 0 || name == "" {
			continue
		}
		switch name {
		case "seg_id":
			seg_id = match[i]
		case "id_val":
			id_val = match[i]
		case "ele_idx":
			v, _ := strconv.ParseInt(match[i], 10, 8)
			ele_idx = int(v)
		case "subele_idx":
			v, _ := strconv.ParseInt(match[i], 10, 8)
			subele_idx = int(v)
		}
	}
	return seg_id, id_val, ele_idx, subele_idx, nil
}

// Is the path empty?
func (x12path *X12Path) Empty() bool {
	return x12path.Path == "" && x12path.SegmentId == "" && x12path.ElementIdx == 0
}

/*
   def _is_child_path(self, root_path, child_path):
       """
       Is the child path really a child of the root path?
       @type root_path: string
       @type child_path: string
       @return: True if a child
       @rtype: boolean
       """
       root = root_path.split('/')
       child = child_path.split('/')
       if len(root) >= len(child):
           return False
       for i in range(len(root)):
           if root[i] != child[i]:
               return False
       return True
*/

// Assemble the segment parts of a X12Path into a string
func (p *X12Path) FormatRefdes() string {
	var parts []string
	if p.SegmentId != "" {
		parts = append(parts, p.SegmentId)
		if p.IdValue != "" {
			parts = append(parts, fmt.Sprintf("[%s]", p.IdValue))
		}
	}
	if p.ElementIdx > 0 {
		parts = append(parts, fmt.Sprintf("%02d", p.ElementIdx))
		if p.SubelementIdx > 0 {
			parts = append(parts, fmt.Sprintf("-%d", p.SubelementIdx))
		}
	}
	return strings.Join(parts, "")
}

// String reassembles the X12Path into a valid X12Path string
// See pkg/net/url/url.go:String
func (p *X12Path) String() string {
	var buf bytes.Buffer
	if p.Path != "" {
		buf.WriteString(p.Path)
	}
	rd := p.FormatRefdes()
	if p.Path != "" && rd != "" {
		buf.WriteByte('/')
	}
	if rd != "" {
		buf.WriteString(rd)
	}
	return buf.String()
}
