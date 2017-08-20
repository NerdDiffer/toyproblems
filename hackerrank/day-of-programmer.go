package main

import (
	"fmt"
	"time"
)

func main() {
	var year int
	_, err := fmt.Scan(&year)

	if err != nil {
		panic(err)
	}

	var p *time.Time

	if year == 1918 {
		d, _ := makeDate(year, 26)
		p = &d
	} else {
		isLeapYear := leapYear(year)

		if isLeapYear {
			d, _ := makeDate(year, 12)
			p = &d
		} else {
			d, _ := makeDate(year, 13)
			p = &d
		}
	}

	dateStr := formatDate(*p)
	fmt.Println(dateStr)
}

func makeDate(y, d int) (time.Time, error) {
	const format = "2006-01-02"
	m := int(time.September)
	val := fmt.Sprintf("%d-%02d-%02d", y, m, d)
	date, err := time.Parse(format, val)
	return date, err
}

func formatDate(d time.Time) string {
	return fmt.Sprintf("%02d.%02d.%d", d.Day(), d.Month(), d.Year())
}

func leapYear(year int) bool {
	if year < 1918 {
		return year%4 == 0
	} else if year > 1918 {
		return year%400 == 0 || (year%4 == 0 && year%100 != 0)
	} else {
		return false
	}
}
