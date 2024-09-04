package main

import (
	"fmt"
	"strconv"
)

type book struct {
	title string
}

func (b book) String() string {
	return fmt.Sprint("The book name is ",b.title);
}

type count int

func (c count) String() string{
	return fmt.Sprint("This is the number ", strconv.Itoa(int(c)))
}

func main(){
	b := book{
		title : "The harry potter",
	}

	var n count = 42

   fmt.Println(b)
   fmt.Println(n)
}