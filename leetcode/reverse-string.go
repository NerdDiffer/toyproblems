package main

import (
	"fmt"
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

func main() {
	fmt.Println(reverseString("racecar"))
	fmt.Println(reverseString("foobar"))
}
