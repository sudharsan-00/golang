package main

import "fmt"

func main() {
	x := 5
	for i := 0; i < 5; i++ {
		fmt.Printf("%v \n", i) 
	}

	for x < 10{
		fmt.Printf("The numbers are %v \n", x)
		x++
	}

}