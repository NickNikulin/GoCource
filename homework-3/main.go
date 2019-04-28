package main

/*
 * Syntax Go. Homework 2. Fleet monitoring/data structures ver. 1.0
 * Nick Nikulin, dated Apr 28, 2019
 */

import "fmt"

type car struct {
	carBrand    string  /* марка авто */
	yearIssue   int     /* год выпуска */
	laggageCap  float64 /* обьем багажника */
	engineRun   string  /* запущен ли двигатель */
	openWindows string  /* открыты ли окна */
	laggageLoad float64 /* насколько заполнен обьем багажника */
}

type truck struct {
	truckBrand  string  /* марка грузовика */
	yearIssue   int     /* год выпуска */
	laggageCap  float64 /* обьем кузова */
	engineRun   string  /* запущен ли двигатель */
	openWindows string  /* открыты ли окна */
	bodyLoad    float64 /* насколько заполнен обьем кузова */
	bodyLow     string  /* Кузов опущен, кузов поднят */

}

func main() {
	carСontrol()
	truckControl()
}

func carСontrol() {
	var freeCar = car{"Лада", 2019, 250, "двигатель заглушен", "окна закрыты", 0.6}
	fmt.Println(freeCar.carBrand)
	fmt.Println(freeCar.engineRun)

	freeCar.engineRun = "двигатель запущен"          // изменю значение
	fmt.Println(freeCar.carBrand, freeCar.engineRun) // Лада двигатель запущен

}

func truckControl() {
	var busyTruck = truck{"Камаз", 2010, 65115, "двигатель заглушен", "окна закрыты", 60000, "Кузов поднят"}
	fmt.Println(busyTruck.truckBrand)
	fmt.Println(busyTruck.bodyLow)

	busyTruck.bodyLow = "кузов опущен"                   // изменю значение
	fmt.Println(busyTruck.truckBrand, busyTruck.bodyLow) // Камаз кузов опущен

}
