package lodash

import "math/rand"

// SliceSeq is a fluent wrapper around a slice that enables lodash-style method
// chaining for transforms that preserve the element type. Because Go methods
// cannot introduce new type parameters, element-type-changing operations (such
// as mapping []T to []R) remain package-level functions; SliceSeq covers the
// large family of slice operations that keep the same element type. Call Value
// to unwrap the final slice.
type SliceSeq[T any] struct {
	items []T
}

// ChainSlice wraps s in a SliceSeq so that same-type slice transforms can be
// applied fluently.
func ChainSlice[T any](s []T) SliceSeq[T] {
	return SliceSeq[T]{items: s}
}

// Value returns the wrapped slice, terminating the chain.
func (s SliceSeq[T]) Value() []T {
	return s.items
}

// Filter keeps only the elements for which pred returns true.
func (s SliceSeq[T]) Filter(pred func(T) bool) SliceSeq[T] {
	return SliceSeq[T]{items: Filter(s.items, pred)}
}

// Reject removes the elements for which pred returns true.
func (s SliceSeq[T]) Reject(pred func(T) bool) SliceSeq[T] {
	return SliceSeq[T]{items: Reject(s.items, pred)}
}

// Reverse reverses the order of the wrapped elements.
func (s SliceSeq[T]) Reverse() SliceSeq[T] {
	return SliceSeq[T]{items: Reverse(s.items)}
}

// Take keeps the first n elements.
func (s SliceSeq[T]) Take(n int) SliceSeq[T] {
	return SliceSeq[T]{items: Take(s.items, n)}
}

// TakeRight keeps the last n elements.
func (s SliceSeq[T]) TakeRight(n int) SliceSeq[T] {
	return SliceSeq[T]{items: TakeRight(s.items, n)}
}

// TakeWhile keeps leading elements while pred returns true.
func (s SliceSeq[T]) TakeWhile(pred func(T) bool) SliceSeq[T] {
	return SliceSeq[T]{items: TakeWhile(s.items, pred)}
}

// Drop removes the first n elements.
func (s SliceSeq[T]) Drop(n int) SliceSeq[T] {
	return SliceSeq[T]{items: Drop(s.items, n)}
}

// DropRight removes the last n elements.
func (s SliceSeq[T]) DropRight(n int) SliceSeq[T] {
	return SliceSeq[T]{items: DropRight(s.items, n)}
}

// DropWhile removes leading elements while pred returns true.
func (s SliceSeq[T]) DropWhile(pred func(T) bool) SliceSeq[T] {
	return SliceSeq[T]{items: DropWhile(s.items, pred)}
}

// Tail removes the first element.
func (s SliceSeq[T]) Tail() SliceSeq[T] {
	return SliceSeq[T]{items: Tail(s.items)}
}

// Initial removes the last element.
func (s SliceSeq[T]) Initial() SliceSeq[T] {
	return SliceSeq[T]{items: Initial(s.items)}
}

// Slice keeps the half-open index range [start, end).
func (s SliceSeq[T]) Slice(start, end int) SliceSeq[T] {
	return SliceSeq[T]{items: Slice(s.items, start, end)}
}

// Concat appends the elements of others after the wrapped elements.
func (s SliceSeq[T]) Concat(others ...[]T) SliceSeq[T] {
	all := append([][]T{s.items}, others...)
	return SliceSeq[T]{items: Concat(all...)}
}

// ForEach invokes fn for each element and returns the SliceSeq unchanged,
// allowing side effects mid-chain.
func (s SliceSeq[T]) ForEach(fn func(T)) SliceSeq[T] {
	ForEach(s.items, fn)
	return s
}

// Tap invokes fn with the wrapped slice for its side effects and returns the
// SliceSeq unchanged.
func (s SliceSeq[T]) Tap(fn func([]T)) SliceSeq[T] {
	fn(s.items)
	return s
}

// Shuffle returns a randomly reordered copy of the wrapped slice using rng, so
// results are deterministic when rng is seeded.
func (s SliceSeq[T]) Shuffle(rng *rand.Rand) SliceSeq[T] {
	return SliceSeq[T]{items: Shuffle(s.items, rng)}
}

// Head returns the first element and true, or the zero value and false when the
// slice is empty. It terminates the chain.
func (s SliceSeq[T]) Head() (T, bool) {
	return Head(s.items)
}

// Last returns the final element and true, or the zero value and false when the
// slice is empty. It terminates the chain.
func (s SliceSeq[T]) Last() (T, bool) {
	return Last(s.items)
}

// Nth returns the element at index n (negative counts from the end) and true,
// or the zero value and false when out of range. It terminates the chain.
func (s SliceSeq[T]) Nth(n int) (T, bool) {
	return Nth(s.items, n)
}

// Sample returns a random element and true, or the zero value and false when
// the slice is empty, drawing from rng. It terminates the chain.
func (s SliceSeq[T]) Sample(rng *rand.Rand) (T, bool) {
	return Sample(s.items, rng)
}

// Size returns the number of wrapped elements.
func (s SliceSeq[T]) Size() int {
	return len(s.items)
}

// IsEmpty reports whether the wrapped slice has no elements.
func (s SliceSeq[T]) IsEmpty() bool {
	return len(s.items) == 0
}

// Join concatenates the string forms of the wrapped elements separated by sep.
// It terminates the chain.
func (s SliceSeq[T]) Join(sep string) string {
	return Join(s.items, sep)
}

// Chunk splits the wrapped slice into groups of at most size elements. It
// returns the groups directly, terminating the chain, because Go's type system
// disallows a method returning a SliceSeq re-parameterised over []T.
func (s SliceSeq[T]) Chunk(size int) [][]T {
	return Chunk(s.items, size)
}
