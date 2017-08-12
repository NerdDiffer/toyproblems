package leetcode

func IsValid(s string) bool {
	openers := map[string]bool{
		"(": true,
		"{": true,
		"[": true,
	}

	closers := map[string]string{
		")": "(",
		"}": "{",
		"]": "[",
	}

	var stack []string

	for _, r := range s {
		ch := string(r)

		if openers[ch] {
			stack = append(stack, ch)
		} else if _, isClosing := closers[ch]; isClosing {
			if len(stack) == 0 {
				return false
			}

			var popped string
			popped, stack = stack[len(stack)-1], stack[:len(stack)-1]
			if popped != closers[ch] {
				return false
			}
		}
		// otherwise, ignore
	}

	return len(stack) == 0
}
