package main

import "fmt"

func main() {
	a := []int {1,2,3,4,5,6}
	b := []int{}
	c := []int{}
	sum := 0
	rum := 0

	for i := 0; i < len(a) ; i++ {
		if a[i]%2==0 {
			b = append(b, a[i])
		} else {
			c = append(c, a[i])
		}
	}
		for _,v := range(b){
			sum += v
		}
		for _,v := range(c){
			rum += v
		}
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(sum)
	fmt.Println(rum)
} 