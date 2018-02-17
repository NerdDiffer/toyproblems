package leetcode

// map of upper-case chars to integers (I, V, X, L, C, D, M)
var charVals = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

func RomanToInt(input string) int {
	sum := 0
	n := len(input)

	if n == 0 {
		return sum
	}

	for i := 0; i < n-1; i++ {
		curr := string(input[i])
		next := string(input[i+1])

		if charVals[curr] >= charVals[next] {
			sum += charVals[curr]
		} else {
			sum -= charVals[curr]
		}
	}

	lastChar := string(input[n-1])
	sum += charVals[lastChar]
	return sum
}
