package leetcode

import "fmt"

func judgeCircle(moves string) bool {
	vertical := 0
	horizontal := 0

	// each letter in moves
	for _, r := range moves {
		switch move := fmt.Sprintf("%c", r); move {
		case "U":
			vertical += 1
		case "D":
			vertical -= 1
		case "L":
			horizontal -= 1
		case "R":
			horizontal += 1
		default:
			// no op
		}
	}

	return vertical == 0 && horizontal == 0
}
