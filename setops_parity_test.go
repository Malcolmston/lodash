package lodash

import (
	"reflect"
	"testing"
)

// Vectors transcribed from lodash test/test.js (v4.17.21) for the *By variants
// of the array set operations, using the floor iteratee exactly as the upstream
// "differenceBy/intersectionBy/unionBy/xorBy should accept an iteratee" tests do.
// These complement the plain set-op vectors in arrays_parity_test.go.

func floorFloat(f float64) float64 { return float64(int(f)) }

// TestParityDifferenceBy mirrors lodash.differenceBy([2.1,1.2],[2.3,3.4],floor).
func TestParityDifferenceBy(t *testing.T) {
	if got := DifferenceBy(floorFloat, []float64{2.1, 1.2}, []float64{2.3, 3.4}); !reflect.DeepEqual(got, []float64{1.2}) {
		t.Errorf("DifferenceBy = %v, want [1.2]", got)
	}
}

// TestParityIntersectionBy mirrors lodash.intersectionBy([2.1,1.2],[2.3,3.4],floor).
func TestParityIntersectionBy(t *testing.T) {
	if got := IntersectionBy(floorFloat, []float64{2.1, 1.2}, []float64{2.3, 3.4}); !reflect.DeepEqual(got, []float64{2.1}) {
		t.Errorf("IntersectionBy = %v, want [2.1]", got)
	}
}

// TestParityUnionBy mirrors lodash.unionBy([2.1],[1.2,2.3],floor).
func TestParityUnionBy(t *testing.T) {
	if got := UnionBy(floorFloat, []float64{2.1}, []float64{1.2, 2.3}); !reflect.DeepEqual(got, []float64{2.1, 1.2}) {
		t.Errorf("UnionBy = %v, want [2.1 1.2]", got)
	}
}

// TestParityXorBy mirrors lodash.xorBy([2.1,1.2],[2.3,3.4],floor).
func TestParityXorBy(t *testing.T) {
	if got := XorBy(floorFloat, []float64{2.1, 1.2}, []float64{2.3, 3.4}); !reflect.DeepEqual(got, []float64{1.2, 3.4}) {
		t.Errorf("XorBy = %v, want [1.2 3.4]", got)
	}
}

// TestParityNthSweep mirrors lodash.nth's positive and negative index sweeps
// over ['a','b','c','d'] (test.js: "should get the nth element" and
// "should work with a negative n").
func TestParityNthSweep(t *testing.T) {
	a := []string{"a", "b", "c", "d"}
	for i, want := range a {
		if v, ok := Nth(a, i); !ok || v != want {
			t.Errorf("Nth(%d) = %v,%v want %v", i, v, ok, want)
		}
	}
	neg := []string{"d", "c", "b", "a"}
	for i, want := range neg {
		if v, ok := Nth(a, -(i + 1)); !ok || v != want {
			t.Errorf("Nth(%d) = %v,%v want %v", -(i + 1), v, ok, want)
		}
	}
}
