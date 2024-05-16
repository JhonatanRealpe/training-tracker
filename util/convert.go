package util

import "strconv"

func StrinToInt(s string) (int, error) {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}
	return i, nil
}
func IntToStrin(i int) string {
	return strconv.Itoa(i)
}
