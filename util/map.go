package util

func GetMapKeys[T1 comparable, T2 any](inputMap map[T1]T2) []T1 {
	result := make([]T1, len(inputMap))
	idx := 0
	for key := range inputMap {
		result[idx] = key
		idx += 1
	}

	return result
}

func GetOrFill[T1 comparable, T2 any](inputMap map[T1]T2, key T1, extras ...T2) T2 {
	var defaultValue T2
	if len(extras) == 0 {
		defaultValue = *new(T2)
	} else {
		defaultValue = extras[0]
	}

	result, ok := inputMap[key]
	if !ok {
		inputMap[key] = defaultValue
		result = defaultValue
	}
	return result
}
