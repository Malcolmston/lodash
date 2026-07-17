package lodash

import (
	"cmp"
	"math/rand"
	"sort"
)

// Map returns a new slice populated with the results of calling fn on every
// element of the input slice. The callback receives each element.
func Map[T, R any](s []T, fn func(T) R) []R {
	r := make([]R, len(s))
	for i, v := range s {
		r[i] = fn(v)
	}
	return r
}

// MapI is like Map but the callback also receives the element index.
func MapI[T, R any](s []T, fn func(T, int) R) []R {
	r := make([]R, len(s))
	for i, v := range s {
		r[i] = fn(v, i)
	}
	return r
}

// Filter returns a new slice containing every element for which fn returns true.
func Filter[T any](s []T, fn func(T) bool) []T {
	r := make([]T, 0, len(s))
	for _, v := range s {
		if fn(v) {
			r = append(r, v)
		}
	}
	return r
}

// Reject is the opposite of Filter: it keeps elements for which fn returns false.
func Reject[T any](s []T, fn func(T) bool) []T {
	r := make([]T, 0, len(s))
	for _, v := range s {
		if !fn(v) {
			r = append(r, v)
		}
	}
	return r
}

// Reduce boils the slice down to a single accumulated value by iterating from
// left to right, seeded with init.
func Reduce[T, A any](s []T, fn func(acc A, cur T) A, init A) A {
	acc := init
	for _, v := range s {
		acc = fn(acc, v)
	}
	return acc
}

// ReduceRight is like Reduce but iterates from right to left.
func ReduceRight[T, A any](s []T, fn func(acc A, cur T) A, init A) A {
	acc := init
	for i := len(s) - 1; i >= 0; i-- {
		acc = fn(acc, s[i])
	}
	return acc
}

// ForEach invokes fn once for each element of the slice.
func ForEach[T any](s []T, fn func(T)) {
	for _, v := range s {
		fn(v)
	}
}

// ForEachI invokes fn once for each element, passing the element and its index.
func ForEachI[T any](s []T, fn func(T, int)) {
	for i, v := range s {
		fn(v, i)
	}
}

// Find returns the first element for which fn returns true, along with true.
// If no element matches, it returns the zero value and false.
func Find[T any](s []T, fn func(T) bool) (T, bool) {
	for _, v := range s {
		if fn(v) {
			return v, true
		}
	}
	var zero T
	return zero, false
}

// FindLast returns the last element for which fn returns true.
func FindLast[T any](s []T, fn func(T) bool) (T, bool) {
	for i := len(s) - 1; i >= 0; i-- {
		if fn(s[i]) {
			return s[i], true
		}
	}
	var zero T
	return zero, false
}

// FindIndex returns the index of the first element for which fn returns true,
// or -1 if none match.
func FindIndex[T any](s []T, fn func(T) bool) int {
	for i, v := range s {
		if fn(v) {
			return i
		}
	}
	return -1
}

// FindLastIndex returns the index of the last element for which fn returns true,
// or -1 if none match.
func FindLastIndex[T any](s []T, fn func(T) bool) int {
	for i := len(s) - 1; i >= 0; i-- {
		if fn(s[i]) {
			return i
		}
	}
	return -1
}

// Every reports whether fn returns true for every element. It returns true for
// an empty slice.
func Every[T any](s []T, fn func(T) bool) bool {
	for _, v := range s {
		if !fn(v) {
			return false
		}
	}
	return true
}

// Some reports whether fn returns true for at least one element. It returns
// false for an empty slice.
func Some[T any](s []T, fn func(T) bool) bool {
	for _, v := range s {
		if fn(v) {
			return true
		}
	}
	return false
}

// Includes reports whether value is present in the slice.
func Includes[T comparable](s []T, value T) bool {
	for _, v := range s {
		if v == value {
			return true
		}
	}
	return false
}

// IndexOf returns the index of the first occurrence of value, or -1.
func IndexOf[T comparable](s []T, value T) int {
	for i, v := range s {
		if v == value {
			return i
		}
	}
	return -1
}

// LastIndexOf returns the index of the last occurrence of value, or -1.
func LastIndexOf[T comparable](s []T, value T) int {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == value {
			return i
		}
	}
	return -1
}

// GroupBy groups the slice elements into a map keyed by the result of fn.
func GroupBy[T any, K comparable](s []T, fn func(T) K) map[K][]T {
	r := make(map[K][]T)
	for _, v := range s {
		k := fn(v)
		r[k] = append(r[k], v)
	}
	return r
}

// KeyBy builds a map keyed by the result of fn. When multiple elements produce
// the same key, later elements overwrite earlier ones.
func KeyBy[T any, K comparable](s []T, fn func(T) K) map[K]T {
	r := make(map[K]T, len(s))
	for _, v := range s {
		r[fn(v)] = v
	}
	return r
}

// CountBy counts elements by the key returned by fn.
func CountBy[T any, K comparable](s []T, fn func(T) K) map[K]int {
	r := make(map[K]int)
	for _, v := range s {
		r[fn(v)]++
	}
	return r
}

