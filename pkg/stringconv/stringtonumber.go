package stringconv

import "strconv"

func StrtoI(str string) (int64, error) {
	number, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0, err
	}
	return number, nil
}
