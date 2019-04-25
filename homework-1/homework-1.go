package main

import "fmt"

func main() {
	fmt.Println("Сколько рублей хотите поменять?")
	var rubl float64
	fmt.Scanln(&rubl)
	const dollar float64 = 64
	dollsum := (rubl / dollar)
	fmt.Println("Получите", dollsum, "$ распишитесь!")
}
