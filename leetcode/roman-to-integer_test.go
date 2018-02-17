package leetcode

import (
	"testing"
)

func TestRomanToInts(t *testing.T) {
	tables := []struct {
		description string
		input       string
		expected    int
	}{
		{description: "no letters", input: "", expected: 0},
		{description: "one letter", input: "V", expected: 5},
		{description: "simple case", input: "LX", expected: 60},
		{description: "simple case", input: "XVI", expected: 16},
		{description: "subtraction", input: "IV", expected: 4},
		{description: "subtraction", input: "IX", expected: 9},
		{description: "subtraction", input: "XL", expected: 40},
		{description: "subtraction", input: "XC", expected: 90},
		{description: "subtraction", input: "CD", expected: 400},
		{description: "subtraction", input: "CM", expected: 900},
		{description: "big numbers", input: "DCXLVII", expected: 647},
		{description: "big numbers", input: "MMDXLIX", expected: 2549},
		{description: "big numbers", input: "MCMXLIV", expected: 1944},
		{description: "big numbers", input: "MCMXCIX", expected: 1999},
	}

	for _, test := range tables {
		actual := RomanToInt(test.input)

		if actual != test.expected {
			t.Errorf("\ndescription: %s\ninput: %v\nactual: %v\nexpected: %v", test.description, test.input, actual, test.expected)
		}
	}
}
