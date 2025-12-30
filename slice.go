package tools

func ForEach[T any](inputSlice []T, fn func(value T, idx int)) {
	for idx, value := range inputSlice {
		fn(value, idx)
	}
}

func Filter[T any](inputSlice []T, fn func(value T, idx int) bool) []T {
	d := make([]T, 0, len(inputSlice))
	for idx, value := range inputSlice {
		if fn(value, idx) {
			d = append(d, value)
		}
	}

	return d
}

func Map[T1, T2 any](inputSlice []T1, fn func(value T1, idx int) T2) []T2 {
	d := make([]T2, 0, len(inputSlice))
	for idx, value := range inputSlice {
		d = append(d, fn(value, idx))
	}

	return d
}

func Find[T any](inputSlice []T, fn func(value T, idx int) bool) (T, bool) {
	for idx, value := range inputSlice {
		if fn(value, idx) {
			return value, true
		}
	}

	var notFound T
	return notFound, false
}

func Copy[T any](inputSlice []T) []T {
	d := make([]T, len(inputSlice))
	for idx, value := range inputSlice {
		d[idx] = value
	}

	return d
}
