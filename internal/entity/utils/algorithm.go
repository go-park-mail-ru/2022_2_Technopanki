package utils

func HasStringArrayElement(elem string, array []string) bool {
	for _, val := range array {
		if val == elem {
			return true
		}
	}
	return false
}
