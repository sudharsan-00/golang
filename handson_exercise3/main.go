package main

import (
	"fmt"
)

// Embed struct
type engine struct {
    electric bool
  }

type vehicle struct{
  engine
  make string
  model string
  doors int
  color string
}
func main(){
v1 := vehicle{
  engine : engine{
    electric: true,
  },
  make: "ford",
  model: "Mustang",
  doors: 4,
  color: "black",   
}
v2 := vehicle{
  engine : engine{
    electric: false,
  },
  make: "benz",
  model: "AMG",
  doors: 4,
  color: "white",   
}

fmt.Println(v1)
fmt.Println(v2)

fmt.Println(v1.electric, v1.make, v1.model)
fmt.Println(v2.electric, v2.make, v2.model)

//fmt.Println(v1.engine.electric, v1.make, v1.model)
//fmt.Println(v2.engine.electric, v2.make, v2.model)

}