package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println("end")
		for j := 1; j < 6; j++ {
			fmt.Printf("outer loop %v \t\t\t inner loop %v\n", i, j)
		}
	}
}
