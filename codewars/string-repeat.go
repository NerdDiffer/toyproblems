package codewars

import "bytes"

func RepeatStr(repititions int, value string) string {
	var buffer bytes.Buffer

	for i := 0; i < repititions; i++ {
		buffer.WriteString(value)
	}

	return buffer.String()
}
