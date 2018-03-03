package codewars

import (
	"strconv"
	"strings"
)

func FizzBuzzCuckooClock(time string) (result string) {
	hh, mm := parse(time)

	if hh > 12 {
		hh -= 12
	} else if hh == 0 {
		hh += 12
	}

	if mm == 0 {
		var cuckoos []string

		for i := 0; i < hh; i++ {
			cuckoos = append(cuckoos, "Cuckoo")
		}

		result = strings.Join(cuckoos[:], " ")
	} else if mm == 30 {
		result = "Cuckoo"
	} else if mm%5 == 0 && mm%3 == 0 {
		result = "Fizz Buzz"
	} else if mm%5 == 0 {
		result = "Buzz"
	} else if mm%3 == 0 {
		result = "Fizz"
	} else {
		result = "tick"
	}

	return
}

func parse(time string) (int, int) {
	hhmm := strings.Split(time, ":")
	var hoursMinutes [2]int64

	for i := 0; i < len(hoursMinutes); i++ {
		num, _ := strconv.ParseInt(hhmm[i], 10, 64)
		hoursMinutes[i] = num
	}

	return int(hoursMinutes[0]), int(hoursMinutes[1])
}
