package leetcode

import (
	"fmt"
)

func combine_r(currInd, nextNum, maxLen, maxNum int, currCombo []int, allCombos *[][]int) {
	if currInd == maxLen {
		comboVals := append([]int(nil), currCombo...)
		*allCombos = append(*allCombos, comboVals)
		return
	} else if nextNum > maxNum {
		return
	}

	currCombo[currInd] = nextNum

	combine_r(currInd+1, nextNum+1, maxLen, maxNum, currCombo, allCombos)
	combine_r(currInd+0, nextNum+1, maxLen, maxNum, currCombo, allCombos)
}

func combine(maxNum, maxLen int) [][]int {
	var allCombos [][]int
	currCombo := make([]int, maxLen)

	combine_r(0, 1, maxLen, maxNum, currCombo, &allCombos)

	return allCombos
}

// change package name & this function name to `main` for quick & dirty QA
func run_combine() {
	combos := combine(4, 2)
	fmt.Println("--- result:\n", combos)
}
