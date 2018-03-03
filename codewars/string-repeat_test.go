package codewars

import "testing"

func TestRepeatStr(t *testing.T) {
	tables := []struct {
		n        int
		s        string
		expected string
	}{
		{6, "I", "IIIIII"},
		{5, "Hello", "HelloHelloHelloHelloHello"},
	}

	for _, test := range tables {
		actual := RepeatStr(test.n, test.s)

		if actual != test.expected {
			t.Errorf("\ninput: (%d, '%s')\nactual: '%s'\nexpected: '%s'\n", test.n, test.s, actual, test.expected)
		}
	}
}
