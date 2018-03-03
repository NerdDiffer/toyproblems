package codewars

func Abs(x int) int {
	if x < 0 {
		return x * -1
	} else {
		return x
	}
}

func Gcdi(x, y int) int {
	if y == 0 {
		return x
	}

	remainder := x % y

	if remainder == 0 {
		return Abs(y) // base case
	} else {
		return Gcdi(y, remainder)
	}
}

func Lcmu(x, y int) int {
	a := Abs(x * y)
	b := Gcdi(x, y)
	return Abs(a / b)
}

func Som(x, y int) int {
	return x + y
}

func Maxi(x, y int) int {
	if x >= y {
		return x
	} else {
		return y
	}
}

func Mini(x, y int) int {
	if x <= y {
		return x
	} else {
		return y
	}
}

type FParam func(int, int) int

func OperArray(f FParam, arr []int, init int) []int {
	length := len(arr)
	result := make([]int, length)
	prev := init

	for i := 0; i < length; i++ {
		val := f(prev, arr[i])
		result[i] = val
		prev = val
	}

	return result
}
