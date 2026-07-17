package lodash

// Seq is a lightweight value wrapper that enables lodash-style method chaining.
// A Seq carries a single value through a pipeline of transforms; call Value to
// unwrap the final result. Because Go methods cannot introduce new type
// parameters, transforms that change the value's type are provided as the
// package-level functions Thru and Tap.
type Seq[T any] struct {
	value T
}

// Chain wraps value in a Seq so that transforms can be applied fluently.
func Chain[T any](value T) Seq[T] {
	return Seq[T]{value: value}
}

// Value returns the wrapped value, terminating the chain.
func (s Seq[T]) Value() T {
	return s.value
}

// Thru applies a same-typed transform to the wrapped value and returns a new
// Seq holding the result. Use the package-level Thru function to transform the
// value into a different type.
func (s Seq[T]) Thru(fn func(T) T) Seq[T] {
	return Seq[T]{value: fn(s.value)}
}

// Tap invokes fn with the wrapped value for its side effects and returns the
// Seq unchanged, allowing you to observe an intermediate value mid-chain.
func (s Seq[T]) Tap(fn func(T)) Seq[T] {
	fn(s.value)
	return s
}

// Thru applies fn to the value wrapped by s and returns a new Seq holding the
// (possibly differently typed) result.
func Thru[T, R any](s Seq[T], fn func(T) R) Seq[R] {
	return Seq[R]{value: fn(s.value)}
}

// Tap invokes fn with the value wrapped by s for its side effects and returns
// the Seq unchanged.
func Tap[T any](s Seq[T], fn func(T)) Seq[T] {
	fn(s.value)
	return s
}
