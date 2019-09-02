package stringconv

import "strconv"

func StrtoI(str string) (int64, error) {
	number, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0, err
	}
	return number, nil
}

func StrtoF(str string) (float64, error) {
	number, err := strconv.ParseFloat(str,64)
	if err != nil {
		return 0.0, err
	}
	return number, nil
}
