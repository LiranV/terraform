// Code generated by "stringer -type=CheckStatus checks.go"; DO NOT EDIT.

package states

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[CheckPassed-9989]
	_ = x[CheckFailed-128293]
	_ = x[CheckPending-129335]
}

const (
	_CheckStatus_name_0 = "CheckPassed"
	_CheckStatus_name_1 = "CheckFailed"
	_CheckStatus_name_2 = "CheckPending"
)

func (i CheckStatus) String() string {
	switch {
	case i == 9989:
		return _CheckStatus_name_0
	case i == 128293:
		return _CheckStatus_name_1
	case i == 129335:
		return _CheckStatus_name_2
	default:
		return "CheckStatus(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}