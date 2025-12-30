package tools

func Ptr[T any](t T) *T {
	temp := t
	return &temp
}

func GetPtrValOrDefault[T any](t *T, defaultValue T) T {
	if t == nil {
		return defaultValue
	}
	return *t
}

func GetPtrValOrZero[T any](t *T) T {
	var zero T
	return GetPtrValOrDefault(t, zero)
}

func Last[T any](s []T) T {
	if len(s) == 0 {
		var zero T
		return zero
	}
	return s[len(s)-1]
}
