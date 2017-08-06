package leetcode

import (
	"strings"
)

func reverseString(s string) string {
	slice := strings.Split(s, "")
	length := len(slice)
	mid := length/2 - 1

	for i := mid; i >= 0; i-- {
		j := length - 1 - i
		slice[i], slice[j] = slice[j], slice[i]
	}

	return strings.Join(slice, "")
}

// Given a string and an integer k, you need to reverse the first k characters for
// every 2k characters counting from the start of the string. If there are less
// than k characters left, reverse all of them. If there are less than 2k but
// greater than or equal to k characters, then reverse the first k characters and
// left the other as original.
func ReverseStr(s string, k int) string {
	if k < 2 {
		return s
	}

	length := len(s)
	result := make([]string, 0, length)
	shouldReverse := true

	for i := 0; i < length; {
		numRemaining := length - i
		var strSlice string

		if numRemaining < k {
			strSlice = s[i:]
		} else {
			strSlice = s[i : i+k]
		}

		toAppend := strSlice

		if shouldReverse {
			toAppend = reverseString(strSlice)
		}

		result = append(result, toAppend)
		shouldReverse = !shouldReverse
		i += k
	}

	return strings.Join(result, "")
}
