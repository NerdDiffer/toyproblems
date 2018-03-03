package codewars

import "testing"

func TestSquareOrSquareRoot(t *testing.T) {
	tables := []struct {
		input    []int
		expected []int
	}{
		{[]int{4, 3, 9, 7, 2, 1}, []int{2, 9, 3, 49, 4, 1}},
	}

	for _, test := range tables {
		result := SquareOrSquareRoot(test.input)

		if len(result) != len(test.expected) {
			t.Errorf("actual result must have same length as expected result")
			panic("stop")
		}

		for i, num := range result {
			if num != test.expected[i] {
				t.Errorf("\nindex: %d\nactual: %d\nexpected: %d\n", i, num, test.expected[i])
			}
		}
	}
}
