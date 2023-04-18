package enumerable

func Slice[T any](values ...T) []T { return values }

func Each[T any](handler func(T), slice []T) {
	for index := range slice {
		handler(slice[index])
	}
}

func Map[T, Y any](handler func(T) Y, slice []T) (ret []Y) {
	Each(func(value T) { ret = append(ret, handler(value)) }, slice)

	return ret
}

func Select[T any](handler func(T) bool, slice []T) (ret []T) {
	Each(func(value T) {
		if handler(value) {
			ret = append(ret, value)
		}
	}, slice)

	return ret
}

func Reject[T any](handler func(T) bool, slice []T) (ret []T) {
	return Select(func(value T) bool { return !handler(value) }, slice)
}
