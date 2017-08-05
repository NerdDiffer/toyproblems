// https://leetcode.com/problems/set-mismatch/
package leetcode

import (
	"sort"
)

func FindErrorNums(nums []int) []int {
	output := make([]int, 2)
	length := len(nums)
	counts := make(map[int]int, length)

	// collect number of times each value appears in input
	for _, num := range nums {
		counts[num] += 1
	}

	for num := 1; num <= length; num++ { // consider all valid numbers
		if counts[num] == 2 { // is it a duplicate?
			output[0] = num
		} else if counts[num] == 0 { // is it missing from input?
			output[1] = num
		}
	}

	return output
}

// 1, 2, 2, 4
// 2, 2
// 2, 3, 3, 4, 5, 6
func _FindErrorNums(nums []int) []int {
	var output []int
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	len := len(nums)

	var missing int
	curr := nums[0]

	if curr != 1 && curr > 0 {
		missing = 1
	}

	for ind := 1; ind < len; ind++ {
		curr = nums[ind]

		if missing == 0 && curr != ind-1 {
			missing = curr
		}

		if missing > 0 && curr == nums[ind-1] {
			output = append(output, curr, missing)
		}

		//if curr == ind {
		//	output = append(output, curr, curr+1)
		//	break
		//} else if ind > 0 && curr == ind+1 && nums[ind-1] == curr {
		//	output = append(output, curr, curr-1)
		//	break
		//}
	}

	return output
}
