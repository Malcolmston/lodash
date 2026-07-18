package lodash

import "time"

// CurryRight transforms a two-argument function into a sequence of unary
// functions that supply the arguments from right to left: the returned function
// takes the last argument first. It is the right-to-left analogue of Curry.
func CurryRight[A, B, R any](fn func(A, B) R) func(B) func(A) R {
	return func(b B) func(A) R {
		return func(a A) R {
			return fn(a, b)
		}
	}
}

// CurryRight3 transforms a three-argument function into a sequence of unary
// functions that supply the arguments from right to left. It is the
// right-to-left analogue of Curry3.
func CurryRight3[A, B, C, R any](fn func(A, B, C) R) func(C) func(B) func(A) R {
	return func(c C) func(B) func(A) R {
		return func(b B) func(A) R {
			return func(a A) R {
				return fn(a, b, c)
			}
		}
	}
}

// Curry4 transforms a four-argument function into a sequence of unary functions
// that supply the arguments from left to right. It extends the Curry/Curry3
// family to arity four.
func Curry4[A, B, C, D, R any](fn func(A, B, C, D) R) func(A) func(B) func(C) func(D) R {
	return func(a A) func(B) func(C) func(D) R {
		return func(b B) func(C) func(D) R {
			return func(c C) func(D) R {
				return func(d D) R {
					return fn(a, b, c, d)
				}
			}
		}
	}
}

// Rest wraps fn, which takes a single slice argument, so that it can be called
// with variadic arguments instead. It is the inverse of Spread and mirrors
// lodash's rest with the collected arguments passed as one slice.
func Rest[T, R any](fn func([]T) R) func(...T) R {
	return func(args ...T) R {
		return fn(args)
	}
}

// Delay schedules fn to run after wait has elapsed and returns the underlying
// *time.Timer. Call the timer's Stop method to cancel the pending invocation
// before it fires. It mirrors lodash's delay.
func Delay(fn func(), wait time.Duration) *time.Timer {
	return time.AfterFunc(wait, fn)
}

// Defer schedules fn to run as soon as possible on a separate goroutine and
// returns the underlying *time.Timer, whose Stop method cancels the pending
// invocation. It mirrors lodash's defer.
func Defer(fn func()) *time.Timer {
	return time.AfterFunc(0, fn)
}
