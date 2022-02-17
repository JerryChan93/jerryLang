package commom

func InArray[T comparable](arr []T, needle T) bool {

	for i := 0; i < len(arr); i++ {
		var elem = arr[i]
		if elem == needle {
			return true
		}

	}

	return false
}
