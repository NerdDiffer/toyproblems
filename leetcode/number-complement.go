package leetcode

import (
	"bytes"
	"errors"
	"strconv"
)

func findComplement(num int) int {
	b := decimalToBinary(num)
	f, _ := flipBits(b)
	d, _ := binaryToDecimal(f)
	return d
}

func decimalToBinary(n int) string {
	i := int64(n)
	return strconv.FormatInt(i, 2)
}

func flipBits(binNum string) (string, error) {
	var buf bytes.Buffer

	for _, r := range binNum {
		if r == 49 { // is it '1'?
			buf.WriteRune(r - 1) // make it '0'
		} else if r == 48 { // is it '0'?
			buf.WriteRune(r + 1) // make it '1'
		} else {
			return "", errors.New("bad input")
		}
	}

	return buf.String(), nil
}

func binaryToDecimal(s string) (int, error) {
	if i, err := strconv.ParseInt(s, 2, 0); err != nil {
		return 0, err
	} else {
		return int(i), nil
	}
}
