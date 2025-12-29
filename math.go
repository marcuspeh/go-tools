package util

import (
	"cmp"
	"math"
)

func Average(data []float64) float64 {
	if len(data) == 0 {
		return 0
	}

	sum := 0.0
	for _, v := range data {
		sum += v
	}
	return sum / float64(len(data))
}

func StdDev(data []float64) float64 {
	if len(data) == 0 {
		return 0
	}

	mean := Average(data)
	variance := 0.0
	for _, v := range data {
		diff := v - mean
		variance += diff * diff
	}
	return math.Sqrt(variance / float64(len(data)))
}

func Sum(data []float64) float64 {
	sum := 0.0
	for _, v := range data {
		sum += v
	}
	return sum
}

func Max[T cmp.Ordered](values ...T) T {
	var maxValue T
	for _, v := range values {
		if v > maxValue {
			maxValue = v
		}
	}

	return maxValue
}

func Min[T cmp.Ordered](values ...T) T {
	var minValue T
	for _, v := range values {
		if v < minValue {
			minValue = v
		}
	}

	return minValue
}
