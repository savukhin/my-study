package utils

func Difference[T comparable, K any, Q any](bigger map[T]K, smaller map[T]Q) map[T]K {
	result := make(map[T]K)

	for key, val := range bigger {
		_, ok := smaller[key]
		if !ok {
			result[key] = val
		}
	}

	return result
}

func DifferenceArrays[T comparable](bigger []T, smaller []T) []T {
	first := make(map[T]bool)
	second := make(map[T]bool)

	for _, val := range bigger {
		first[val] = true
	}
	for _, val := range smaller {
		second[val] = true
	}

	diff := Difference(first, second)
	result := make([]T, 0)

	for key, _ := range diff {
		result = append(result, key)
	}

	return result
}
