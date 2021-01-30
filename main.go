package main

import (

	"fmt"
)

var x bool
var y = 43
var z int = 21

func main() {
	fmt.Println("Hello, playground")


	fmt.Println(x)
	x = false
	fmt.Println(x)

	s := "Hello, Zoli"
	fmt.Printf("%s\n", s)
	fmt.Printf("%q\n", s)
	fmt.Printf("%x\n", s)
	fmt.Printf("---%x\n", "Zoli")
	for i :=0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
}


