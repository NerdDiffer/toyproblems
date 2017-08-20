package main

import (
	"bytes"
	"fmt"
)

func main() {
	var n int
	fmt.Scanf("%v\n%v", &n)
	makeStaircase(n)
}

func repeat(limit int, ch string) string {
	var buf bytes.Buffer

	for i := 0; i < limit; i++ {
		buf.WriteString(ch)
	}

	return buf.String()
}

func makeStaircase(n int) {
	for i := 1; i <= n; i++ {
		lpad := repeat(n-i, " ")
		step := repeat(i, "#")
		fmt.Printf("%s%s\n", lpad, step)
	}

	fmt.Println() // need that blank line at end!
}
