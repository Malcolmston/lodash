package lodash

import (
	"reflect"
	"testing"
)

// Vectors transcribed from lodash test/test.js (v4.17.21) for the remaining
// numeric helpers not covered in math_parity_test.go.

// TestParityRoundCeilFloor mirrors the shared each(['ceil','floor','round'])
// precision assertions in test.js.
func TestParityRoundCeilFloor(t *testing.T) {
	// func(4.006): ceil 5, floor 4, round 4.
	if got := Ceil(4.006, 0); got != 5 {
		t.Errorf("Ceil(4.006,0) = %v, want 5", got)
	}
	if got := Floor(4.006, 0); got != 4 {
		t.Errorf("Floor(4.006,0) = %v, want 4", got)
	}
	if got := Round(4.006, 0); got != 4 {
		t.Errorf("Round(4.006,0) = %v, want 4", got)
	}
	// func(4.016, 2): floor 4.01, else 4.02.
	if got := Floor(4.016, 2); got != 4.01 {
		t.Errorf("Floor(4.016,2) = %v, want 4.01", got)
	}
	if got := Ceil(4.016, 2); got != 4.02 {
		t.Errorf("Ceil(4.016,2) = %v, want 4.02", got)
	}
	if got := Round(4.016, 2); got != 4.02 {
		t.Errorf("Round(4.016,2) = %v, want 4.02", got)
	}
	// func(4.1, 2) === 4.1 for all three.
	for _, got := range []float64{Ceil(4.1, 2), Floor(4.1, 2), Round(4.1, 2)} {
		if got != 4.1 {
			t.Errorf("round-family(4.1,2) = %v, want 4.1", got)
		}
	}
	// func(4160, -2): floor 4100, else 4200.
	if got := Floor(4160, -2); got != 4100 {
		t.Errorf("Floor(4160,-2) = %v, want 4100", got)
	}
	if got := Ceil(4160, -2); got != 4200 {
		t.Errorf("Ceil(4160,-2) = %v, want 4200", got)
	}
	if got := Round(4160, -2); got != 4200 {
		t.Errorf("Round(4160,-2) = %v, want 4200", got)
	}
}

// TestParityMaxMin mirrors lodash.max / lodash.min.
func TestParityMaxMin(t *testing.T) {
	if v, ok := Max([]int{1, 2, 3}); !ok || v != 3 {
		t.Errorf("Max([1 2 3]) = %d,%v, want 3,true", v, ok)
	}
	if v, ok := Min([]int{1, 2, 3}); !ok || v != 1 {
		t.Errorf("Min([1 2 3]) = %d,%v, want 1,true", v, ok)
	}
	if v, ok := Max([]string{"a", "b"}); !ok || v != "b" {
		t.Errorf("Max([a b]) = %q,%v, want b,true", v, ok)
	}
	if v, ok := Min([]string{"a", "b"}); !ok || v != "a" {
		t.Errorf("Min([a b]) = %q,%v, want a,true", v, ok)
	}
	// lodash: _.max([]) === undefined.
	if _, ok := Max([]int{}); ok {
		t.Errorf("Max([]) ok = true, want false")
	}
}

// TestParityRange mirrors lodash.range / lodash.rangeRight documented examples.
func TestParityRange(t *testing.T) {
	if got := Range(4); !reflect.DeepEqual(got, []int{0, 1, 2, 3}) {
		t.Errorf("Range(4) = %v, want [0 1 2 3]", got)
	}
	if got := RangeStep(1, 11, 2); !reflect.DeepEqual(got, []int{1, 3, 5, 7, 9}) {
		t.Errorf("RangeStep(1,11,2) = %v, want [1 3 5 7 9]", got)
	}
	if got := RangeStep(0, -4, -1); !reflect.DeepEqual(got, []int{0, -1, -2, -3}) {
		t.Errorf("RangeStep(0,-4,-1) = %v, want [0 -1 -2 -3]", got)
	}
	if got := RangeRight(4); !reflect.DeepEqual(got, []int{3, 2, 1, 0}) {
		t.Errorf("RangeRight(4) = %v, want [3 2 1 0]", got)
	}
}

// TestParitySum mirrors lodash.sum (documented example and empty case).
func TestParitySum(t *testing.T) {
	if got := Sum([]int{4, 2, 8, 6}); got != 20 {
		t.Errorf("Sum([4 2 8 6]) = %d, want 20", got)
	}
	if got := Sum([]int{}); got != 0 {
		t.Errorf("Sum([]) = %d, want 0", got)
	}
}

// TestParityArithmetic mirrors lodash.add / subtract / multiply / divide.
func TestParityArithmetic(t *testing.T) {
	if got := Add(6, 4); got != 10 {
		t.Errorf("Add(6,4) = %d, want 10", got)
	}
	if got := Subtract(6, 4); got != 2 {
		t.Errorf("Subtract(6,4) = %d, want 2", got)
	}
	if got := Multiply(6, 4); got != 24 {
		t.Errorf("Multiply(6,4) = %d, want 24", got)
	}
	if got := Divide(6.0, 4.0); got != 1.5 {
		t.Errorf("Divide(6,4) = %v, want 1.5", got)
	}
}
