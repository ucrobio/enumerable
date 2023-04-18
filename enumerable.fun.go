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
