package leetcode

import "testing"

func TestFindComplement(t *testing.T) {
	tables := []struct {
		input    int
		expected int
	}{
		{5, 2},
		{1, 0},
	}

	for _, test := range tables {
		actual := findComplement(test.input)

		if actual != test.expected {
			t.Errorf("\nInput: %d\nActual: %d\nExpected: %d\n", test.input, actual, test.expected)
		}
	}
}

func TestDecimalToBinary(t *testing.T) {
	tables := []struct {
		input    int
		expected string
	}{
		{5, "101"},
		{4, "100"},
		{2, "10"},
	}

	for _, test := range tables {
		actual := decimalToBinary(test.input)

		if actual != test.expected {
			t.Errorf("\ninput: %d\nactual: '%s'\nexpected: '%s'", test.input, actual, test.expected)
		}
	}
}

func TestFlipBits(t *testing.T) {
	tables := []struct {
		input    string
		expected string
	}{
		{"1", "0"},
		{"0", "1"},
		{"101", "010"},
		{"010", "101"},
		{"100", "011"},
		{"011", "100"},
	}

	for _, test := range tables {
		actual, _ := flipBits(test.input)

		if actual != test.expected {
			t.Errorf("\ninput: '%s'\nactual: '%s'\nexpected: '%s'\n", test.input, actual, test.expected)
		}
	}
}

func TestBinaryToDecimal(t *testing.T) {
	tables := []struct {
		input    string
		expected int
	}{
		{"1", 1},
		{"0", 0},
		{"101", 5},
		{"010", 2},
	}

	for _, test := range tables {
		actual, _ := binaryToDecimal(test.input)

		if actual != test.expected {
			t.Errorf("\ninput: '%s'\nactual: %d\nexpected: %d\n", test.input, actual, test.expected)
		}
	}
}
