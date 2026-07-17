package lodash

import "cmp"

// Number is a constraint permitting any built-in numeric type.
type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64
}

// Sum returns the sum of all elements. It returns the zero value for an empty
// slice.
func Sum[T Number](s []T) T {
	var total T
	for _, v := range s {
		total += v
	}
	return total
}

// SumBy returns the sum of applying fn to each element.
func SumBy[T any, R Number](s []T, fn func(T) R) R {
	var total R
	for _, v := range s {
		total += fn(v)
	}
	return total
}

// Mean returns the arithmetic mean of the elements as a float64. It returns 0
// for an empty slice.
func Mean[T Number](s []T) float64 {
	if len(s) == 0 {
		return 0
	}
	return float64(Sum(s)) / float64(len(s))
}

// MeanBy returns the mean of applying fn to each element as a float64.
func MeanBy[T any, R Number](s []T, fn func(T) R) float64 {
	if len(s) == 0 {
		return 0
	}
	return float64(SumBy(s, fn)) / float64(len(s))
}

// Min returns the smallest element and true, or the zero value and false when
// the slice is empty.
func Min[T cmp.Ordered](s []T) (T, bool) {
	if len(s) == 0 {
		var zero T
		return zero, false
	}
	m := s[0]
	for _, v := range s[1:] {
		if v < m {
			m = v
		}
	}
	return m, true
}

// MinBy returns the element whose key (as produced by fn) is smallest.
func MinBy[T any, K cmp.Ordered](s []T, fn func(T) K) (T, bool) {
	if len(s) == 0 {
		var zero T
		return zero, false
	}
	best := s[0]
	bestKey := fn(best)
	for _, v := range s[1:] {
		k := fn(v)
		if k < bestKey {
			best, bestKey = v, k
		}
	}
	return best, true
}

// Max returns the largest element and true, or the zero value and false when
// the slice is empty.
func Max[T cmp.Ordered](s []T) (T, bool) {
	if len(s) == 0 {
		var zero T
		return zero, false
	}
	m := s[0]
	for _, v := range s[1:] {
		if v > m {
			m = v
		}
	}
	return m, true
}

// MaxBy returns the element whose key (as produced by fn) is largest.
func MaxBy[T any, K cmp.Ordered](s []T, fn func(T) K) (T, bool) {
	if len(s) == 0 {
		var zero T
		return zero, false
	}
	best := s[0]
	bestKey := fn(best)
	for _, v := range s[1:] {
		k := fn(v)
		if k > bestKey {
			best, bestKey = v, k
		}
	}
	return best, true
}

// Clamp constrains value to lie within the inclusive range [lower, upper]. If
// lower is greater than upper the bounds are swapped.
func Clamp[T cmp.Ordered](value, lower, upper T) T {
	if lower > upper {
		lower, upper = upper, lower
	}
	if value < lower {
		return lower
	}
	if value > upper {
		return upper
	}
	return value
}

// Range returns a slice of integers from 0 up to (but not including) end. If end
// is negative it counts down.
func Range(end int) []int {
	return RangeStep(0, end, 1)
}

// RangeStep returns a slice of integers from start up to (but not including) end
// advancing by step. The step must be non-zero; the direction is inferred from
// the sign of (end - start) when step's sign disagrees.
func RangeStep(start, end, step int) []int {
	if step == 0 {
		panic("lodash: RangeStep step must not be zero")
	}
	var r []int
	if start < end {
		if step < 0 {
			step = -step
		}
		for i := start; i < end; i += step {
			r = append(r, i)
		}
	} else {
		if step > 0 {
			step = -step
		}
		for i := start; i > end; i += step {
			r = append(r, i)
		}
	}
	return r
}

// InRange reports whether value lies within the half-open interval [lower, upper).
// If lower is greater than upper the bounds are swapped.
func InRange[T Number](value, lower, upper T) bool {
	if lower > upper {
		lower, upper = upper, lower
	}
	return value >= lower && value < upper
}
