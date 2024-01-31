package main

import "fmt"

func main() {
	/* xs := []string{"Almond Biscotti Café", "Banana Pudding ", "Balsamic Strawberry (GF)", "Bittersweet Chocolate (GF)", "Blueberry Pancake (GF)", "Bourbon Turtle (GF)", "Browned Butter Cookie Dough", "Chocolate Covered Black Cherry (GF)", "Chocolate Fudge Brownie", "Chocolate Peanut Butter (GF)", "Coffee (GF)", "Cookies & Cream", "Fresh Basil (GF)", "Garden Mint Chip (GF)", "Lavender Lemon Honey (GF)", "Lemon Bar", "Madagascar Vanilla (GF)", "Matcha (GF)", "Midnight Chocolate Sorbet (GF, V)", "Non-Dairy Chocolate Peanut Butter (GF, V)", "Non-Dairy Coconut Matcha (GF, V)", "Non-Dairy Sunbutter Cinnamon (GF, V)", "Orange Cream (GF) ", "Peanut Butter Cookie Dough", "Raspberry Sorbet (GF, V)", "Salty Caramel (GF)", "Slate Slate Different", "Strawberry Lemonade Sorbet (GF, V)", "Vanilla Caramel Blondie", "Vietnamese Cinnamon (GF)", "Wolverine Tracks (GF)"}
	fmt.Println(xs)
	fmt.Printf("%T\n", xs)
	// slicing of slice
	fmt.Printf("xs - %#v\n", xs[1:5])

	// printing slice value with index
	for i, v := range xs {
		fmt.Printf("index %v - value %v\n", i, v)
	}
	fmt.Println("------------------")

	xi := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	xi = append(xi[:4], xi[5:]...)
	fmt.Printf("xi - %#v\n", xi)
	fmt.Println("------------------")*/
    
    //slice with make function
	xd := make([]int, 0, 10)

	xd = append(xd, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println(len(xd))
	fmt.Println(cap(xd))
	fmt.Println(xd)

	xd = append(xd, 11, 12, 13, 14)
	fmt.Println(len(xd))
	fmt.Println(cap(xd))
	fmt.Println(xd)

	fmt.Println("------------------")
	// multidimentional slice

	xv := []string{"sudhars", "chocolate", "oreo", "lion"}
	xc := []string{"varshini", "blackcurrent", "vennila", "tiger"}

	xss := [][]string{xv, xc}

	fmt.Println(xss)

}
