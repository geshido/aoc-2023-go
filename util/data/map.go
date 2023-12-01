package data

func Map[T1, T2 any](list []T1, fn func(T1) T2) []T2 {
	res := make([]T2, len(list))

	for idx, v := range list {
		res[idx] = fn(v)
	}

	return res
}
