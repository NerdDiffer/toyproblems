package codewars

import "testing"

func TestFizzBuzzCuckooClock(t *testing.T) {
	tables := []struct {
		input    string
		expected string
	}{
		{"13:34", "tick"},
		{"21:00", "Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo"},
		{"11:15", "Fizz Buzz"},
		{"03:03", "Fizz"},
		{"14:30", "Cuckoo"},
		{"08:55", "Buzz"},
		{"00:00", "Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo"},
		{"12:00", "Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo"},
	}

	for _, test := range tables {
		actual := FizzBuzzCuckooClock(test.input)

		if actual != test.expected {
			t.Errorf("\ninput: '%s'\nactual: '%s'\nexpected: '%s'\n", test.input, actual, test.expected)
		}
	}
}
