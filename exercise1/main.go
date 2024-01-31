package main

import (
	"fmt"
)

func main() {
  x := []string {"james", "Bond", "Shaken, not stirred"}
  y := []string {"Miss", "Moneypenny", "Im 008"}

  

  fmt.Println(x)
  fmt.Println(y)
  
  z := [][]string{x,y}

  for i, v := range z{
    fmt.Println(i, v)
    for a, b := range v{
      fmt.Println(a,b)
    }
    }  
   
}