package main

import "fmt"

func main() {
	fmt.Println("Сколько рублей хотите поменять?")
	var rub float64
	fmt.Scanln(&rub)
	const dollar float64 = 64
	dollAmount := (rub / dollar)
	fmt.Println("Получите", dollAmount, "$ распишитесь!")
}
