package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := rand.Intn(10)
	y := rand.Intn(10)
	fmt.Printf("random x and y is %v %v\t\t", x, y)
	/*
		if x < 4 && y < 4 {
			fmt.Println("x and y is less than 4")
		} else if x > 6 && y > 6 {
			fmt.Println("x and y is greater than 6")
		} else if x <= 4 && y <= 4 {
			fmt.Println("x and y is less than or equal to 4")
		} else if x != 5 {
			fmt.Println("y is not equal to 5")
		} else {
			fmt.Println("none of the previous were met")
		}
	*/
	switch {
	case x < 4 && y < 4:
		fmt.Println("x and y is less than 4")
	case x > 6 && y > 6:
		fmt.Println("x and y is greater than 6")
	case x <= 4 && y <= 4:
		fmt.Println("x and y is less than or equal to 4")
	case x != 5:
		fmt.Println("y is not equal to 5")
	default:
		fmt.Println("none of the previous were met")
	}
}
