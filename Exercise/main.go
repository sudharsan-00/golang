package main

import "fmt"

func main() {
	xs := make([]int, 5)

	for i := 0; i < 5; i++ {
		xs[i] = i
	}
	for i, v := range xs {
		fmt.Printf("%v - %T - %v\n", v, v, i)
	}
	fmt.Println("****************")

	xd := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	for i, v := range xd {
		fmt.Printf("%T - %v - %v\n", v, v, i)
	}

	fmt.Println("******************")
	
	fmt.Println(xd[:5])
	fmt.Println(xd[5:])
	fmt.Println(xd[2:7])
	fmt.Println(xd[1:6])
	
	fmt.Println("******************")
	
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
    fmt.Println(x)
	x = append(x, 51)
    fmt.Println(x)
	
	x = append(x, 53,54,55)
    fmt.Println(x)
	
	y := []int{56,57,58,59,60}
	
	x = append(x,y...)
    fmt.Println(x)
	
	fmt.Println("******************")
	//slicing of slice

    z := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
    fmt.Println(z)
	z = append(z[:3],z[6:]...)
    fmt.Println(z)
	
}