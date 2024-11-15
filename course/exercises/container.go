package exercises

type Container struct {
	Data []int
}

func (container *Container) Add(value int) {
	container.Data = append(container.Data, value)
}

func (container *Container) Exist(value int) bool {
	result := false
	for _, item := range container.Data {
		if value == item {
			result = true
		}
	}
	return result
}

func (container *Container) Delete(value int) bool {
	result := false
	for i, item := range container.Data {
		if value == item {
			container.Data = append(container.Data[:i], container.Data[i+1:]...)
			result = true
		}
	}
	return result
}
