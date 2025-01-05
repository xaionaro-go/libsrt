package libsrt

import "C"

func boolToInt(v bool) C.int {
	if v {
		return 1
	}
	return 0
}
