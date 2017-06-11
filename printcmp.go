package helpers

import "fmt"

// print & compare two integers
func PrintCmpInts(actual int, expected int) bool {
	fmt.Printf("actual: %d\n", actual)
	fmt.Printf("expected: %d\n", expected)
	result := actual == expected
	fmt.Println(result)
	return result
}

// print & compare two strings
func PrintCmpStrings(actual string, expected string) bool {
	fmt.Printf("actual: %s\n", actual)
	fmt.Printf("expected: %s\n", expected)
	result := actual == expected
	fmt.Println(result)
	return result
}
