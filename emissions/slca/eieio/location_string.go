// Code generated by "stringer -type=Location"; DO NOT EDIT.

package eieio

import "strconv"

const _Location_name = "DomesticImportedTotal"

var _Location_index = [...]uint8{0, 8, 16, 21}

func (i Location) String() string {
	if i < 0 || i >= Location(len(_Location_index)-1) {
		return "Location(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Location_name[_Location_index[i]:_Location_index[i+1]]
}
