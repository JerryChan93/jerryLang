package commom

func Filter[T any](arr []T, in bool, f func(T) bool) []T {
	result := []T{}
	for _, elem := range arr {
		choose := f(elem)
		if (in && choose) || (!in && !choose) {
			result = append(result, elem)
		}
	}
	return result
}

func FilterIn[T any](arr []T, f func(T) bool) []T {
	return Filter(arr, true, f)
}
func FilterOut[T any](arr []T, f func(T) bool) []T {
	return Filter(arr, false, f)
}
