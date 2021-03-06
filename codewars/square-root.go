package codewars

import "math"

func SquareOrSquareRoot(input []int) []int {
	result := make([]int, len(input))

	for i := 0; i < len(input); i += 1 {
		num := input[i]
		sqrt := math.Sqrt(float64(num))

		// is the square root an integer?
		if sqrt == math.Floor(sqrt) || sqrt == math.Ceil(sqrt) {
			// then add the square root to the result
			result[i] = int(sqrt)
		} else {
			// then add the square to the result
			result[i] = num * num
		}
	}

	return result
}
