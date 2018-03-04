package codewars

import (
	"fmt"
)

var vals = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

func Decode(roman string) int {
	sum := 0
	n := len(roman)

	if n < 1 {
		return sum
	}

	var curr, next int

	for i := 0; i < n-1; i++ {
		curr = vals[string(roman[i])]
		next = vals[string(roman[i+1])]

		if curr < next {
			sum -= curr
		} else {
			sum += curr
		}
	}

	last := vals[string(roman[n-1])]
	sum += int(last)

	return sum
}

func show(s string, i int) {
	fmt.Printf("%v %T, %v %T, %v %T", i, i, s[i], s[i], string(s[i]), string(s[i]))
	fmt.Println()
}
