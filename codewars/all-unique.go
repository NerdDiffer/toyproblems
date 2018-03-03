package codewars

func HasUniqueChar(str string) bool {
	var counts = map[rune]bool{}

	for _, val := range str {
		if _, isPresent := counts[val]; isPresent {
			return false
		} else {
			counts[val] = true
		}
	}

	return true
}
