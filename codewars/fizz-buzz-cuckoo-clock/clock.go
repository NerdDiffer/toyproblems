package main

import "fmt"

import (
	"strconv"
	"strings"
)

func FizzBuzzCuckooClock(time string) (result string) {
	hh, mm := parse_time(time)

	if hh > 12 {
		hh -= 12
	} else if hh == 0 {
		hh += 12
	}

	if mm == 0 {
		var cuckoos []string
		hh_as_int := int(hh)

		for i := 0; i < hh_as_int; i++ {
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

func parse_time(time string) (hh, mm int64) {
	hhmm := strings.Split(time, ":")
	var hours_and_minutes [2]int64

	for i := 0; i < 2; i++ {
		num, _ := strconv.ParseInt(hhmm[i], 10, 64) // ignore 2nd return value
		hours_and_minutes[i] = num
	}

	hh = hours_and_minutes[0]
	mm = hours_and_minutes[1]
	return
}

func main() {
	fmt.Println(FizzBuzzCuckooClock("21:00"))
}
