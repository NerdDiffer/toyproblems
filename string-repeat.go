package main

import (
	"bytes"
	"fmt"
)

func RepeatStr(repititions int, value string) string {
	var buffer bytes.Buffer

	for i := 0; i < repititions; i++ {
		buffer.WriteString(value)
	}

	return buffer.String()
}

func main() {
	fmt.Println(RepeatStr(6, "I"))
	fmt.Println(RepeatStr(5, "Hello"))
}
