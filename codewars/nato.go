// https://www.codewars.com/kata/if-you-can-read-this-dot-dot-dot/train/go
package codewars

import (
	"strings"
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
