package lodash

import "cmp"

// This file provides "collection" helpers that treat a map as a collection of
// values, mirroring the way lodash's collection methods accept objects as well
// as arrays. Unless noted, callbacks receive each value followed by its key.
//
// Because Go map iteration order is unspecified, helpers that return a slice
// produce their elements in an unspecified order, and helpers that pick a
// single element ("first match") are only deterministic when at most one
// element can satisfy the predicate.

// EveryMap reports whether pred returns true for every value/key pair in m.
// It returns true for an empty map.
func EveryMap[K comparable, V any](m map[K]V, pred func(V, K) bool) bool {
	for k, v := range m {
		if !pred(v, k) {
			return false
		}
	}
	return true
}

// SomeMap reports whether pred returns true for at least one value/key pair in
// m. It returns false for an empty map.
func SomeMap[K comparable, V any](m map[K]V, pred func(V, K) bool) bool {
	for k, v := range m {
		if pred(v, k) {
			return true
		}
	}
	return false
}

// NoneMap reports whether pred returns false for every value/key pair in m. It
// returns true for an empty map.
func NoneMap[K comparable, V any](m map[K]V, pred func(V, K) bool) bool {
	return !SomeMap(m, pred)
}

// FindMap returns a value for which pred returns true, along with true. When
// multiple pairs match, which one is returned is unspecified. It returns the
// zero value and false when nothing matches.
func FindMap[K comparable, V any](m map[K]V, pred func(V, K) bool) (V, bool) {
	for k, v := range m {
		if pred(v, k) {
			return v, true
		}
	}
	var zero V
	return zero, false
}

// FilterMap returns the values of m for which pred returns true. The result
// order is unspecified.
func FilterMap[K comparable, V any](m map[K]V, pred func(V, K) bool) []V {
	r := make([]V, 0, len(m))
	for k, v := range m {
		if pred(v, k) {
			r = append(r, v)
		}
	}
	return r
}

// RejectMap returns the values of m for which pred returns false. The result
// order is unspecified.
func RejectMap[K comparable, V any](m map[K]V, pred func(V, K) bool) []V {
	r := make([]V, 0, len(m))
	for k, v := range m {
		if !pred(v, k) {
			r = append(r, v)
		}
	}
	return r
}

// MapToSlice applies fn to every value/key pair in m and returns the results as
// a slice. The result order is unspecified. It mirrors lodash's _.map applied
// to an object.
func MapToSlice[K comparable, V, R any](m map[K]V, fn func(V, K) R) []R {
	r := make([]R, 0, len(m))
	for k, v := range m {
		r = append(r, fn(v, k))
	}
	return r
}

// ReduceMap boils m down to a single accumulated value by applying fn to each
// value/key pair, seeded with init. Because iteration order is unspecified, fn
// should be associative and commutative for a deterministic result.
func ReduceMap[K comparable, V, A any](m map[K]V, fn func(acc A, v V, k K) A, init A) A {
	acc := init
	for k, v := range m {
		acc = fn(acc, v, k)
	}
	return acc
}

// PartitionMap splits the values of m into two slices: truthy holds the values
// for which pred returns true and falsy holds the rest. Element order within
// each slice is unspecified.
func PartitionMap[K comparable, V any](m map[K]V, pred func(V, K) bool) (truthy, falsy []V) {
	for k, v := range m {
		if pred(v, k) {
			truthy = append(truthy, v)
		} else {
			falsy = append(falsy, v)
		}
	}
	return truthy, falsy
}

// GroupByMap groups the values of m by the key returned by fn, producing a map
// from group key to the slice of values in that group. Element order within
// each group is unspecified.
func GroupByMap[K comparable, V any, G comparable](m map[K]V, fn func(V) G) map[G][]V {
	r := make(map[G][]V)
	for _, v := range m {
		g := fn(v)
		r[g] = append(r[g], v)
	}
	return r
}

// IncludesValue reports whether value is one of the values stored in m. It
// mirrors lodash's _.includes applied to an object.
func IncludesValue[K, V comparable](m map[K]V, value V) bool {
	for _, v := range m {
		if v == value {
			return true
		}
	}
	return false
}

// MinByMap returns the value of m for which fn yields the smallest ordered key,
// along with true, or the zero value and false when m is empty. Ties are broken
// arbitrarily.
func MinByMap[K comparable, V any, R cmp.Ordered](m map[K]V, fn func(V) R) (V, bool) {
	var (
		best     V
		bestRank R
		found    bool
	)
	for _, v := range m {
		rank := fn(v)
		if !found || rank < bestRank {
			best, bestRank, found = v, rank, true
		}
	}
	return best, found
}

// MaxByMap returns the value of m for which fn yields the largest ordered key,
// along with true, or the zero value and false when m is empty. Ties are broken
// arbitrarily.
func MaxByMap[K comparable, V any, R cmp.Ordered](m map[K]V, fn func(V) R) (V, bool) {
	var (
		best     V
		bestRank R
		found    bool
	)
	for _, v := range m {
		rank := fn(v)
		if !found || rank > bestRank {
			best, bestRank, found = v, rank, true
		}
	}
	return best, found
}
