package main

import (
	"fmt"
)

func main() {
	course := "docker course"
	fmt.Println("course is ", course)
	changeCourse(&course)
	fmt.Println("course is now ", course)
}

func changeCourse(course *string) string {
	*course = "newCourse"
	return *course
}
