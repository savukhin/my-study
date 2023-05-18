package utils

func Reverse[T any](source []T) []T {
	result := make([]T, len(source))
	for i, elem := range source {
		result[len(source)-1-i] = elem
	}

	return result
}
