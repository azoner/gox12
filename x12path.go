/*
Parses an x12 path

An x12 path is comprised of a sequence of loop identifiers, a segment
identifier, and element position, and a composite position.

The last loop id might be a segment id.
/LOOP_1/LOOP_2
/LOOP_1/LOOP_2/SEG
/LOOP_1/LOOP_2/SEG02
/LOOP_1/LOOP_2/SEG[424]02-1
SEG[434]02-1
02-1
02
*/

package gox12

import (
	"fmt"
	"regexp"
	"strings"
    "strconv"
)

// x12 path
type X12Path struct {
	SegmentId     string
	IdValue       string
	ElementIdx    int
	SubelementIdx int
	isRelative    bool
	Loops         []string
}

func NewX12Path(path_str string) (x12path X12Path, err error) {
	re_seg_id := "(?P<seg_id>[A-Z][A-Z0-9]{1,2})?"
	re_id_val := "(\\[(?P<id_val>[A-Z0-9]+)\\])?"
	re_ele_idx := "(?P<ele_idx>[0-9]{2})?"
	re_subele_idx := "(-(?P<subele_idx>[0-9]+))?"
	re_str := fmt.Sprintf("^%s%s%s%s$", re_seg_id, re_id_val, re_ele_idx, re_subele_idx)
	re := regexp.MustCompile(re_str)

	if path_str == "" {
		x12path.isRelative = true
		return
	}
	if path_str[0] == '/' {
		x12path.isRelative = false
		path_str = path_str[1:]
	} else {
		x12path.isRelative = true
	}
	loops := strings.Split(path_str, "/")
	//if len(loops) == 0 {
	//    err = Error
	//    return nil, err
	//}
	if len(loops) > 0 && loops[len(loops)-1] == "" {
		// Ended in a /, so no segment
		loops = loops[:len(loops)-1]
		x12path.Loops = loops
	}
	if len(loops) > 0 {
		seg_str := loops[len(loops)-1]
		match := re.FindStringSubmatch(seg_str)
		if match == nil {
			// no segment component
			return
		}
		for i, name := range re.SubexpNames() {
			// Ignore the whole regexp match and unnamed groups
			if i == 0 || name == "" {
				continue
			}
			switch name {
			case "seg_id":
				x12path.SegmentId = match[i]
			case "id_val":
				x12path.IdValue = match[i]
			case "ele_idx":
                v, _ := strconv.ParseInt(match[i], 10, 16)
				x12path.ElementIdx = int(v)
			case "subele_idx":
                v, _ := strconv.ParseInt(match[i], 10, 16)
				x12path.SubelementIdx = int(v)
			}
		}
		x12path.Loops = x12path.Loops[:len(loops)-1]
		if x12path.SegmentId == "" && x12path.IdValue != "" {
			err = fmt.Errorf("Path '%s' is invalid. Must specify a segment identifier with a qualifier", path_str)
			return
		}
		if x12path.SegmentId == "" && (x12path.ElementIdx != 0 || x12path.SubelementIdx != 0) && len(x12path.Loops) > 0 {
			err = fmt.Errorf("Path '%s' is invalid. Must specify a segment identifier", path_str)
			return
		}
	}
	return
}

// Is the path empty?
func (x12path *X12Path) Empty() bool {
	return x12path.isRelative == true && len(x12path.Loops) == 0 && x12path.SegmentId == "" && x12path.ElementIdx == 0
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

func (p *X12Path) FormatRefdes() string {
	var parts []string
	if p.SegmentId != "" {
		parts = append(parts, p.SegmentId)
		if p.IdValue != "" {
			parts = append(parts, fmt.Sprintf("[%s]", p.IdValue))
		}
	}
	if p.ElementIdx > 0 {
		parts = append(parts, fmt.Sprintf("%02i", p.ElementIdx))
		if p.SubelementIdx > 0 {
			parts = append(parts, fmt.Sprintf("-%i", p.SubelementIdx))
		}
	}
	return strings.Join(parts, "")
}

func (p *X12Path) String() string {
	var parts []string
	if p.isRelative {
		parts = append(parts, "/")
	}
	parts = append(parts, strings.Join(p.Loops, "/"))
	if p.SegmentId == "" && len(p.Loops) > 0 {
		parts = append(parts, "/")
	}
	parts = append(parts, p.FormatRefdes())
	return strings.Join(parts, "")
}
