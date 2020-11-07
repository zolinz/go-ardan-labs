package main

import (
	"example.com/hello/pkg/hello"
	"fmt"
)

var x bool
var y = 43
var z int = 21

func main() {
	fmt.Println("Hello, playground")
	fmt.Print(hello.Proverb())

	fmt.Println(x)
	x = false
	fmt.Println(x)

	s := "Hello, 世界"
	fmt.Printf("%s\n", s)
	fmt.Printf("%q\n", s)
	fmt.Printf("%x\n", s)
	fmt.Printf("---%x\n", "世")
	for i :=0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
}


