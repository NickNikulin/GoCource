package main

/*
 Syntax Go. Homework 4 - Methods and interfaces
 A simple experiment Bruce Lee Style consisting of two packages
 - brain and geek.
 Nick Nikulin, dated Мay 04, 2019
 Teacher, Sergey Iryupin
*/

import (
	"fmt"
	"gocourse1/homework-4/brain"
	"gocourse1/homework-4/geek"
)

func main() {
	var doIt geek.Maker
	fmt.Println(doIt)

	doIt = &brain.EmptyCupBrain{}
	fmt.Println(doIt)
	doIt.Make()

	doIt = &brain.BeLikeWaterBrain{}
	fmt.Println(doIt)
	doIt.Make()

}
