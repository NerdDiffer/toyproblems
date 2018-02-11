package leetcode

func SearchInsertPosition(list []int, val int) int {
	length := len(list)
	curr := list[0]

	if length < 1 || val <= curr {
		return 0
	}

	var prev int

	for i := 1; i < length; i += 1 {
		prev = list[i-1]
		curr = list[i]

		if val > prev && val <= curr {
			return i
		}
	}

	return length
}
