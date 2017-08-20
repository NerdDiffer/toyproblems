package main

import "fmt"

func main() {
	var n, p int
	_, err := fmt.Scanf("%v\n%v", &n, &p)

	if err != nil {
		panic(err)
	}

	var pagesFromBeginning int

	if isEven(p) {
		pagesFromBeginning = p / 2
	} else {
		pagesFromBeginning = (p - 1) / 2
	}

	var pagesFromEnd int

	if isEven(n) {
		if isEven(p) {
			pagesFromEnd = (n - p) / 2
		} else {
			pagesFromEnd = (n - p + 1) / 2
		}
	} else {
		pagesFromEnd = (n - p) / 2
	}

	if pagesFromBeginning < pagesFromEnd {
		fmt.Println(pagesFromBeginning)
	} else {
		fmt.Println(pagesFromEnd)
	}
}

func isEven(num int) bool {
	return num%2 == 0
}
