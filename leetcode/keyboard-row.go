// https://leetcode.com/problems/keyboard-row
package leetcode

import (
	"strings"
)

func FindWords(words []string) []string {
	// map of letter to row number
	var keyToRow = map[string]int{
		"q": 1, "w": 1, "e": 1, "r": 1, "t": 1, "y": 1, "u": 1, "i": 1, "o": 1, "p": 1,
		"a": 2, "s": 2, "d": 2, "f": 2, "g": 2, "h": 2, "j": 2, "k": 2, "l": 2,
		"z": 3, "x": 3, "c": 3, "v": 3, "b": 3, "n": 3, "m": 3,
	}

	getMapKey := func(word string, ind int) string {
		return strings.ToLower(string(word[ind]))
	}

	result := []string{}

	// for each word in the list of words:
	for _, word := range words {
		is_same := true

		key := getMapKey(word, 0)
		cmp := keyToRow[key]

		// look at every letter in the word
		for j := 1; j < len(word); j++ {
			// are all of the (lower-case) letters in the same row?
			key = getMapKey(word, j)
			row := keyToRow[key]

			// if no, move to next word in list of words
			if row != cmp {
				is_same = false
				break
			}
		}

		// if yes, add to result
		if is_same {
			result = append(result, word)
		}
	}

	return result
}
