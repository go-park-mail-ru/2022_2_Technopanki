package utils

func HasArrayElement[T comparable](elem T, array []T) bool {
	for _, val := range array {
		if val == elem {
			return false
		}
	}
	return true
}
