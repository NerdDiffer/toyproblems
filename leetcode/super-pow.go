package leetcode

func SuperPow_v1(base int, exps []int) int {
	pow := 1
	maxExp := numerizeExpArr(exps)

	for exp := 1; exp <= maxExp; exp++ {
		pow *= base
	}

	return pow % 1337
}

func numerizeExpArr(arr []int) int {
	var exp int
	placeVal := 1
	length := len(arr)

	for i := length - 1; i > -1; i-- {
		exp += arr[i] * placeVal
		placeVal *= 10
	}

	return exp
}

const LEET = 1337

func SuperPow(base int, exps []int) int {
	if len(exps) == 0 {
		return 1
	}

	// pop and save ref to last val
	lastNum, exps := exps[len(exps)-1], exps[:len(exps)-1]

	moreBase := SuperPow(base, exps)

	return powerMod(moreBase, 10, LEET) * powerMod(base, lastNum, LEET) % LEET
}

func powerMod(base, exp, mod int) int {
	power := 1
	base %= mod

	for i := 0; i < exp; i++ {
		power *= base
		power %= mod
	}

	return power
}
