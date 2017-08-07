package leetcode

func MoveZeroes(nums []int) {
	numZeroes := 0

	for i := 0; i < len(nums); {
		if nums[i] != 0 {
			i++
		} else {
			nums = append(nums[:i], nums[i+1:]...) // remove the 0
			numZeroes += 1                         // track how many to add in later
		}
	}

	zeroes := make([]int, numZeroes)
	nums = append(nums, zeroes...) // combine both slices
}
