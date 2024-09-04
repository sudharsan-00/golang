package main

import (
	"fmt"
)
//ananymous struct
func main(){
  p1 :=struct {
    first string
    friends map[string]int
    favDrinks []string
}{
    first: "james",
    friends: map[string]int{
          "alex" : 23,
          "david": 22,
    },
    favDrinks: []string{"Mango", "straberry", "papaya"},
  }

  for k,v := range p1.friends{
    fmt.Println(p1.first, "- friends -",k,v)
  }

  for k,v := range p1.favDrinks{
    fmt.Println(p1.first, "- drinks -",k,v)
  }
}

