package exercises

func AnyOf[T any](items []T, criteria func(T) bool) bool {
	for _, item := range items {
		if criteria(item) {
			return true
		}
	}
	return false
}

func FindIf[T any](items []T, criteria func(T) bool) int {
	for index, item := range items {
		if criteria(item) {
			return index
		}
	}
	return -1
}

func AdjacentFind[T any](items []T, criteria func(T, T) bool) int {
	for index, item := range items {
		if index+1 >= len(items) {
			return -1
		}
		if criteria(item, items[index+1]) {
			return index
		}
	}
	return -1
}

func Equal[T comparable](left []T, right []T) bool {
	if (left == nil) != (right == nil) {
		return false
	}

	if len(left) != len(right) {
		return false
	}

	for index, item := range left {
		if item != right[index] {
			return false
		}
	}
	return true
}

func ReplaceIf[T comparable](items []T, element T, criteria func(T) bool) int {
	if items == nil {
		return -1
	}

	for index, item := range items {
		if criteria(item) {
			items[index] = element
			return index
		}
	}

	return -1
}

func RemoveIf[T any](items *[]T, criteria func(T) bool) bool {
	if items == nil {
		return false
	}

	tmp := make([]T, len(*items))
	i := 0
	for _, item := range *items {
		if criteria(item) == true {
			tmp[i] = item
			i++
		}
	}
	items = &tmp
	return true
}
