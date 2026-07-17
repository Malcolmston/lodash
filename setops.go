package lodash

import (
	"cmp"
	"sort"
	"strings"
)

// Head returns the first element of the slice and true, or the zero value and
// false when the slice is empty.
func Head[T any](s []T) (T, bool) {
	if len(s) == 0 {
		var zero T
		return zero, false
	}
	return s[0], true
}

// First is an alias for Head: it returns the first element of the slice.
func First[T any](s []T) (T, bool) { return Head(s) }

// Last returns the final element of the slice and true, or the zero value and
// false when the slice is empty.
func Last[T any](s []T) (T, bool) {
	if len(s) == 0 {
		var zero T
		return zero, false
	}
	return s[len(s)-1], true
}

// Nth returns the element at index n and true. Negative indices count from the
// end (-1 is the last element). An out-of-range index yields the zero value and
// false.
func Nth[T any](s []T, n int) (T, bool) {
	if n < 0 {
		n += len(s)
	}
	if n < 0 || n >= len(s) {
		var zero T
		return zero, false
	}
	return s[n], true
}

// Initial returns a new slice containing all but the last element.
func Initial[T any](s []T) []T {
	if len(s) <= 1 {
		return []T{}
	}
	r := make([]T, len(s)-1)
	copy(r, s[:len(s)-1])
	return r
}

// Tail returns a new slice containing all but the first element.
func Tail[T any](s []T) []T {
	if len(s) <= 1 {
		return []T{}
	}
	r := make([]T, len(s)-1)
	copy(r, s[1:])
	return r
}

// Slice returns the portion of s between start (inclusive) and end (exclusive).
// Negative indices count from the end. Out-of-range bounds are clamped. The
// input is not mutated.
func Slice[T any](s []T, start, end int) []T {
	n := len(s)
	if start < 0 {
		start += n
	}
	if end < 0 {
		end += n
	}
	if start < 0 {
		start = 0
	}
	if end > n {
		end = n
	}
	if start >= end {
		return []T{}
	}
	r := make([]T, end-start)
	copy(r, s[start:end])
	return r
}

// Join concatenates the string forms of the elements using sep. Elements are
// rendered with ToString.
func Join[T any](s []T, sep string) string {
	parts := make([]string, len(s))
	for i, v := range s {
		parts[i] = ToString(v)
	}
	return strings.Join(parts, sep)
}

// DifferenceBy returns the elements of s whose key (as produced by fn) does not
// appear among the keys of any of the others slices, de-duplicated by key.
func DifferenceBy[T any, K comparable](fn func(T) K, s []T, others ...[]T) []T {
	exclude := make(map[K]struct{})
	for _, o := range others {
		for _, v := range o {
			exclude[fn(v)] = struct{}{}
		}
	}
	seen := make(map[K]struct{})
	r := make([]T, 0, len(s))
	for _, v := range s {
		k := fn(v)
		if _, bad := exclude[k]; bad {
			continue
		}
		if _, dup := seen[k]; dup {
			continue
		}
		seen[k] = struct{}{}
		r = append(r, v)
	}
	return r
}

// DifferenceWith returns the elements of s that are not considered equal, by the
// eq comparator, to any element of the others slices.
func DifferenceWith[T any](eq func(a, b T) bool, s []T, others ...[]T) []T {
	r := make([]T, 0, len(s))
	for _, v := range s {
		excluded := false
		for _, o := range others {
			for _, w := range o {
				if eq(v, w) {
					excluded = true
					break
				}
			}
			if excluded {
				break
			}
		}
		if !excluded {
			r = append(r, v)
		}
	}
	return r
}

