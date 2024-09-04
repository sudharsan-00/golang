package main

import "fmt"

func main() {
    foo()

	hash("James")

	s := star("alex")
	fmt.Println(s)
    
	s1, n := moon("James",20)
	fmt.Println(s1, n)
	
}
// no params, no returns
func foo() {
	fmt.Println("i am from foo")
}

// 1 params, no return
func hash(s string){
    fmt.Println("My name is ",s)
}

// 1 params, 1 return
func star(s string) string{
	return fmt.Sprint("My name is ",s)
}

// 2 params, 2 returns
func moon(name string, age int) (string,int){
  age *= 7
  return fmt.Sprint(name, " is this old in paralell universe"),age
} 

/*
func syntax

no params, no returns
1 params, no return
1 params, 1 return
2 params, 2 returns
*/
