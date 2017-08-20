package main

import "fmt"

func main() {
	var n, k int
	_, _ = fmt.Scanf("%v", &n)
	_, _ = fmt.Scanf("%v", &k)

	var highest int
	_, _ = fmt.Scanf("%v", &highest)

	var curr int

	for i := 1; i < n; i++ {
		_, _ = fmt.Scanf("%v", &curr)
		highest = max(highest, curr)
	}

	diff := highest - k
	fmt.Println(max(diff, 0))
}

func max(a, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}
