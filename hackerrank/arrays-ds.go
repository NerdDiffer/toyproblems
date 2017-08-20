package main

import "fmt"

func main() {
	var length int
	_, _ = fmt.Scanf("%d", &length)

	nums := populate(length)
	reverse(nums)
	printResult(nums)
}

func populate(length int) []int {
	var nums = make([]int, length)
	var num int

	for i := 0; i < length; i++ {
		_, err := fmt.Scanf("%d", &num)

		if err != nil {
			panic(err)
		}

		nums[i] = num
	}

	return nums
}

func reverse(s []int) {
	n := len(s)
	i := 0
	j := n - 1

	var t int

	for i < j {
		// swap
		t = s[i]
		s[i] = s[j]
		s[j] = t

		i++
		j--
	}
}

func printResult(s []int) {
	n := len(s)

	for i := 0; i < n-1; i++ {
		fmt.Printf("%d ", s[i])
	}

	fmt.Printf("%d\n", s[n-1])
}
