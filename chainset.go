package lodash

// SetSeq is a fluent wrapper around a slice of comparable elements. It
// complements SliceSeq with the set-oriented operations that require the
// element type to be comparable (deduplication, membership and the algebraic
// set combinators). Call Value to unwrap the final slice.
type SetSeq[T comparable] struct {
	items []T
}

// ChainSet wraps s in a SetSeq so that set-oriented transforms can be applied
// fluently.
func ChainSet[T comparable](s []T) SetSeq[T] {
	return SetSeq[T]{items: s}
}

// Value returns the wrapped slice, terminating the chain.
func (s SetSeq[T]) Value() []T {
	return s.items
}

// Uniq removes duplicate elements, keeping the first occurrence of each in
// order.
func (s SetSeq[T]) Uniq() SetSeq[T] {
	return SetSeq[T]{items: Uniq(s.items)}
}

// Compact removes zero-valued elements (for example 0, "" or false).
func (s SetSeq[T]) Compact() SetSeq[T] {
	return SetSeq[T]{items: Compact(s.items)}
}

// Without removes every element equal to any of values.
func (s SetSeq[T]) Without(values ...T) SetSeq[T] {
	return SetSeq[T]{items: Without(s.items, values...)}
}

// Filter keeps only the elements for which pred returns true.
func (s SetSeq[T]) Filter(pred func(T) bool) SetSeq[T] {
	return SetSeq[T]{items: Filter(s.items, pred)}
}

// Reverse reverses the order of the wrapped elements.
func (s SetSeq[T]) Reverse() SetSeq[T] {
	return SetSeq[T]{items: Reverse(s.items)}
}

// Union returns the unique union of the wrapped slice with others, preserving
// first-seen order.
func (s SetSeq[T]) Union(others ...[]T) SetSeq[T] {
	all := append([][]T{s.items}, others...)
	return SetSeq[T]{items: Union(all...)}
}

// Intersection keeps only the elements that also appear in every one of others.
func (s SetSeq[T]) Intersection(others ...[]T) SetSeq[T] {
	all := append([][]T{s.items}, others...)
	return SetSeq[T]{items: Intersection(all...)}
}

// Difference removes the elements that appear in any of others.
func (s SetSeq[T]) Difference(others ...[]T) SetSeq[T] {
	return SetSeq[T]{items: Difference(s.items, others...)}
}

// Includes reports whether value is present in the wrapped slice. It terminates
// the chain.
func (s SetSeq[T]) Includes(value T) bool {
	return Includes(s.items, value)
}

// IndexOf returns the index of the first occurrence of value, or -1 when it is
// absent. It terminates the chain.
func (s SetSeq[T]) IndexOf(value T) int {
	return IndexOf(s.items, value)
}

// Size returns the number of wrapped elements.
func (s SetSeq[T]) Size() int {
	return len(s.items)
}
