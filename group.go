package maputil

// GroupBy groups the elements of a slice by a key derived from each element
// using the provided key function. The result is a map from keys to slices
// of elements that share that key.
func GroupBy[T any, K comparable](slice []T, keyFn func(T) K) map[K][]T {
	result := make(map[K][]T)
	for _, item := range slice {
		k := keyFn(item)
		result[k] = append(result[k], item)
	}
	return result
}

// CountBy counts the elements of a slice per group, where the group key
// is derived from each element using the provided key function.
func CountBy[T any, K comparable](slice []T, keyFn func(T) K) map[K]int {
	result := make(map[K]int)
	for _, item := range slice {
		result[keyFn(item)]++
	}
	return result
}

// UniqueBy returns a map from group key to a single element per group.
// When multiple elements share the same key, the last element wins.
func UniqueBy[T any, K comparable](slice []T, keyFn func(T) K) map[K]T {
	result := make(map[K]T)
	for _, item := range slice {
		result[keyFn(item)] = item
	}
	return result
}
