package leetcode

import "testing"

func TestJudgeRouteCircle(t *testing.T) {
	tables := []struct {
		input    string
		expected bool
	}{
		{"UD", true},
		{"LL", false},
		{"", true},
	}

	for _, test := range tables {
		actual := judgeCircle(test.input)

		if actual != test.expected {
			t.Errorf("\ninput: %s\nactual: %t\nexpected: %t", test.input, actual, test.expected)
		}
	}
}
