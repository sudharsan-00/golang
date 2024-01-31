package main

import (
	"fmt"
)

func main() {
 x := make(map[string][]string)

 x["bond_james"]  = []string{"chocolate", "Movie", "adventure"}
 x["alex_carry"]  = []string{"icecreme", "cricket", "cycling"}
 x["leonel_ronaldo"] = []string{"juices", "football", "playing"}

 for k,v := range x{
  fmt.Println(k)
  for i, v2 := range v{
    fmt.Println(i, v2)
  }
  delete(x, "alex_carry")
  fmt.Printf("%#v\n",x)
 }
   
}