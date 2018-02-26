package leetcode

import "math"

func selfDividingNumbers(left int, right int) []int {
	results := make([]int, 0)

	for n := left; n <= right; n++ {
		if n%10 == 0 {
			continue
		}

		if isSelfDivisible(n) {
			results = append(results, n)
		}
	}

	return results
}

func isSelfDivisible(num int) bool {
	n := howManyDigits(num)
	var d int

	for p := 0; p < n; p++ {
		d = digitAtPlace(num, p)

		if d == 0 || num%d != 0 {
			return false
		}
	}

	return true
}

// https://stackoverflow.com/a/1306751/2908123
func howManyDigits(number int) int {
	f := float64(number)
	log := math.Log10(f)
	return int(log) + 1
}

// https://stackoverflow.com/a/39644726/2908123
func digitAtPlace(number, place int) int {
	d := math.Pow10(place)
	t := number / int(d) // move target digit to the 1's place
	return t % 10
}
