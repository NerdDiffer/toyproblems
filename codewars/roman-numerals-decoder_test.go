package codewars

import "testing"

func TestDecoder(t *testing.T) {
	tables := []struct {
		input    string
		expected int
	}{
		{"XXI", 21},
		{"I", 1},
		{"VI", 6},
		{"IV", 4},
		{"MMVIII", 2008},
		{"MDCLXVI", 1666},
	}

	for _, test := range tables {
		actual := Decode(test.input)

		if actual != test.expected {
			t.Errorf("\ninput: '%s'\nactual: %d\nexpected: %d\n", test.input, actual, test.expected)
		}
	}
}
