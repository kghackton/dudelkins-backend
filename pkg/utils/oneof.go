package utils

func OneOf(value int, array []int) bool {
	for _, element := range array {
		if value == element {
			return true
		}
	}
	return false
}
