package util

import (
	"testing"
)

func TestAverage(t *testing.T) {
	data := []float64{1, 2, 3, 4, 5}
	avg := Average(data)
	if avg != 3.0 {
		t.Errorf("Expected 3.0, got %f", avg)
	}

	if Average(nil) != 0 {
		t.Error("Expected 0 for nil slice")
	}
}

func TestSum(t *testing.T) {
	data := []float64{1, 2, 3}
	sum := Sum(data)
	if sum != 6.0 {
		t.Errorf("Expected 6.0, got %f", sum)
	}
}

func TestStdDev(t *testing.T) {
	// Population std dev of 2, 4, 4, 4, 5, 5, 7, 9
	// Mean = 5
	// Variance = (9+1+1+1+0+0+4+16)/8 = 32/8 = 4
	// StdDev = 2
	data := []float64{2, 4, 4, 4, 5, 5, 7, 9}
	stdDev := StdDev(data)
	if stdDev != 2.0 {
		t.Errorf("Expected 2.0, got %f", stdDev)
	}

	if StdDev(nil) != 0 {
		t.Error("Expected 0 for nil slice")
	}
}

func TestMax(t *testing.T) {
	if Max(1, 5, 3) != 5 {
		t.Error("Expected 5")
	}
	if Max(1.5, 0.5) != 1.5 {
		t.Error("Expected 1.5")
	}
	// Max of nothing? The function signature is `values ...T`.
	// Loop `for _, v := range values` won't run. `maxValue` (zero value) returned.
	if Max[int]() != 0 {
		t.Error("Expected 0")
	}
}

func TestMin(t *testing.T) {
	if Min(1, 5, 3) != 1 {
		t.Error("Expected 1")
	}
	// Min of nothing -> zero value
	if Min[int]() != 0 {
		t.Error("Expected 0")
	}
	// Note: The implementation of Min initializes `minValue` to zero value (T).
	// If all values are positive, Min will return 0 if passed empty? No.
	// `var minValue T` -> 0.
	// Loop: `if v < minValue`. 5 < 0 is false.
	// So Min(5, 10) returns 0?
	// Let's check the code.
}