// IntersectionBy returns the elements present (by key) in every provided slice,
// ordered by first appearance in the first slice and de-duplicated by key.
func IntersectionBy[T any, K comparable](fn func(T) K, slices ...[]T) []T {
	if len(slices) == 0 {
		return nil
	}
	sets := make([]map[K]struct{}, len(slices))
	for i, s := range slices {
		m := make(map[K]struct{}, len(s))
		for _, v := range s {
			m[fn(v)] = struct{}{}
		}
		sets[i] = m
	}
	seen := make(map[K]struct{})
	var r []T
	for _, v := range slices[0] {
		k := fn(v)
		if _, dup := seen[k]; dup {
			continue
		}
		inAll := true
		for i := 1; i < len(sets); i++ {
			if _, ok := sets[i][k]; !ok {
				inAll = false
				break
			}
		}
		if inAll {
			seen[k] = struct{}{}
			r = append(r, v)
		}
	}
	return r
}

// IntersectionWith returns the elements of the first slice that are equal, by
// the eq comparator, to some element of every other slice.
func IntersectionWith[T any](eq func(a, b T) bool, slices ...[]T) []T {
	if len(slices) == 0 {
		return nil
	}
	var r []T
	for _, v := range slices[0] {
		inAll := true
		for i := 1; i < len(slices); i++ {
			found := false
			for _, w := range slices[i] {
				if eq(v, w) {
					found = true
					break
				}
			}
			if !found {
				inAll = false
				break
			}
		}
		if inAll {
			r = append(r, v)
		}
	}
	return r
}

// UnionBy returns the unique elements (by key) from all provided slices,
// preserving order of first appearance.
func UnionBy[T any, K comparable](fn func(T) K, slices ...[]T) []T {
	seen := make(map[K]struct{})
	var r []T
	for _, s := range slices {
		for _, v := range s {
			k := fn(v)
			if _, ok := seen[k]; !ok {
				seen[k] = struct{}{}
				r = append(r, v)
			}
		}
	}
	return r
}

// UnionWith returns the unique elements from all provided slices using the eq
// comparator to detect duplicates, preserving order of first appearance.
func UnionWith[T any](eq func(a, b T) bool, slices ...[]T) []T {
	var r []T
	for _, s := range slices {
		for _, v := range s {
			dup := false
			for _, w := range r {
				if eq(v, w) {
					dup = true
					break
				}
			}
			if !dup {
				r = append(r, v)
			}
		}
	}
	return r
}

// Xor returns the symmetric difference of the provided slices: the unique values
// that appear in exactly one of the slices.
func Xor[T comparable](slices ...[]T) []T {
	return XorBy(func(v T) T { return v }, slices...)
}

// XorBy is like Xor but element identity is determined by the key produced by
// fn.
func XorBy[T any, K comparable](fn func(T) K, slices ...[]T) []T {
	count := make(map[K]int)
	for _, s := range slices {
		local := make(map[K]struct{})
		for _, v := range s {
			k := fn(v)
			if _, ok := local[k]; !ok {
				local[k] = struct{}{}
				count[k]++
			}
		}
	}
	var r []T
	seen := make(map[K]struct{})
	for _, s := range slices {
		for _, v := range s {
			k := fn(v)
			if count[k] == 1 {
				if _, ok := seen[k]; !ok {
					seen[k] = struct{}{}
					r = append(r, v)
				}
			}
		}
	}
	return r
}

// XorWith is like Xor but element identity is determined by the eq comparator.
func XorWith[T any](eq func(a, b T) bool, slices ...[]T) []T {
	var all []T
	for _, s := range slices {
		all = append(all, s...)
	}
	var r []T
	for _, v := range all {
		count := 0
		for _, w := range all {
			if eq(v, w) {
				count++
			}
		}
		if count == 1 {
			r = append(r, v)
		}
	}
	return r
}

// SortedUniq removes duplicate values from a sorted slice, keeping one of each
// run of equal elements. The input is not mutated.
func SortedUniq[T comparable](s []T) []T {
	r := make([]T, 0, len(s))
	for i, v := range s {
		if i == 0 || v != s[i-1] {
			r = append(r, v)
		}
	}
	return r
}

// SortedUniqBy removes duplicates from a sorted slice, comparing elements by the
// key produced by fn. The input is not mutated.
func SortedUniqBy[T any, K comparable](s []T, fn func(T) K) []T {
	r := make([]T, 0, len(s))
	var prev K
	for i, v := range s {
		k := fn(v)
		if i == 0 || k != prev {
			r = append(r, v)
			prev = k
		}
	}
	return r
}

