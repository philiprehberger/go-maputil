// Package maputil provides generic map utilities for Go.
//
// All functions leverage Go generics to work with any map types
// that satisfy the required constraints.
package maputil

import (
	"cmp"
	"slices"
)

// Filter returns a new map containing only the entries for which
// the predicate function returns true.
func Filter[K comparable, V any](m map[K]V, predicate func(K, V) bool) map[K]V {
	result := make(map[K]V)
	for k, v := range m {
		if predicate(k, v) {
			result[k] = v
		}
	}
	return result
}

// Map transforms the values of a map using the given transform function,
// returning a new map with the same keys and transformed values.
func Map[K comparable, V any, R any](m map[K]V, transform func(K, V) R) map[K]R {
	result := make(map[K]R, len(m))
	for k, v := range m {
		result[k] = transform(k, v)
	}
	return result
}

// MapKeys transforms the keys of a map using the given function,
// returning a new map with transformed keys and the same values.
// If multiple old keys map to the same new key, the last one wins
// (iteration order is non-deterministic).
func MapKeys[K comparable, V any, NK comparable](m map[K]V, fn func(K) NK) map[NK]V {
	result := make(map[NK]V, len(m))
	for k, v := range m {
		result[fn(k)] = v
	}
	return result
}

// Merge combines multiple maps into one. When the same key appears
// in more than one map, the value from the last map wins.
func Merge[K comparable, V any](maps ...map[K]V) map[K]V {
	result := make(map[K]V)
	for _, m := range maps {
		for k, v := range m {
			result[k] = v
		}
	}
	return result
}

// MergeWith combines multiple maps into one, using the provided
// conflict resolution function when the same key appears in more
// than one map. The conflictFn receives the key, the existing value,
// and the new value, and returns the value to keep.
func MergeWith[K comparable, V any](conflictFn func(K, V, V) V, maps ...map[K]V) map[K]V {
	result := make(map[K]V)
	for _, m := range maps {
		for k, v := range m {
			if existing, ok := result[k]; ok {
				result[k] = conflictFn(k, existing, v)
			} else {
				result[k] = v
			}
		}
	}
	return result
}

// Pick returns a new map containing only the entries whose keys
// are in the provided list. Keys not present in the original map
// are silently skipped.
func Pick[K comparable, V any](m map[K]V, keys ...K) map[K]V {
	result := make(map[K]V)
	for _, k := range keys {
		if v, ok := m[k]; ok {
			result[k] = v
		}
	}
	return result
}

// Omit returns a new map containing all entries except those whose
// keys are in the provided list.
func Omit[K comparable, V any](m map[K]V, keys ...K) map[K]V {
	exclude := make(map[K]struct{}, len(keys))
	for _, k := range keys {
		exclude[k] = struct{}{}
	}
	result := make(map[K]V)
	for k, v := range m {
		if _, skip := exclude[k]; !skip {
			result[k] = v
		}
	}
	return result
}

// Invert swaps the keys and values of a map. If multiple keys map
// to the same value, the last one wins (iteration order is non-deterministic).
func Invert[K comparable, V comparable](m map[K]V) map[V]K {
	result := make(map[V]K, len(m))
	for k, v := range m {
		result[v] = k
	}
	return result
}

// Keys returns a slice of all keys in the map. The order is not guaranteed.
func Keys[K comparable, V any](m map[K]V) []K {
	result := make([]K, 0, len(m))
	for k := range m {
		result = append(result, k)
	}
	return result
}

// SortedKeys returns a sorted slice of all keys in the map.
func SortedKeys[K cmp.Ordered, V any](m map[K]V) []K {
	result := make([]K, 0, len(m))
	for k := range m {
		result = append(result, k)
	}
	slices.Sort(result)
	return result
}

// Values returns a slice of all values in the map. The order is not guaranteed.
func Values[K comparable, V any](m map[K]V) []V {
	result := make([]V, 0, len(m))
	for _, v := range m {
		result = append(result, v)
	}
	return result
}

// Contains reports whether the map contains the given key.
func Contains[K comparable, V any](m map[K]V, key K) bool {
	_, ok := m[key]
	return ok
}

// Size returns the number of entries in the map.
// It is equivalent to len(m) but useful in generic function chains.
func Size[K comparable, V any](m map[K]V) int {
	return len(m)
}
