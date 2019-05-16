package statistic

import "testing"

func TestSum(t *testing.T) {
	var v float64
	v = Sum([]float64{1, 2, 3, 4, 5})
	if v != 15 {
		t.Error("Expected 15, got ", v)
	}
}
