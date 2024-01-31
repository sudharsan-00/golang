package main

import (
	"fmt"
	"math/rand"
)

func main() {
	xi := []int{42, 43, 45, 46, 47}
	for i, v := range xi {
		fmt.Printf("index %v \t value %v\n", i, v)
	}

	m := map[string]int{
		"james": 22,
		"jack": 23,
	}
	for k, v := range m {
		fmt.Printf("key %v \t value %v\n", k, v)
	}

	fmt.Println("***************")
	c := 0
	for i := 0; i < 100; i++ {
		if x := rand.Intn(5); x == 3 {
			fmt.Printf("iteration count %v \t total count %v \t x is %v\n", i, c, x)
			c++
		}
	}
}
