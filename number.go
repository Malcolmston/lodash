package lodash

import (
	"math"
	"math/rand"
	"strconv"
	"strings"
)

// Random returns a pseudo-random integer in the inclusive range [min, max]
// drawn from rng, so results are deterministic when rng is seeded. If min is
// greater than max the bounds are swapped.
func Random(min, max int, rng *rand.Rand) int {
	if min > max {
		min, max = max, min
	}
	return min + rng.Intn(max-min+1)
}

// RandomFloat returns a pseudo-random float64 in the half-open range [min, max)
// drawn from rng. If min is greater than max the bounds are swapped.
func RandomFloat(min, max float64, rng *rand.Rand) float64 {
	if min > max {
		min, max = max, min
	}
	return min + rng.Float64()*(max-min)
}

// shiftExp returns value multiplied by 10^shift, computed by adjusting the
// decimal exponent of value's shortest textual representation rather than by
// floating-point multiplication. This mirrors lodash's createRound, which
// shifts the exponent through a string round-trip so that results like
// Floor(4.1, 2) are exactly 4.1 instead of the 4.09 a naive value*100 would
// yield. Non-finite and zero values are returned unchanged.
func shiftExp(value float64, shift int) float64 {
	if value == 0 || math.IsNaN(value) || math.IsInf(value, 0) {
		return value
	}
	s := strconv.FormatFloat(value, 'e', -1, 64)
	i := strings.IndexByte(s, 'e')
	exp, err := strconv.Atoi(s[i+1:])
	if err != nil {
		return value
	}
	out, err := strconv.ParseFloat(s[:i]+"e"+strconv.Itoa(exp+shift), 64)
	if err != nil {
		return value
	}
	return out
}

// roundWith applies op to value at the given decimal precision, using the same
// exponent-shifting strategy as lodash's _.round/_.ceil/_.floor to avoid binary
// floating-point artifacts.
func roundWith(value float64, precision int, op func(float64) float64) float64 {
	if precision == 0 || value == 0 || math.IsNaN(value) || math.IsInf(value, 0) {
		return op(value)
	}
	return shiftExp(op(shiftExp(value, precision)), -precision)
}

// Round rounds value to precision decimal places using round-half-away-from-zero.
// A negative precision rounds to the left of the decimal point. Like lodash's
// _.round it shifts the decimal exponent rather than scaling by a power of ten,
// so Round(4.006, 2) is exactly 4.01.
func Round(value float64, precision int) float64 {
	return roundWith(value, precision, math.Round)
}

// Ceil rounds value up to precision decimal places. A negative precision rounds
// to the left of the decimal point. Like lodash's _.ceil it shifts the decimal
// exponent to avoid floating-point error, so Ceil(4.016, 2) is exactly 4.02.
func Ceil(value float64, precision int) float64 {
	return roundWith(value, precision, math.Ceil)
}

// Floor rounds value down to precision decimal places. A negative precision
// rounds to the left of the decimal point. Like lodash's _.floor it shifts the
// decimal exponent to avoid floating-point error, so Floor(4.1, 2) is exactly
// 4.1 rather than 4.09.
func Floor(value float64, precision int) float64 {
	return roundWith(value, precision, math.Floor)
}

// Add returns the sum of a and b.
func Add[T Number](a, b T) T { return a + b }

// Subtract returns the difference a minus b.
func Subtract[T Number](a, b T) T { return a - b }

// Multiply returns the product of a and b.
func Multiply[T Number](a, b T) T { return a * b }

// Divide returns the quotient a divided by b.
func Divide[T Number](a, b T) T { return a / b }

// RangeRight is like Range but the returned integers are in descending order.
func RangeRight(end int) []int {
	return Reverse(Range(end))
}
