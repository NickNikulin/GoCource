package main

/*
Syntax Go. Homework 7.1 - goroutines and chanels.
Task: Change code and spin the spinner for 10 seconds
Nick Nikulin, dated Мay 25, 2019
Teacher, Sergey Iryupin
*/

import (
	"fmt"
	"time"
)

func main() {
	go spinner(50 * time.Millisecond)
	time.Sleep(10 * time.Second)
}

func spinner(delay time.Duration) {
	for {
		for _, r := range "-\\|/" {
			fmt.Printf("%c\r", r)
			time.Sleep(delay)
		}
	}
}
