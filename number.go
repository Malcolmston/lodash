package lodash

import (
	"math"
	"math/rand"
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

// scaleFor returns 10 raised to precision, used to shift the decimal point for
// precision-aware rounding.
func scaleFor(precision int) float64 {
	return math.Pow(10, float64(precision))
}

// Round rounds value to precision decimal places using round-half-away-from-zero.
// A negative precision rounds to the left of the decimal point.
func Round(value float64, precision int) float64 {
	scale := scaleFor(precision)
	return math.Round(value*scale) / scale
}

// Ceil rounds value up to precision decimal places. A negative precision rounds
// to the left of the decimal point.
func Ceil(value float64, precision int) float64 {
	scale := scaleFor(precision)
	return math.Ceil(value*scale) / scale
}

// Floor rounds value down to precision decimal places. A negative precision
// rounds to the left of the decimal point.
func Floor(value float64, precision int) float64 {
	scale := scaleFor(precision)
	return math.Floor(value*scale) / scale
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
