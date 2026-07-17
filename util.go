package lodash

import (
	"strconv"
	"sync/atomic"
)

// Identity returns its argument unchanged. It is useful as a default transform.
func Identity[T any](value T) T { return value }

// Constant returns a function that always returns value, ignoring any later
// changes to the surrounding state.
func Constant[T any](value T) func() T {
	return func() T { return value }
}

// Noop does nothing. It is handy as a default callback.
func Noop() {}

// Times invokes fn n times, collecting the results into a slice. The call index
// (0-based) is passed to fn. A non-positive n yields an empty slice.
func Times[R any](n int, fn func(int) R) []R {
	if n < 0 {
		n = 0
	}
	r := make([]R, n)
	for i := 0; i < n; i++ {
		r[i] = fn(i)
	}
	return r
}

// StubArray returns a new empty slice of T on every call.
func StubArray[T any]() []T { return []T{} }

// StubObject returns a new empty map[string]any on every call.
func StubObject() map[string]any { return map[string]any{} }

// StubString returns the empty string.
func StubString() string { return "" }

// StubTrue always returns true.
func StubTrue() bool { return true }

// StubFalse always returns false.
func StubFalse() bool { return false }

var uniqueIDCounter uint64

// UniqueID returns a process-unique identifier string composed of the optional
// prefix followed by a monotonically increasing counter. It is safe for
// concurrent use.
func UniqueID(prefix string) string {
	n := atomic.AddUint64(&uniqueIDCounter, 1)
	return prefix + strconv.FormatUint(n, 10)
}

// Property returns a function that reads the value at path from a nested
// map[string]any using Get semantics (dotted and bracketed paths).
func Property(path string) func(map[string]any) (any, bool) {
	return func(m map[string]any) (any, bool) {
		return Get(m, path)
	}
}

// PropertyOf returns a function that reads the value at a given path from the
// bound object using Get semantics.
func PropertyOf(object map[string]any) func(string) (any, bool) {
	return func(path string) (any, bool) {
		return Get(object, path)
	}
}

// Matches returns a predicate that reports whether its argument contains all of
// the key/value pairs in source, using IsMatch semantics.
func Matches(source map[string]any) func(map[string]any) bool {
	return func(object map[string]any) bool {
		return IsMatch(object, source)
	}
}

// Attempt invokes fn and recovers any panic, returning the panic value as an
// error. On success it returns the function's result and a nil error.
func Attempt[R any](fn func() R) (result R, err error) {
	defer func() {
		if r := recover(); r != nil {
			if e, ok := r.(error); ok {
				err = e
			} else {
				err = &attemptError{value: r}
			}
		}
	}()
	return fn(), nil
}

type attemptError struct{ value any }

func (e *attemptError) Error() string {
	return ToString(e.value)
}
