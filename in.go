package validator

// In checks if the value under validation must be included in the given list of values.
//
// Example:
//
//	result := In("apple", "banana", "orange", "apple")
//	// result will be true because "apple" is included in the list of acceptable values.
func In[T comparable](value T, acceptableValues ...T) bool {
	for i := range acceptableValues {
		if value == acceptableValues[i] {
			return true
		}
	}

	return false
}
