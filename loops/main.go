package main

import (
	"fmt"
	"strings"
)

// linear search
func main() {
	a := "Hello"
	strSlice := strings.Split(a, "")
	target := "l"
	count := 0

	for i := 0; i < len(strSlice); i++ {
		if strSlice[i] == target {
			count += 1
		}
	}
	fmt.Println(count)
}
