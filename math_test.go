package lodash

import (
	"math"
	"reflect"
	"testing"
)

func TestSumMean(t *testing.T) {
	if Sum([]int{1, 2, 3, 4}) != 10 {
		t.Fatal("sum")
	}
	if Sum([]float64{}) != 0 {
		t.Fatal("sum empty")
	}
	if SumBy([]string{"a", "bb", "ccc"}, func(s string) int { return len(s) }) != 6 {
		t.Fatal("sumby")
	}
	if Mean([]int{2, 4, 6}) != 4 {
		t.Fatal("mean")
	}
	if Mean([]int{}) != 0 {
		t.Fatal("mean empty")
	}
	if MeanBy([]string{"a", "bbb"}, func(s string) int { return len(s) }) != 2 {
		t.Fatal("meanby")
	}
	if MeanBy([]int{}, func(x int) int { return x }) != 0 {
		t.Fatal("meanby empty")
	}
}

func TestMinMax(t *testing.T) {
	mn, ok := Min([]int{3, 1, 2})
	if !ok || mn != 1 {
		t.Fatalf("min = %d", mn)
	}
	mx, ok := Max([]int{3, 1, 2})
	if !ok || mx != 3 {
		t.Fatalf("max = %d", mx)
	}
	if _, ok := Min([]int{}); ok {
		t.Fatal("min empty")
	}
	if _, ok := Max([]int{}); ok {
		t.Fatal("max empty")
	}
	pl := []person{{"A", 30}, {"B", 20}, {"C", 40}}
	yng, _ := MinBy(pl, func(p person) int { return p.Age })
	if yng.Name != "B" {
		t.Fatalf("minby = %v", yng)
	}
	old, _ := MaxBy(pl, func(p person) int { return p.Age })
	if old.Name != "C" {
		t.Fatalf("maxby = %v", old)
	}
	if _, ok := MinBy([]person{}, func(p person) int { return p.Age }); ok {
		t.Fatal("minby empty")
	}
	if _, ok := MaxBy([]person{}, func(p person) int { return p.Age }); ok {
		t.Fatal("maxby empty")
	}
}

func TestClamp(t *testing.T) {
	if Clamp(5, 1, 10) != 5 {
		t.Fatal("clamp mid")
	}
	if Clamp(-3, 1, 10) != 1 {
		t.Fatal("clamp low")
	}
	if Clamp(99, 1, 10) != 10 {
		t.Fatal("clamp high")
	}
	if Clamp(5, 10, 1) != 5 {
		t.Fatal("clamp swapped bounds")
	}
}

func TestRange(t *testing.T) {
	if !reflect.DeepEqual(Range(4), []int{0, 1, 2, 3}) {
		t.Fatal("range")
	}
	if !reflect.DeepEqual(Range(-3), []int{0, -1, -2}) {
		t.Fatal("range negative")
	}
	if !reflect.DeepEqual(RangeStep(1, 10, 2), []int{1, 3, 5, 7, 9}) {
		t.Fatal("rangestep")
	}
	if !reflect.DeepEqual(RangeStep(10, 0, -2), []int{10, 8, 6, 4, 2}) {
		t.Fatal("rangestep down")
	}
	// Sign correction: positive step but descending target.
	if !reflect.DeepEqual(RangeStep(5, 0, 1), []int{5, 4, 3, 2, 1}) {
		t.Fatal("rangestep sign fix")
	}
	if len(Range(0)) != 0 {
		t.Fatal("range zero")
	}
}

func TestRangeStepPanics(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Fatal("expected panic on zero step")
		}
	}()
	RangeStep(0, 10, 0)
}

func TestInRange(t *testing.T) {
	if !InRange(3, 1, 5) {
		t.Fatal("inrange in")
	}
	if InRange(5, 1, 5) {
		t.Fatal("inrange upper exclusive")
	}
	if !InRange(1, 1, 5) {
		t.Fatal("inrange lower inclusive")
	}
	if !InRange(3.5, 5.0, 1.0) {
		t.Fatal("inrange swapped")
	}
	if InRange(-1, 0, 5) {
		t.Fatal("inrange below")
	}
}

func TestMeanFloatPrecision(t *testing.T) {
	got := Mean([]float64{1, 2})
	if math.Abs(got-1.5) > 1e-9 {
		t.Fatalf("mean float = %v", got)
	}
}
