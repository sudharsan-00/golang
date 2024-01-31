package main

import (
	"fmt"
	"github.com/sudharsan-00/puppy"
)

func main() {
	puppy.From1()
	puppy.From2()
	fmt.Println(puppy.Bark())
	fmt.Println(puppy.Barks()) 
	fmt.Println(puppy.Bigbark())
	fmt.Println(puppy.Bigbarks())
}