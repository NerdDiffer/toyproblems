package leetcode

func CanPlaceFlowers(flowerbed []int, n int) bool {
	prevFlowerInd := -2

	numOpen := 0
	length := len(flowerbed)

	isPottable := func(currInd int) bool {
		if currInd == length-1 {
			return currInd-prevFlowerInd >= 2
		} else if currInd == 0 {
			return flowerbed[currInd+1] != 1
		} else {
			return flowerbed[currInd-1] != 1 && flowerbed[currInd+1] != 1
		}
	}

	for currInd, plotVal := range flowerbed {
		if plotVal == 1 {
			if currInd-prevFlowerInd < 2 {
				return false
			}

			prevFlowerInd = currInd
		} else if isPottable(currInd) {
			flowerbed[currInd] = 1
			numOpen += 1
			prevFlowerInd = currInd
		}
	}

	return numOpen >= n
}
