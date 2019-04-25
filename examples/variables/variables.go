package main

import (
	"fmt"
)

func main() {
	var i, j int = 1, 2
	fmt.Printf("i is %d and j is %d\n", i, j)

	course := "docker course"
	fmt.Println("course is ", course)
	changeCourse(&course)
	fmt.Println("course is now ", course)

	const noChange string = "I cant change"
	fmt.Println(noChange)
	// noChange = "oh no"
}

func changeCourse(course *string) string {
	*course = "newCourse"
	return *course
}
