package codewars

import "testing"

func TestAccum(t *testing.T) {
	tables := []struct {
		input    string
		expected string
	}{
		{"abcd", "A-Bb-Ccc-Dddd"},
		{"RqaEzty", "R-Qq-Aaa-Eeee-Zzzzz-Tttttt-Yyyyyyy"},
		{"cwAt", "C-Ww-Aaa-Tttt"},
		{"ZpglnRxqenU", "Z-Pp-Ggg-Llll-Nnnnn-Rrrrrr-Xxxxxxx-Qqqqqqqq-Eeeeeeeee-Nnnnnnnnnn-Uuuuuuuuuuu"},
	}

	for _, test := range tables {
		actual := Accum(test.input)

		if actual != test.expected {
			t.Errorf("\ninput: '%s'\nactual: '%s'\nexpected: '%s'\n", test.input, actual, test.expected)
		}
	}
}
