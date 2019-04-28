package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello arrays")

	var ar1 [5]int
	fmt.Println(ar1)

	ar1[3] = 123
	fmt.Println(ar1)

	fmt.Println(len(ar1))

	//declare and init on one line
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println(b)

	//now slices

}
