package main

/*
Syntax Go. Homework 7.2 - goroutines and chanels.
Task:Change the program conveyor. Change the number of transmitted values.
Ensure the correct completion of all goroutines.
Nick Nikulin, dated Мay 25, 2019
Teacher, Sergey Iryupin
*/

import (
	"fmt"
)

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	// генерация
	go func() {
		for x := 0; x < 10; x++ {
			naturals <- x
		}
		close(naturals)
	}()

	// возведение в квадрат
	go func() {
		for x := range naturals {
			squares <- x * x
		}
		close(squares)
	}()

	// печать
	for x := range squares {
		fmt.Println(x)
	}
}
