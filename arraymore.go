package lodash

// UniqWith returns a new slice with duplicate elements removed, using eq to
// decide whether two elements are equal. The first occurrence of each distinct
// element is kept, in order. It is the comparator-based analogue of Uniq for
// element types that are not comparable with ==.
func UniqWith[T any](s []T, eq func(a, b T) bool) []T {
	r := make([]T, 0, len(s))
	for _, v := range s {
		dup := false
		for _, kept := range r {
			if eq(kept, v) {
				dup = true
				break
			}
		}
		if !dup {
			r = append(r, v)
		}
	}
	return r
}

// PullAllWith returns a new slice with every element of s that is equal to any
// element of values removed, using eq for comparison. The input slice is not
// modified. It is the comparator-based analogue of PullAll.
func PullAllWith[T any](s []T, values []T, eq func(a, b T) bool) []T {
	r := make([]T, 0, len(s))
	for _, v := range s {
		remove := false
		for _, target := range values {
			if eq(v, target) {
				remove = true
				break
			}
		}
		if !remove {
			r = append(r, v)
		}
	}
	return r
}

// ZipObjectDeep builds a nested map[string]any by treating each element of
// paths as a property path (see ToPath) and assigning the corresponding element
// of values at that path. Paths without a matching value receive nil. It
// mirrors lodash's zipObjectDeep.
func ZipObjectDeep(paths []string, values []any) map[string]any {
	r := map[string]any{}
	for i, p := range paths {
		var v any
		if i < len(values) {
			v = values[i]
		}
		r = Set(r, p, v)
	}
	return r
}

// FlatMapDeep maps every element of s through fn and then recursively flattens
// the collected results to any depth, returning a flat []T. Each result of fn
// may itself contain nested []any values, which are flattened as well. It
// mirrors lodash's flatMapDeep.
func FlatMapDeep[T any](s []any, fn func(any) []any) []T {
	mapped := make([]any, 0, len(s))
	for _, v := range s {
		for _, out := range fn(v) {
			mapped = append(mapped, out)
		}
	}
	return FlattenDeep[T](mapped)
}

// FlatMapDepth maps every element of s through fn and then flattens the
// collected results by up to depth levels, returning a flat []T. A depth of 0
// performs no flattening. It mirrors lodash's flatMapDepth.
func FlatMapDepth[T any](s []any, fn func(any) []any, depth int) []T {
	mapped := make([]any, 0, len(s))
	for _, v := range s {
		for _, out := range fn(v) {
			mapped = append(mapped, out)
		}
	}
	return FlattenDepth[T](mapped, depth)
}

// FillRange returns a copy of s with the elements in the half-open index range
// [start, end) replaced by value. Negative indexes count from the end of the
// slice, and the range is clamped to the bounds of s. The input slice is not
// modified. It mirrors lodash's fill with explicit start and end arguments.
func FillRange[T any](s []T, value T, start, end int) []T {
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
	r := make([]T, n)
	copy(r, s)
	for i := start; i < end; i++ {
		r[i] = value
	}
	return r
}
