package lodash

import "testing"

// Vectors transcribed from lodash test/test.js (v4.17.21) for numeric helpers.

// TestParityClamp mirrors lodash.clamp (three-argument lower/upper form).
func TestParityClamp(t *testing.T) {
	cases := []struct {
		v, lo, hi, want float64
	}{
		{-10, -5, 5, -5},
		{-10.2, -5.5, 5.5, -5.5},
		{10, -5, 5, 5},
		{10.6, -5.6, 5.4, 5.4},
		{-4, -5, 5, -4},
		{-5, -5, 5, -5},
		{-5.5, -5.6, 5.6, -5.5},
		{4, -5, 5, 4},
		{5, -5, 5, 5},
		{4.5, -5.1, 5.2, 4.5},
	}
	for _, c := range cases {
		if got := Clamp(c.v, c.lo, c.hi); got != c.want {
			t.Errorf("Clamp(%v,%v,%v) = %v, want %v", c.v, c.lo, c.hi, got, c.want)
		}
	}
}

// TestParityInRange mirrors lodash.inRange. lodash's two-argument form
// inRange(v, end) maps to InRange(v, 0, end); the three-argument form maps
// directly, including start > end (bounds are swapped).
func TestParityInRange(t *testing.T) {
	two := []struct {
		v, end float64
		want   bool
	}{
		{3, 5, true},
		{5, 5, false},
		{6, 5, false},
		{0.5, 5, true},
		{5.2, 5, false},
	}
	for _, c := range two {
		if got := InRange(c.v, 0, c.end); got != c.want {
			t.Errorf("InRange(%v,0,%v) = %v, want %v", c.v, c.end, got, c.want)
		}
	}
	three := []struct {
		v, lo, hi float64
		want      bool
	}{
		{1, 1, 5, true},
		{3, 1, 5, true},
		{0, 1, 5, false},
		{5, 1, 5, false},
		{2, 5, 1, true},
		{-3, -2, -6, true},
		{1.2, 1, 5, true},
		{0.5, 1, 5, false},
	}
	for _, c := range three {
		if got := InRange(c.v, c.lo, c.hi); got != c.want {
			t.Errorf("InRange(%v,%v,%v) = %v, want %v", c.v, c.lo, c.hi, got, c.want)
		}
	}
}

// TestParityMean mirrors lodash.mean.
func TestParityMean(t *testing.T) {
	if got := Mean([]float64{4, 2, 8, 6}); got != 5 {
		t.Errorf("Mean = %v, want 5", got)
	}
}

// TestParityToInteger mirrors lodash.toInteger conversions.
func TestParityToInteger(t *testing.T) {
	if got := ToInteger(-5.6); got != -5 {
		t.Errorf("ToInteger(-5.6) = %d, want -5", got)
	}
	if got := ToInteger("5.6"); got != 5 {
		t.Errorf("ToInteger(5.6) = %d, want 5", got)
	}
	if got := ToInteger(nil); got != 0 {
		t.Errorf("ToInteger(nil) = %d, want 0", got)
	}
}
