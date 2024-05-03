package maputil

// MapCopy Copy map
func MapCopy[K comparable, V any](m1 map[K]V) map[K]V {
	m2 := make(map[K]V, len(m1))
	for k, v := range m2 {
		m2[k] = v
	}
	return m2
}

func Keys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func KeysV1[K comparable, V, NK any](m map[K]V, keyFunc func(K) NK) []NK {
	if keyFunc == nil {
		return []NK{}
	}
	keys := make([]NK, 0, len(m))
	for k := range m {
		keys = append(keys, keyFunc(k))
	}
	return keys
}

func Values[K comparable, V any](m map[K]V) []V {
	values := make([]V, 0, len(m))
	for _, v := range m {
		values = append(values, v)
	}
	return values
}

func ValuesV1[K comparable, V, NV any](m map[K]V, valFunc func(V) NV) []NV {
	if valFunc == nil {
		return []NV{}
	}
	values := make([]NV, 0, len(m))
	for _, v := range m {
		values = append(values, valFunc(v))
	}
	return values
}
