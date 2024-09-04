package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

type secretagent struct{
	person
	ltk bool
}

func main() {
sa1 := secretagent {
	person:	person{
	    first: "james",
		last:  "bond",
		age:   34,
	},
	ltk: true,
}

	p2 := person{
		first: "chris",
		last:  "hermsworth",
		age:   27,
	}

	fmt.Printf("%T \t %#v", sa1, sa1)

	fmt.Println(sa1)
	fmt.Println(p2)
    fmt.Println(sa1.person)
	fmt.Println(sa1.first, sa1.last, sa1.ltk)
}