// Partition splits the slice into two slices: elements for which fn is true and
// elements for which it is false.
func Partition[T any](s []T, fn func(T) bool) (truthy, falsy []T) {
	for _, v := range s {
		if fn(v) {
			truthy = append(truthy, v)
		} else {
			falsy = append(falsy, v)
		}
	}
	return truthy, falsy
}

// Uniq returns a new slice with duplicate values removed, preserving the order
// of first appearance.
func Uniq[T comparable](s []T) []T {
	seen := make(map[T]struct{}, len(s))
	r := make([]T, 0, len(s))
	for _, v := range s {
		if _, ok := seen[v]; !ok {
			seen[v] = struct{}{}
			r = append(r, v)
		}
	}
	return r
}

// UniqBy is like Uniq but uniqueness is determined by the key produced by fn.
func UniqBy[T any, K comparable](s []T, fn func(T) K) []T {
	seen := make(map[K]struct{}, len(s))
	r := make([]T, 0, len(s))
	for _, v := range s {
		k := fn(v)
		if _, ok := seen[k]; !ok {
			seen[k] = struct{}{}
			r = append(r, v)
		}
	}
	return r
}

// Chunk splits the slice into groups of at most size elements. It panics if
// size is less than 1.
func Chunk[T any](s []T, size int) [][]T {
	if size < 1 {
		panic("lodash: Chunk size must be at least 1")
	}
	r := make([][]T, 0, (len(s)+size-1)/size)
	for i := 0; i < len(s); i += size {
		end := i + size
		if end > len(s) {
			end = len(s)
		}
		chunk := make([]T, end-i)
		copy(chunk, s[i:end])
		r = append(r, chunk)
	}
	return r
}

// Flatten flattens a slice of slices by a single level.
func Flatten[T any](s [][]T) []T {
	n := 0
	for _, sub := range s {
		n += len(sub)
	}
	r := make([]T, 0, n)
	for _, sub := range s {
		r = append(r, sub...)
	}
	return r
}

// FlattenDeep recursively flattens a nested structure into a flat slice. The
// nested value must be either a T or a []any of the same shape.
func FlattenDeep[T any](s []any) []T {
	var r []T
	var walk func(items []any)
	walk = func(items []any) {
		for _, item := range items {
			switch v := item.(type) {
			case []any:
				walk(v)
			case T:
				r = append(r, v)
			}
		}
	}
	walk(s)
	return r
}

// Compact returns a new slice with all zero values removed.
func Compact[T comparable](s []T) []T {
	var zero T
	r := make([]T, 0, len(s))
	for _, v := range s {
		if v != zero {
			r = append(r, v)
		}
	}
	return r
}

// Reverse returns a new slice with the elements in reverse order. The input is
// not mutated.
func Reverse[T any](s []T) []T {
	r := make([]T, len(s))
	for i, v := range s {
		r[len(s)-1-i] = v
	}
	return r
}

// Pair holds two related values produced by Zip.
type Pair[A, B any] struct {
	First  A
	Second B
}

// Zip combines two slices into a slice of pairs. The result length equals the
// length of the shorter input.
func Zip[A, B any](a []A, b []B) []Pair[A, B] {
	n := len(a)
	if len(b) < n {
		n = len(b)
	}
	r := make([]Pair[A, B], n)
	for i := 0; i < n; i++ {
		r[i] = Pair[A, B]{First: a[i], Second: b[i]}
	}
	return r
}

// Unzip splits a slice of pairs back into two slices.
func Unzip[A, B any](pairs []Pair[A, B]) ([]A, []B) {
	a := make([]A, len(pairs))
	b := make([]B, len(pairs))
	for i, p := range pairs {
		a[i] = p.First
		b[i] = p.Second
	}
	return a, b
}

// Difference returns the elements of s that are not present in any of the
// others slices, preserving order and de-duplicating.
func Difference[T comparable](s []T, others ...[]T) []T {
	exclude := make(map[T]struct{})
	for _, o := range others {
		for _, v := range o {
			exclude[v] = struct{}{}
		}
	}
	seen := make(map[T]struct{})
	r := make([]T, 0, len(s))
	for _, v := range s {
		if _, bad := exclude[v]; bad {
			continue
		}
		if _, dup := seen[v]; dup {
			continue
		}
		seen[v] = struct{}{}
		r = append(r, v)
	}
	return r
}

// Intersection returns the unique elements common to all provided slices,
// ordered by their first appearance in the first slice.
func Intersection[T comparable](slices ...[]T) []T {
	if len(slices) == 0 {
		return nil
	}
	counts := make([]map[T]struct{}, len(slices))
	for i, s := range slices {
		m := make(map[T]struct{}, len(s))
		for _, v := range s {
			m[v] = struct{}{}
		}
		counts[i] = m
	}
	seen := make(map[T]struct{})
	var r []T
	for _, v := range slices[0] {
		if _, dup := seen[v]; dup {
			continue
		}
		inAll := true
		for i := 1; i < len(counts); i++ {
			if _, ok := counts[i][v]; !ok {
				inAll = false
				break
			}
		}
		if inAll {
			seen[v] = struct{}{}
			r = append(r, v)
		}
	}
	return r
}

