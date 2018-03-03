package codewars

import "testing"

func TestToNato(t *testing.T) {
	tables := []struct {
		input    string
		expected string
	}{
		{"If you can read", "India Foxtrot Yankee Oscar Uniform Charlie Alfa November Romeo Echo Alfa Delta"},
		{"Did not see that coming", "Delta India Delta November Oscar Tango Sierra Echo Echo Tango Hotel Alfa Tango Charlie Oscar Mike India November Golf"},
		{"go for it!", "Golf Oscar Foxtrot Oscar Romeo India Tango !"},
	}

	for _, test := range tables {
		actual := ToNato(test.input)

		if actual != test.expected {
			t.Errorf("\ninput: '%s'\nactual: '%s'\nexpected: '%s'\n", test.input, actual, test.expected)
		}
	}
}