// PullAll returns a new slice with every occurrence of the given values removed.
// The input is not mutated (unlike lodash's mutating _.pullAll).
func PullAll[T comparable](s []T, values []T) []T {
	remove := make(map[T]struct{}, len(values))
	for _, v := range values {
		remove[v] = struct{}{}
	}
	r := make([]T, 0, len(s))
	for _, v := range s {
		if _, ok := remove[v]; !ok {
			r = append(r, v)
		}
	}
	return r
}

// PullAllBy is like PullAll but elements are matched by the key produced by fn.
func PullAllBy[T any, K comparable](s []T, values []T, fn func(T) K) []T {
	remove := make(map[K]struct{}, len(values))
	for _, v := range values {
		remove[fn(v)] = struct{}{}
	}
	r := make([]T, 0, len(s))
	for _, v := range s {
		if _, ok := remove[fn(v)]; !ok {
			r = append(r, v)
		}
	}
	return r
}

// Pull returns a new slice with every occurrence of the given values removed.
// It is the variadic form of PullAll.
func Pull[T comparable](s []T, values ...T) []T {
	return PullAll(s, values)
}

// PullAt returns a new slice with the elements at the given indices removed.
// Negative and out-of-range indices are ignored. The input is not mutated.
func PullAt[T any](s []T, indexes ...int) []T {
	drop := make(map[int]struct{}, len(indexes))
	for _, i := range indexes {
		if i < 0 {
			i += len(s)
		}
		drop[i] = struct{}{}
	}
	r := make([]T, 0, len(s))
	for i, v := range s {
		if _, ok := drop[i]; !ok {
			r = append(r, v)
		}
	}
	return r
}

// Remove returns a new slice containing only the elements for which pred returns
// false, i.e. the elements that lodash's mutating _.remove would leave behind.
// The input is not mutated.
func Remove[T any](s []T, pred func(T) bool) []T {
	r := make([]T, 0, len(s))
	for _, v := range s {
		if !pred(v) {
			r = append(r, v)
		}
	}
	return r
}

// FlatMap maps each element to a slice via fn and flattens the results by one
// level.
func FlatMap[T, R any](s []T, fn func(T) []R) []R {
	var r []R
	for _, v := range s {
		r = append(r, fn(v)...)
	}
	return r
}

// FlattenDepth flattens a nested structure up to depth levels deep. A depth of 0
// returns a copy of the input unchanged; higher depths flatten further.
func FlattenDepth[T any](s []any, depth int) []T {
	var r []T
	var walk func(items []any, d int)
	walk = func(items []any, d int) {
		for _, item := range items {
			if sub, ok := item.([]any); ok && d > 0 {
				walk(sub, d-1)
				continue
			}
			if v, ok := item.(T); ok {
				r = append(r, v)
			}
		}
	}
	walk(s, depth)
	return r
}

// TakeWhile returns the leading run of elements for which pred returns true.
func TakeWhile[T any](s []T, pred func(T) bool) []T {
	r := make([]T, 0, len(s))
	for _, v := range s {
		if !pred(v) {
			break
		}
		r = append(r, v)
	}
	return r
}

// TakeRightWhile returns the trailing run of elements for which pred returns
// true, in their original order.
func TakeRightWhile[T any](s []T, pred func(T) bool) []T {
	i := len(s)
	for i > 0 && pred(s[i-1]) {
		i--
	}
	r := make([]T, len(s)-i)
	copy(r, s[i:])
	return r
}

// DropWhile drops the leading run of elements for which pred returns true and
// returns the rest.
func DropWhile[T any](s []T, pred func(T) bool) []T {
	i := 0
	for i < len(s) && pred(s[i]) {
		i++
	}
	r := make([]T, len(s)-i)
	copy(r, s[i:])
	return r
}

