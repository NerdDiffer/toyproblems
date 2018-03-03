package codewars

import "testing"

func TestHasUniqueChar(t *testing.T) {
	tables := []struct {
		input    string
		expected bool
	}{
		{"  nAa", false},
	}

	for _, test := range tables {
		actual := HasUniqueChar(test.input)

		if actual != test.expected {
			t.Errorf("\ninput: '%s'\nactual: %t\nexpected: %t\n", test.input, actual, test.expected)
		}
	}
}
