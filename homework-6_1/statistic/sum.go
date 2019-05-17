package statistic

/*
 Syntax Go. Homework 6.1 - Standard library part 2
Supplement the code from the testing section with a
function for counting the sum of transferred elements
and a test for this function.
 Nick Nikulin, dated Мay 17, 2019
 Teacher, Sergey Iryupin
*/

import "fmt"

func Sum(x []float64) float64 {

	totalx := 0.0
	for _, valuex := range x {
		totalx += valuex
	}
	return totalx
}

func main() {
	x := []float64{1, 2, 3, 4, 5}
	fmt.Println(Sum(x))
}
