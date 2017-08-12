package leetcode

import "testing"

func TestCanPlaceFLowers(t *testing.T) {
	tables := []struct {
		flowerbed     []int
		numNewFlowers int
		expected      bool
	}{
		{[]int{1, 0, 0, 0, 1}, 1, true},
		{[]int{1, 0, 0, 0, 1}, 2, false},
		{[]int{1, 0, 0, 0, 0, 1}, 2, false},
		{[]int{0, 0, 0, 1}, 1, true},
		{[]int{1, 0, 0, 0, 1, 0, 0}, 2, true},
	}

	for _, test := range tables {
		actual := CanPlaceFlowers(test.flowerbed, test.numNewFlowers)

		if actual != test.expected {
			t.Errorf("\ninput: %v, %d\nactual: %t\nexpected: %t", test.flowerbed, test.numNewFlowers, actual, test.expected)
		}
	}
}
