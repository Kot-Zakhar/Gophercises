// Code generated by "stringer --type CardValue"; DO NOT EDIT.

package deck

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[VJoker-0]
	_ = x[VA-1]
	_ = x[V1-2]
	_ = x[V2-3]
	_ = x[V3-4]
	_ = x[V4-5]
	_ = x[V5-6]
	_ = x[V6-7]
	_ = x[V7-8]
	_ = x[V8-9]
	_ = x[V9-10]
	_ = x[V10-11]
	_ = x[VJ-12]
	_ = x[VQ-13]
	_ = x[VK-14]
}

const _CardValue_name = "VJokerVAV1V2V3V4V5V6V7V8V9V10VJVQVK"

var _CardValue_index = [...]uint8{0, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 29, 31, 33, 35}

func (i CardValue) String() string {
	if i < 0 || i >= CardValue(len(_CardValue_index)-1) {
		return "CardValue(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _CardValue_name[_CardValue_index[i]:_CardValue_index[i+1]]
}
