package main

import "fmt"

func main() {
    foo()

	func() {
		fmt.Println("This is anonymous func")
	}()

	func(s string) {
		fmt.Printf("This is anonymous func returning my name %v", s)
	}("Alex")
}

func foo() {
	fmt.Println("This is foo")
}
