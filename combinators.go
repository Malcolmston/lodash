package lodash

// Negate returns a predicate that reports the logical negation of pred.
func Negate[T any](pred func(T) bool) func(T) bool {
	return func(v T) bool { return !pred(v) }
}

// Curry converts a two-argument function into a sequence of unary functions, so
// that fn(a, b) can be called as Curry(fn)(a)(b).
func Curry[A, B, R any](fn func(A, B) R) func(A) func(B) R {
	return func(a A) func(B) R {
		return func(b B) R {
			return fn(a, b)
		}
	}
}

// Curry3 converts a three-argument function into a sequence of unary functions,
// so that fn(a, b, c) can be called as Curry3(fn)(a)(b)(c).
func Curry3[A, B, C, R any](fn func(A, B, C) R) func(A) func(B) func(C) R {
	return func(a A) func(B) func(C) R {
		return func(b B) func(C) R {
			return func(c C) R {
				return fn(a, b, c)
			}
		}
	}
}

// Partial binds the first argument of a two-argument function, returning a unary
// function of the remaining argument.
func Partial[A, B, R any](fn func(A, B) R, a A) func(B) R {
	return func(b B) R { return fn(a, b) }
}

// PartialRight binds the last argument of a two-argument function, returning a
// unary function of the remaining (first) argument.
func PartialRight[A, B, R any](fn func(A, B) R, b B) func(A) R {
	return func(a A) R { return fn(a, b) }
}

// Flip returns a function that invokes fn with its two arguments swapped.
func Flip[A, B, R any](fn func(A, B) R) func(B, A) R {
	return func(b B, a A) R { return fn(a, b) }
}

// Flow composes the given same-typed transforms into a single function that
// applies them left to right: the output of each becomes the input of the next.
// Flow with no functions returns the identity function.
func Flow[T any](fns ...func(T) T) func(T) T {
	return func(v T) T {
		for _, fn := range fns {
			v = fn(v)
		}
		return v
	}
}

// FlowRight composes the given same-typed transforms into a single function that
// applies them right to left (mathematical composition).
func FlowRight[T any](fns ...func(T) T) func(T) T {
	return func(v T) T {
		for i := len(fns) - 1; i >= 0; i-- {
			v = fns[i](v)
		}
		return v
	}
}

// CondPair couples a predicate with the transform to run when it matches, for
// use with Cond.
type CondPair[T, R any] struct {
	// Pred reports whether Result should be applied to the input.
	Pred func(T) bool
	// Result produces the output when Pred matches.
	Result func(T) R
}

// Cond returns a function that runs the input through the first CondPair whose
// predicate matches, returning that pair's transformed result. If none match it
// returns the zero value of R and false.
func Cond[T, R any](pairs ...CondPair[T, R]) func(T) (R, bool) {
	return func(v T) (R, bool) {
		for _, p := range pairs {
			if p.Pred(v) {
				return p.Result(v), true
			}
		}
		var zero R
		return zero, false
	}
}

// Over returns a function that runs its input through every transform in fns and
// collects the results into a slice, preserving order.
func Over[T, R any](fns ...func(T) R) func(T) []R {
	return func(v T) []R {
		r := make([]R, len(fns))
		for i, fn := range fns {
			r[i] = fn(v)
		}
		return r
	}
}

// OverEvery returns a predicate that reports whether the input satisfies every
// predicate in preds. With no predicates it always returns true.
func OverEvery[T any](preds ...func(T) bool) func(T) bool {
	return func(v T) bool {
		for _, p := range preds {
			if !p(v) {
				return false
			}
		}
		return true
	}
}

// OverSome returns a predicate that reports whether the input satisfies at least
// one predicate in preds. With no predicates it always returns false.
func OverSome[T any](preds ...func(T) bool) func(T) bool {
	return func(v T) bool {
		for _, p := range preds {
			if p(v) {
				return true
			}
		}
		return false
	}
}

// OverArgs returns a function that transforms its variadic arguments with the
// corresponding transform in transforms (by position) before passing them to
// fn. Arguments without a matching transform are passed through unchanged.
func OverArgs[T, R any](fn func(...T) R, transforms ...func(T) T) func(...T) R {
	return func(args ...T) R {
		out := make([]T, len(args))
		for i, a := range args {
			if i < len(transforms) {
				out[i] = transforms[i](a)
			} else {
				out[i] = a
			}
		}
		return fn(out...)
	}
}

// Ary caps a variadic function so that it receives at most n arguments; any
// extra arguments are dropped. A negative n is treated as zero.
func Ary[T, R any](fn func(...T) R, n int) func(...T) R {
	if n < 0 {
		n = 0
	}
	return func(args ...T) R {
		if len(args) > n {
			args = args[:n]
		}
		return fn(args...)
	}
}

// Unary caps a variadic function so that it receives at most a single argument.
func Unary[T, R any](fn func(...T) R) func(T) R {
	return func(v T) R { return fn(v) }
}

// Rearg returns a function that reorders its variadic arguments according to
// indexes before calling fn: the i-th argument passed to fn is the argument at
// position indexes[i] of the call. Out-of-range indexes contribute the zero
// value.
func Rearg[T, R any](fn func(...T) R, indexes ...int) func(...T) R {
	return func(args ...T) R {
		out := make([]T, len(indexes))
		for i, idx := range indexes {
			if idx >= 0 && idx < len(args) {
				out[i] = args[idx]
			}
		}
		return fn(out...)
	}
}

// Spread converts a variadic function into one that takes its arguments as a
// single slice.
func Spread[T, R any](fn func(...T) R) func([]T) R {
	return func(args []T) R { return fn(args...) }
}

// Wrap returns a function that calls wrapper with value followed by the caller's
// argument, letting wrapper decorate or intercept the bound value.
func Wrap[T, A, R any](value T, wrapper func(T, A) R) func(A) R {
	return func(a A) R { return wrapper(value, a) }
}

// Compose composes the given same-typed transforms right to left. It is an alias
// for FlowRight, matching the common functional-programming name.
func Compose[T any](fns ...func(T) T) func(T) T {
	return FlowRight(fns...)
}

// NthArg returns a function that returns the argument at position n from its
// variadic arguments. A negative n counts from the end. When n is out of range
// the zero value is returned.
func NthArg[T any](n int) func(...T) T {
	return func(args ...T) T {
		i := n
		if i < 0 {
			i += len(args)
		}
		if i < 0 || i >= len(args) {
			var zero T
			return zero
		}
		return args[i]
	}
}
