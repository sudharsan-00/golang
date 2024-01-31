package main

import (
	"fmt"
)

/*
	func main() {
		inputString := "hello"
		letterToCount := "l"
		count := 0

		for _, char := range inputString {
			if strings.ToLower(string(char)) == strings.ToLower(letterToCount) {
				count++
			}
		}
		fmt.Printf("The letter %s appears %d times in the string.\n", letterToCount, count)
	}
*/
func main() {
	a := "Hello, World!"
	reversed := reverseString(a)
	fmt.Printf("Original: %s\n", a)
	fmt.Printf("Reversed: %s\n", reversed)
}

func reverseString(a string) string {

	b := []rune(a)
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	reversedString := string(b)
	return reversedString
}