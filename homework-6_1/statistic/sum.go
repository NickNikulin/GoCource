package statistic

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
