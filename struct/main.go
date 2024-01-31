package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func main() {
	p1 := person{
		first: "james",
		last:  "bond",
		age:   34,
	}

	p2 := person{
		first: "chris",
		last:  "hermsworth",
		age:   27,
	}

	fmt.Printf("%T \t %#v", p1, p1)

	fmt.Println(p1)
	fmt.Println(p2)

}
