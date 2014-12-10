package math

import (
	"strconv"
)

func NumDigits(num int) int {
	return len(strconv.Itoa(num))
}
