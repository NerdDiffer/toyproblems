// https://leetcode.com/problems/single-number/
package leetcode

func SingleNumber(nums []int) int {
	counts := make(map[int]int, len(nums))

	for _, num := range nums {
		if _, ok := counts[num]; ok {
			delete(counts, num)
		} else {
			counts[num] = 1
		}
	}

	var remainingNum int

	for num := range counts {
		remainingNum = num
	}

	return remainingNum
}
