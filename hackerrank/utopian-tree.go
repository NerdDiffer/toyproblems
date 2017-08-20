package main

import "fmt"

func main() {
	var numTests int
	_, _ = fmt.Scanf("%d", &numTests)

	var numCycles int

	for i := 0; i < numTests; i++ {
		_, err := fmt.Scanf("%d", &numCycles)

		if err != nil {
			panic(err)
		}

		fmt.Println(getHeight(numCycles))
	}
}

func getHeight(n int) int {
	height := 1

	for i := 0; i < n; i++ {
		if i%2 == 0 {
			height *= 2
		} else {
			height += 1
		}
	}

	return height
}
