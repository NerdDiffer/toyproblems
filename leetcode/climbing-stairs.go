// https://leetcode.com/problems/climbing-stairs/
package main

import (
	"toyproblems/toyproblemhelpers"
)

func climbStairs(n int) int {
	if n < 1 {
		return 0
	} else if n < 3 {
		return n
	} else {
		cache := make([]int, n)
		cache[0] = 1
		cache[1] = 2

		for i := 2; i < n; i++ {
			cache[i] = cache[i-1] + cache[i-2]
		}

		return cache[n-1]
	}
}

func main() {
	test(3, 3)
	test(4, 5)
	test(5, 8)
}

func test(input, expected int) bool {
	actual := climbStairs(input)
	return toyproblemhelpers.PrintCmpInts(actual, expected)
}
