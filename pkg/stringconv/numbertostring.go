package stringconv

import "strconv"

func FtoStr(number float64) string {
	s := strconv.FormatFloat(number, 'E', -1, 64)
	return s
}
