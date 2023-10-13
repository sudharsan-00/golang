package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	ch1 := make(chan string)
	ch2 := make(chan string)

	d1 := time.Duration(rand.Int63n(250))
	d2 := time.Duration(rand.Int63n(250))

	go func() {
		time.Sleep(d1 * time.Millisecond)
		ch1 <- "sudharsan"
	}()

	go func() {
		time.Sleep(d2 * time.Millisecond)
		ch2 <- "varshini"
	}()

	select {
	case v1 := <-ch1:
		fmt.Println("Name of the king", v1)
	case v2 := <-ch2:
		fmt.Println("name of the queen", v2)
	}
}
