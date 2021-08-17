// Code generated by "stringer -type=CardSuit"; DO NOT EDIT.

package deck

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[S_Joker-0]
	_ = x[S_Spades-1]
	_ = x[S_Diamonds-2]
	_ = x[S_Clubs-3]
	_ = x[S_Hearts-4]
}

const _CardSuit_name = "S_JokerS_SpadesS_DiamondsS_ClubsS_Hearts"

var _CardSuit_index = [...]uint8{0, 7, 15, 25, 32, 40}

func (i CardSuit) String() string {
	if i < 0 || i >= CardSuit(len(_CardSuit_index)-1) {
		return "CardSuit(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _CardSuit_name[_CardSuit_index[i]:_CardSuit_index[i+1]]
}