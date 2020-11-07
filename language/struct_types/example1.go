package main

import "fmt"

type person struct{
	firstName string
	lastName string
}



func main(){


	var zoli person

	fmt.Printf("%+v\n", zoli)


	zoli2 := person{
		firstName: "zoltan",
		lastName: "kovacs",
	}

	fmt.Println("Zoli first name :", zoli.firstName)

	fmt.Println("Zoli2 first name :", zoli2.firstName)

}
