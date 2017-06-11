package main

import "fmt"

func IsTriangle(a, b, c int) bool {
	return (a+b > c) && (a+c > b) && (b+c > a)
}

func main() {
	fmt.Println(IsTriangle(4, 2, 3))
}
