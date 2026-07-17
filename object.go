package lodash

import (
	"sort"
	"strings"

	"cmp"
)

// Keys returns the keys of the map in unspecified order.
func Keys[K comparable, V any](m map[K]V) []K {
	r := make([]K, 0, len(m))
	for k := range m {
		r = append(r, k)
	}
	return r
}

// SortedKeys returns the keys of the map sorted in ascending order.
func SortedKeys[K cmp.Ordered, V any](m map[K]V) []K {
	r := Keys(m)
	sort.Slice(r, func(i, j int) bool { return r[i] < r[j] })
	return r
}

// Values returns the values of the map in unspecified order.
func Values[K comparable, V any](m map[K]V) []V {
	r := make([]V, 0, len(m))
	for _, v := range m {
		r = append(r, v)
	}
	return r
}

// Entry is a single key/value pair from a map.
type Entry[K comparable, V any] struct {
	Key   K
	Value V
}

// Entries returns the key/value pairs of the map in unspecified order.
func Entries[K comparable, V any](m map[K]V) []Entry[K, V] {
	r := make([]Entry[K, V], 0, len(m))
	for k, v := range m {
		r = append(r, Entry[K, V]{Key: k, Value: v})
	}
	return r
}

// FromEntries builds a map from a slice of entries.
func FromEntries[K comparable, V any](entries []Entry[K, V]) map[K]V {
	r := make(map[K]V, len(entries))
	for _, e := range entries {
		r[e.Key] = e.Value
	}
	return r
}

// Pick returns a new map containing only the given keys that are present in m.
func Pick[K comparable, V any](m map[K]V, keys ...K) map[K]V {
	r := make(map[K]V, len(keys))
	for _, k := range keys {
		if v, ok := m[k]; ok {
			r[k] = v
		}
	}
	return r
}

// PickBy returns a new map containing entries for which fn returns true.
func PickBy[K comparable, V any](m map[K]V, fn func(K, V) bool) map[K]V {
	r := make(map[K]V)
	for k, v := range m {
		if fn(k, v) {
			r[k] = v
		}
	}
	return r
}

// Omit returns a new map with the given keys removed.
func Omit[K comparable, V any](m map[K]V, keys ...K) map[K]V {
	drop := make(map[K]struct{}, len(keys))
	for _, k := range keys {
		drop[k] = struct{}{}
	}
	r := make(map[K]V, len(m))
	for k, v := range m {
		if _, ok := drop[k]; !ok {
			r[k] = v
		}
	}
	return r
}

// OmitBy returns a new map excluding entries for which fn returns true.
func OmitBy[K comparable, V any](m map[K]V, fn func(K, V) bool) map[K]V {
	r := make(map[K]V)
	for k, v := range m {
		if !fn(k, v) {
			r[k] = v
		}
	}
	return r
}

// MapKeys returns a new map whose keys are produced by applying fn to each
// original key/value pair.
func MapKeys[K comparable, V any, K2 comparable](m map[K]V, fn func(K, V) K2) map[K2]V {
	r := make(map[K2]V, len(m))
	for k, v := range m {
		r[fn(k, v)] = v
	}
	return r
}

// MapValues returns a new map whose values are produced by applying fn to each
// original key/value pair.
func MapValues[K comparable, V any, V2 any](m map[K]V, fn func(K, V) V2) map[K]V2 {
	r := make(map[K]V2, len(m))
	for k, v := range m {
		r[k] = fn(k, v)
	}
	return r
}

// Invert returns a new map with keys and values swapped. When multiple keys map
// to the same value, later iterations overwrite earlier ones.
func Invert[K comparable, V comparable](m map[K]V) map[V]K {
	r := make(map[V]K, len(m))
	for k, v := range m {
		r[v] = k
	}
	return r
}

// Merge returns a new map combining all provided maps. Later maps override
// earlier ones for duplicate keys.
func Merge[K comparable, V any](maps ...map[K]V) map[K]V {
	r := make(map[K]V)
	for _, m := range maps {
		for k, v := range m {
			r[k] = v
		}
	}
	return r
}

// Assign copies all key/value pairs from sources into dst, overwriting existing
// keys, and returns dst.
func Assign[K comparable, V any](dst map[K]V, sources ...map[K]V) map[K]V {
	for _, m := range sources {
		for k, v := range m {
			dst[k] = v
		}
	}
	return dst
}

// Get retrieves a nested value from a map[string]any using a dot-separated path
// such as "a.b.c". The second return value reports whether the full path was
// resolved.
func Get(m map[string]any, path string) (any, bool) {
	if path == "" {
		return nil, false
	}
	parts := strings.Split(path, ".")
	var cur any = m
	for _, p := range parts {
		asMap, ok := cur.(map[string]any)
		if !ok {
			return nil, false
		}
		cur, ok = asMap[p]
		if !ok {
			return nil, false
		}
	}
	return cur, true
}

// Has reports whether the dot-separated path resolves to a value in the nested
// map.
func Has(m map[string]any, path string) bool {
	_, ok := Get(m, path)
	return ok
}