// DropRightWhile drops the trailing run of elements for which pred returns true
// and returns the rest.
func DropRightWhile[T any](s []T, pred func(T) bool) []T {
	i := len(s)
	for i > 0 && pred(s[i-1]) {
		i--
	}
	r := make([]T, i)
	copy(r, s[:i])
	return r
}

// ZipWith combines two slices element-wise using fn, producing a slice whose
// length equals the shorter input.
func ZipWith[A, B, R any](a []A, b []B, fn func(A, B) R) []R {
	n := len(a)
	if len(b) < n {
		n = len(b)
	}
	r := make([]R, n)
	for i := 0; i < n; i++ {
		r[i] = fn(a[i], b[i])
	}
	return r
}

// UnzipWith splits a slice of pairs into two slices, then combines them
// element-wise using fn.
func UnzipWith[A, B, R any](pairs []Pair[A, B], fn func(A, B) R) []R {
	r := make([]R, len(pairs))
	for i, p := range pairs {
		r[i] = fn(p.First, p.Second)
	}
	return r
}

// SortedIndex returns the lowest index at which value can be inserted into the
// sorted slice s to keep it sorted.
func SortedIndex[T cmp.Ordered](s []T, value T) int {
	return sort.Search(len(s), func(i int) bool { return s[i] >= value })
}

// SortedLastIndex returns the highest index at which value can be inserted into
// the sorted slice s to keep it sorted.
func SortedLastIndex[T cmp.Ordered](s []T, value T) int {
	return sort.Search(len(s), func(i int) bool { return s[i] > value })
}

// SortedIndexBy returns the lowest insertion index for value in s, comparing
// elements by the key produced by fn.
func SortedIndexBy[T any, K cmp.Ordered](s []T, value T, fn func(T) K) int {
	target := fn(value)
	return sort.Search(len(s), func(i int) bool { return fn(s[i]) >= target })
}

// SortedLastIndexBy returns the highest insertion index for value in s,
// comparing elements by the key produced by fn.
func SortedLastIndexBy[T any, K cmp.Ordered](s []T, value T, fn func(T) K) int {
	target := fn(value)
	return sort.Search(len(s), func(i int) bool { return fn(s[i]) > target })
}

// SortedIndexOf returns the index of the first occurrence of value in the sorted
// slice s using binary search, or -1 if it is absent.
func SortedIndexOf[T cmp.Ordered](s []T, value T) int {
	i := SortedIndex(s, value)
	if i < len(s) && s[i] == value {
		return i
	}
	return -1
}

// SortedLastIndexOf returns the index of the last occurrence of value in the
// sorted slice s using binary search, or -1 if it is absent.
func SortedLastIndexOf[T cmp.Ordered](s []T, value T) int {
	i := SortedLastIndex(s, value) - 1
	if i >= 0 && i < len(s) && s[i] == value {
		return i
	}
	return -1
}

// IndexOfFrom returns the index of the first occurrence of value at or after
// index from, or -1 if it is absent. A negative from counts from the end.
func IndexOfFrom[T comparable](s []T, value T, from int) int {
	if from < 0 {
		from += len(s)
	}
	if from < 0 {
		from = 0
	}
	for i := from; i < len(s); i++ {
		if s[i] == value {
			return i
		}
	}
	return -1
}

// LastIndexOfFrom returns the index of the last occurrence of value at or before
// index from, or -1 if it is absent. A negative from counts from the end.
func LastIndexOfFrom[T comparable](s []T, value T, from int) int {
	if from < 0 {
		from += len(s)
	}
	if from >= len(s) {
		from = len(s) - 1
	}
	for i := from; i >= 0; i-- {
		if s[i] == value {
			return i
		}
	}
	return -1
}

// None reports whether pred returns false for every element. It is the negation
// of Some and returns true for an empty slice.
func None[T any](s []T, pred func(T) bool) bool {
	for _, v := range s {
		if pred(v) {
			return false
		}
	}
	return true
}

// ForEachRight invokes fn once for each element from right to left.
func ForEachRight[T any](s []T, fn func(T)) {
	for i := len(s) - 1; i >= 0; i-- {
		fn(s[i])
	}
}
