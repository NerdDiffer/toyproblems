package codewars

import "fmt"

func IsTriangle(a, b, c int) bool {
	return (a+b > c) && (a+c > b) && (b+c > a)
}

func runIsTriangle() {
	fmt.Println(IsTriangle(4, 2, 3))
}
