package utils

import (
	"strconv"
)

func AddCommas(n int) string {
	str := strconv.Itoa(n)
	length := len(str)

	if length < 4 {
		return str
	}

	var result string

	for i := 0; i < length; i++ {
		if i != 0 && (length-i)%3 == 0 {
			result += "."
		}
		result += string(str[i])
	}

	return result
}
