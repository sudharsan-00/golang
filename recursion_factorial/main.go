package main

import "fmt"

func main() {
     x := factorial(4)
	 fmt.Println(x)

	 y := factloop(4)
	 fmt.Println(y)
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n - 1)  // recursion (func calling func itself)
}

func factloop(n int) int{   // recursion in loop
	total := 1
	for n > 0 {
		total *= n
		n--
	}
	return(total)
}