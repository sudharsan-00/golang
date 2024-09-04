package main

import (
	
	"fmt"
	"io/ioutil"
	"log"
)

/*
func main() {
	xb, err := readfile("poem.txt")
	if err != nil {
		log.Fatalf("Readfile in main : %s", err)
	}
	fmt.Println(xb)
	fmt.Println(string(xb))
}

func readfile(fileName string) ([]byte, error) {
	xb, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("error in readfile func : %s", &err)
	}
	return xb, err
}

*/


func main(){

	xb, err := readfile("poem.txt")
	if err != nil{
		log.Fatalf("%s",err)
	}
	fmt.Println(xb)
	fmt.Println(string(xb))

}

func readfile(fileName string) ([]byte, error){
	xb, err := ioutil.ReadFile(fileName)
	if err != nil{
		return nil, fmt.Errorf("%s",err)
	}
	return xb,
}