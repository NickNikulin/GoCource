package main

/*
 * Syntax Go. Homework 2
 * Nick Nikulin, dated Apr 26, 2019
 */

import "fmt"

func main() {
	DevineByTwo()
}

func DevineByTwo() {
	fmt.Println("Введите целое число")
	var numb int64
	fmt.Scanln(&numb)

	if Check(numb, 2) {
		fmt.Println("Четное")
		return
	}

	fmt.Println("Нечетное")
}

func Check(num int64, dev int64) bool {
	if num%dev == 0 {
		return true
	} else {
		return false
	}
}
