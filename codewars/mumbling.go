package main

import (
	"bytes"
	"fmt"
	"strings"
)

func Accum(s string) string {
	var accum []string

	for ind, val := range s {
		repeated := repeat(ind+1, string(val))
		accum = append(accum, repeated)
	}

	return strings.Join(accum, "-")
}

func repeat(repetitions int, value string) string {
	var buffer bytes.Buffer

	buffer.WriteString(strings.ToUpper(value))

	for i := 1; i < repetitions; i++ {
		buffer.WriteString(strings.ToLower(value))
	}

	return buffer.String()
}

func main() {
	fmt.Println(Accum("ZpglnRxqenU"))
}
