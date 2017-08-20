package main

import "fmt"

func main() {
	var n int
	_, err := fmt.Scan(&n)

	if err != nil {
		panic(err)
	}

	sum := 0
	var num int

	for i := 0; i < n; i++ {
		_, err := fmt.Scanf("%v", &num)

		if err != nil {
			panic(err)
		}

		sum += num
	}

	fmt.Println(sum)
}
