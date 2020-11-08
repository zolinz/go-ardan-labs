package main

import "fmt"

func main() {

	count := 10

	fmt.Println(count)

	increment(count)

	fmt.Println("in main" , count)


	increment2(&count)

	fmt.Println("in main" , count)

}


func increment(inc int){

	inc++
	fmt.Println(inc)
}


//receive the address into a pointer var
func increment2(inc *int){

	*inc++
   fmt.Println(*inc)
}
