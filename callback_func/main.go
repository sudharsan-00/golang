package main

import "fmt"

func main() {
	x := doMath(45, 16, add)
	fmt.Println(x)

	y := doMath(45, 16, sub)
	fmt.Println(y)

	fmt.Printf("%T\n",add)
	fmt.Printf("%T\n",sub)
	fmt.Printf("%T\n",doMath)
}
func doMath(a int, b int, f func(int, int) int) int{
	return f(a, b)
}
func add(a int, b int) int {
	return a + b
}

func sub(a int, b int) int {
	return a - b
}

/* func main(){
   fmt.Println(printSquare(square, 3))
}

func printSquare(f func(int) int, a int) string{
	x := f(a)
	return fmt.Sprintf("The number %v squared is %v",a,x)
}

func square(n int) int{
	return n * n
}
*/