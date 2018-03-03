package codewars

import "testing"

func TestOperArray(t *testing.T) {
	a := []int{18, 69, -90, -78, 65, 40}

	tables := []struct {
		op       FParam // the operation to perform
		s        []int  // data to operate on
		i        int    // initial value
		expected []int
	}{
		{Gcdi, a, a[0], []int{18, 3, 3, 3, 1, 1}},
		{Lcmu, a, a[0], []int{18, 414, 2070, 26910, 26910, 107640}},
		{Som, a, 0, []int{18, 87, -3, -81, -16, 24}},
		{Mini, a, a[0], []int{18, 18, -90, -90, -90, -90}},
		{Maxi, a, a[0], []int{18, 69, 69, 69, 69, 69}},
	}

	for _, test := range tables {
		result := OperArray(test.op, test.s, test.i)

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
