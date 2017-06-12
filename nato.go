// https://www.codewars.com/kata/if-you-can-read-this-dot-dot-dot/train/go
package main

import (
	"strings"
	"toyproblems/toyproblemhelpers"
	"unicode"
	"unicode/utf8"
)

func ToNato(words string) string {
	nato := []string{"Alfa", "Bravo", "Charlie", "Delta", "Echo", "Foxtrot", "Golf", "Hotel", "India", "Juliett", "Kilo", "Lima", "Mike", "November", "Oscar", "Papa", "Quebec", "Romeo", "Sierra", "Tango", "Uniform", "Victor", "Whiskey", "X-ray", "Yankee", "Zulu"}

	var result []string

	for _, r := range words {
		if unicode.IsSpace(r) {
			continue
		} else if unicode.IsPunct(r) {
			result = append(result, string(r))
		} else {
			ch := string(r)
			upperCh := strings.ToUpper(ch)
			unicodeVal, _ := utf8.DecodeRuneInString(upperCh)
			mapped := nato[unicodeVal-65]
			result = append(result, mapped)
		}
	}

	return strings.Join(result, " ")
}

func main() {
	test("If you can read", "India Foxtrot Yankee Oscar Uniform Charlie Alfa November Romeo Echo Alfa Delta")
	test("Did not see that coming", "Delta India Delta November Oscar Tango Sierra Echo Echo Tango Hotel Alfa Tango Charlie Oscar Mike India November Golf")
	test("go for it!", "Golf Oscar Foxtrot Oscar Romeo India Tango !")
}

// make these cheap printing/assertions less ugly looking
func test(input string, expected string) bool {
	actual := ToNato(input)
	return toyproblemhelpers.PrintCmpStrings(actual, expected)
}
