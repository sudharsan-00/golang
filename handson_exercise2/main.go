package main

import (
	"fmt"
)
//Map struct

type person struct {
  first string
  last string
  favIC []string
}

func main() {
    p1 := person{
      first: "James",
      last: "Bond",
      favIC: []string{"chocolate", "vannila", "blackcurrent"},
    }
    
    p2 := person{
      first: "Alex",
      last: "carry",
      favIC: []string{"straberry", "mango", "pista"},
    }

    fmt.Println(p1)
    fmt.Println(p2)

    fmt.Println(p1.favIC)
    fmt.Println(p2.favIC)

    for _,v := range p1.favIC{
      fmt.Println(p1.first, "favorite is",v)
    }
    for _,v := range p2.favIC{
      fmt.Println(p2.first, "favorite is",v)
    }

    m := map[string]person{
      p1.last: p1,
      p2.last: p2,
    }

    for _,v := range m{
      fmt.Println(v)
    for _,v1:= range v.favIC{
      fmt.Println(v.first, v.last, v1)
    }
   }
}