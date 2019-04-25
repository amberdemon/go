package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("looping")

	i := 1

	for i < 3 {
		fmt.Println("i is ", i)
		i++
	}

	for j := 0; j < 10; j++ {
		fmt.Println("j is ", j)
	}

	for {
		fmt.Println("break")
		break
	}

	for k := 0; k < 5; k++ {
		if k%2 == 0 {
			continue
		}
		fmt.Println("k is ", k)
	}

	checkDOW()
}

func checkDOW() {
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Its a weekend")
	default:
		fmt.Println(":( its a weekday")
	}
}
