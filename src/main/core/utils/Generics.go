package utils

func Map[T any, R any](list []T, fn func(T) R) []R {
	result := make([]R, len(list))
	for i, v := range list {
		result[i] = fn(v)
	}
	return result
}

func ChunkBy[T any](input []T, chunkQuantity int) [][]*T {
	if len(input) == 0 {
		return [][]*T{}
	}

	size := (len(input) + (chunkQuantity - 1)) / chunkQuantity

	divided := make([][]*T, chunkQuantity)

	for i, v := range input {
		index := i / size
		divided[index] = append(divided[index], &v)
	}

	return divided
}
