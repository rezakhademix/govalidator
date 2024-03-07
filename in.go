package validator

// In checks value under validation must be included in the given list of values
func In[T comparable](value T, acceptableValues ...T) bool {
	for i := range acceptableValues {
		if value == acceptableValues[i] {
			return true
		}
	}

	return false
}
