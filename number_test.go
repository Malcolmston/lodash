package lodash_test

import (
	"math/rand"
	"testing"

	lodash "github.com/malcolmston/lodash"
)

func TestRandomSeeded(t *testing.T) {
	rng1 := rand.New(rand.NewSource(1))
	rng2 := rand.New(rand.NewSource(1))
	for i := 0; i < 100; i++ {
		a := lodash.Random(1, 6, rng1)
		b := lodash.Random(1, 6, rng2)
		if a != b {
			t.Fatalf("Random not deterministic: %d != %d", a, b)
		}
		if a < 1 || a > 6 {
			t.Fatalf("Random out of range: %d", a)
		}
	}
	// swapped bounds
	rng := rand.New(rand.NewSource(2))
	v := lodash.Random(6, 1, rng)
	if v < 1 || v > 6 {
		t.Fatalf("Random swapped bounds out of range: %d", v)
	}
	f := lodash.RandomFloat(0, 1, rand.New(rand.NewSource(3)))
	if f < 0 || f >= 1 {
		t.Fatalf("RandomFloat out of range: %f", f)
	}
	if lodash.RandomFloat(5, 5, rand.New(rand.NewSource(3))) != 5 {
		t.Fatalf("RandomFloat equal bounds")
	}
}

func TestRoundFloorCeil(t *testing.T) {
	if lodash.Round(4.006, 2) != 4.01 {
		t.Errorf("Round = %v", lodash.Round(4.006, 2))
	}
	if lodash.Round(4060, -2) != 4100 {
		t.Errorf("Round negative precision = %v", lodash.Round(4060, -2))
	}
	if lodash.Floor(4.006, 2) != 4.0 {
		t.Errorf("Floor = %v", lodash.Floor(4.006, 2))
	}
	if lodash.Floor(4060, -2) != 4000 {
		t.Errorf("Floor negative precision")
	}
	if lodash.Ceil(4.001, 2) != 4.01 {
		t.Errorf("Ceil = %v", lodash.Ceil(4.001, 2))
	}
	if lodash.Ceil(4040, -2) != 4100 {
		t.Errorf("Ceil negative precision")
	}
}

func TestArithmetic(t *testing.T) {
	if lodash.Add(2, 3) != 5 || lodash.Subtract(5, 2) != 3 ||
		lodash.Multiply(4, 3) != 12 || lodash.Divide(10, 2) != 5 {
		t.Errorf("arithmetic")
	}
	if lodash.Add(1.5, 2.5) != 4.0 {
		t.Errorf("Add float")
	}
}

func TestRangeRight(t *testing.T) {
	if !lodash.IsEqual(lodash.RangeRight(4), []int{3, 2, 1, 0}) {
		t.Errorf("RangeRight = %v", lodash.RangeRight(4))
	}
}
