package util

import (
	"strconv"
)

func StrToIntPtr(s string) *int {
	i, err := strconv.Atoi(s)
	if err != nil {
		return nil
	}
	return &i
}