// Union returns the unique elements from all provided slices, preserving the
// order of first appearance.
func Union[T comparable](slices ...[]T) []T {
	seen := make(map[T]struct{})
	var r []T
	for _, s := range slices {
		for _, v := range s {
			if _, ok := seen[v]; !ok {
				seen[v] = struct{}{}
				r = append(r, v)
			}
		}
	}
	return r
}

// Without returns a new slice excluding all the given values.
func Without[T comparable](s []T, values ...T) []T {
	exclude := make(map[T]struct{}, len(values))
	for _, v := range values {
		exclude[v] = struct{}{}
	}
	r := make([]T, 0, len(s))
	for _, v := range s {
		if _, ok := exclude[v]; !ok {
			r = append(r, v)
		}
	}
	return r
}

// Take returns the first n elements of the slice. If n exceeds the length, the
// whole slice is returned. Negative n yields an empty slice.
func Take[T any](s []T, n int) []T {
	if n < 0 {
		n = 0
	}
	if n > len(s) {
		n = len(s)
	}
	r := make([]T, n)
	copy(r, s[:n])
	return r
}

// TakeRight returns the last n elements of the slice.
func TakeRight[T any](s []T, n int) []T {
	if n < 0 {
		n = 0
	}
	if n > len(s) {
		n = len(s)
	}
	r := make([]T, n)
	copy(r, s[len(s)-n:])
	return r
}

// Drop returns the slice with the first n elements removed.
func Drop[T any](s []T, n int) []T {
	if n < 0 {
		n = 0
	}
	if n > len(s) {
		n = len(s)
	}
	r := make([]T, len(s)-n)
	copy(r, s[n:])
	return r
}

// DropRight returns the slice with the last n elements removed.
func DropRight[T any](s []T, n int) []T {
	if n < 0 {
		n = 0
	}
	if n > len(s) {
		n = len(s)
	}
	r := make([]T, len(s)-n)
	copy(r, s[:len(s)-n])
	return r
}

// Sample returns a single random element from the slice using the provided
// random source, making it deterministic when seeded. The second return value
// is false when the slice is empty.
func Sample[T any](s []T, rng *rand.Rand) (T, bool) {
	if len(s) == 0 {
		var zero T
		return zero, false
	}
	return s[rng.Intn(len(s))], true
}

// SampleN returns n unique random elements from the slice using the provided
// random source. If n exceeds the length, a shuffled copy of the whole slice is
// returned.
func SampleN[T any](s []T, n int, rng *rand.Rand) []T {
	if n < 0 {
		n = 0
	}
	if n > len(s) {
		n = len(s)
	}
	idx := rng.Perm(len(s))[:n]
	r := make([]T, n)
	for i, j := range idx {
		r[i] = s[j]
	}
	return r
}

// Shuffle returns a new slice with the elements randomly reordered using the
// provided random source. The input is not mutated.
func Shuffle[T any](s []T, rng *rand.Rand) []T {
	r := make([]T, len(s))
	copy(r, s)
	rng.Shuffle(len(r), func(i, j int) {
		r[i], r[j] = r[j], r[i]
	})
	return r
}

// SortBy returns a new slice sorted in ascending order by the key returned by
// fn. The sort is stable and the input is not mutated.
func SortBy[T any, K cmp.Ordered](s []T, fn func(T) K) []T {
	r := make([]T, len(s))
	copy(r, s)
	sort.SliceStable(r, func(i, j int) bool {
		return fn(r[i]) < fn(r[j])
	})
	return r
}

// Order describes a sort direction for OrderBy.
type Order int

const (
	// Asc sorts in ascending order.
	Asc Order = iota
	// Desc sorts in descending order.
	Desc
)

// OrderBy returns a new slice sorted by multiple key functions with per-key sort
// directions. The keys are compared in order until a difference is found. The
// number of orders may be shorter than keys; missing directions default to Asc.
func OrderBy[T any](s []T, keys []func(a, b T) int, orders []Order) []T {
	r := make([]T, len(s))
	copy(r, s)
	sort.SliceStable(r, func(i, j int) bool {
		for k, key := range keys {
			c := key(r[i], r[j])
			if c == 0 {
				continue
			}
			ord := Asc
			if k < len(orders) {
				ord = orders[k]
			}
			if ord == Desc {
				return c > 0
			}
			return c < 0
		}
		return false
	})
	return r
}

// Concat returns a new slice with all provided slices concatenated together.
func Concat[T any](slices ...[]T) []T {
	n := 0
	for _, s := range slices {
		n += len(s)
	}
	r := make([]T, 0, n)
	for _, s := range slices {
		r = append(r, s...)
	}
	return r
}

// Fill returns a new slice of length n where every element is value.
func Fill[T any](value T, n int) []T {
	if n < 0 {
		n = 0
	}
	r := make([]T, n)
	for i := range r {
		r[i] = value
	}
	return r
}
