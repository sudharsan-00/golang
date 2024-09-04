package main

import "fmt"

func main(){
    fmt.Printf("%T\n",add)
    fmt.Printf("%T\n",sub)
    fmt.Printf("%T\n",doMath)

	x := doMath(42,5,add)
	fmt.Println(x)

	y := doMath(42,16,sub)
	fmt.Println(y)

}

func doMath(a int, b int, f func(a int,b int)int)int{
	return f(a,b)
}
func add(a int, b int) int {
	return a+b
}

func sub(a int, b int) int{
	return a-b
}