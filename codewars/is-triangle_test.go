package codewars

import "testing"

func TestIsTriangle(t *testing.T) {
	tables := []struct {
		a        int
		b        int
		c        int
		expected bool
	}{
		{4, 2, 3, true},
	}

	for _, test := range tables {
		actual := IsTriangle(test.a, test.b, test.c)

		if actual != test.expected {
			t.Errorf("\ninput: (%d, %d, %d)\nactual: %t\nexpected: %t\n", test.a, test.b, test.c, actual, test.expected)
		}
	}
}
