package validator

// Unique checks whether the values in the provided slice are unique.
//
// Example:
//
//	values := []int{1, 2, 3, 4, 5}
//	result := Unique(values)
//	// result will be true because all values in the slice are unique.
func Unique[T comparable](values []T) bool {
	uniqueValues := make(map[T]bool)

	for _, value := range values {
		uniqueValues[value] = true
	}

	return len(values) == len(uniqueValues)
}
