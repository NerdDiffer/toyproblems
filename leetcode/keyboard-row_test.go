package leetcode

import (
	tph "github.com/NerdDiffer/toyproblems/helpers"
	"testing"
)

func TestFindWords(t *testing.T) {
	tables := []struct {
		input    []string
		expected []string
	}{
		{
			input:    []string{"Hello", "Alaska", "Dad", "Peace", "TypeWriter"},
			expected: []string{"Alaska", "Dad", "TypeWriter"},
		},
	}

	for _, test := range tables {
		actual := FindWords(test.input)

		if !tph.CmpStringSlices(actual, test.expected) {
			t.Errorf("actual: %s, expected: %s", actual, test.expected)
		}
	}
}
