package main

import (
	"fmt"
	"math"
)

var a int
// changing positive number to negative number
func main() {
	fmt.Println("Enter the number :")
	fmt.Scan(&a)

	d := math.Sqrt(math.Pow(float64(a),2))

	fmt.Println(int(d))
}
