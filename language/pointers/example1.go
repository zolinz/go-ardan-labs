package main

import "fmt"

func main() {

	count := 10

	fmt.Println(count)

	increment(count)

	fmt.Println("in main" , count)


	increment2()



}


func increment(inc int){

	inc++
	fmt.Println(inc)
}


func increment2(inc *int){

  &inc++
   fmt.Println(inc)
}
