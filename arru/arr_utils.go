package arru

func Filter[T any](arr []T, passFunc func(T) bool) []T {

	if passFunc == nil {
		return arr
	}
	result := make([]T, 0, len(arr))
	for i := range arr {
		t := arr[i]
		if !passFunc(t) {
			continue
		}
		result = append(result, t)
	}
	return result
}

func Conv2Map[K comparable, AT any](arr []AT, keyFunc func(AT) K) map[K]AT {
	if keyFunc == nil {
		return make(map[K]AT)
	}
	mp := make(map[K]AT, len(arr))
	for i := range arr {
		k := keyFunc(arr[i])
		mp[k] = arr[i]
	}
	return mp
}
