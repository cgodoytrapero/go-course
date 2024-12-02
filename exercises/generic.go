package exercises

func AnyOf[T any](items []T, criteria func(T) bool) bool {

	for _, item := range items {
		if criteria(item) {
			return true
		}
	}
	return false
}
