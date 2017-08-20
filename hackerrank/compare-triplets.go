package main

import "fmt"

func main() {
	alice := make([]int, 3)
	bob := make([]int, 3)
	populate(alice)
	populate(bob)

	sumAlice, sumBob := cmp(alice, bob)

	fmt.Printf("%d %d\n", sumAlice, sumBob)
}

func cmp(scoresA, scoresB []int) (int, int) {
	sumA := 0
	sumB := 0

	var a, b int

	for i := 0; i < 3; i++ {
		a = scoresA[i]
		b = scoresB[i]
		if a > b {
			sumA += 1
		} else if b > a {
			sumB += 1
		} else {
			// no op
		}
	}

	return sumA, sumB
}

func populate(scores []int) {
	var score int

	for i := 0; i < len(scores); i++ {
		_, err := fmt.Scanf("%v", &score)

		if err != nil {
			panic(err)
		}

		scores[i] = score
	}
}
