package validator

func In[T comparable](value T, acceptableValues ...T) bool {
	for i := range acceptableValues {
		if value == acceptableValues[i] {
			return true
		}
	}

	return false
}
