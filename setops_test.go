package lodash_test

import (
	"testing"

	lodash "github.com/malcolmston/lodash"
)

func TestHeadFirstLastNth(t *testing.T) {
	s := []int{10, 20, 30}
	if v, ok := lodash.Head(s); !ok || v != 10 {
		t.Errorf("Head")
	}
	if v, _ := lodash.First(s); v != 10 {
		t.Errorf("First")
	}
	if v, ok := lodash.Last(s); !ok || v != 30 {
		t.Errorf("Last")
	}
	if v, _ := lodash.Nth(s, 1); v != 20 {
		t.Errorf("Nth")
	}
	if v, _ := lodash.Nth(s, -1); v != 30 {
		t.Errorf("Nth neg")
	}
	if _, ok := lodash.Nth(s, 9); ok {
		t.Errorf("Nth out of range")
	}
	if _, ok := lodash.Head([]int{}); ok {
		t.Errorf("Head empty")
	}
}

func TestInitialTailSlice(t *testing.T) {
	s := []int{1, 2, 3, 4}
	if !lodash.IsEqual(lodash.Initial(s), []int{1, 2, 3}) {
		t.Errorf("Initial")
	}
	if !lodash.IsEqual(lodash.Tail(s), []int{2, 3, 4}) {
		t.Errorf("Tail")
	}
	if !lodash.IsEqual(lodash.Slice(s, 1, 3), []int{2, 3}) {
		t.Errorf("Slice")
	}
	if !lodash.IsEqual(lodash.Slice(s, -2, 4), []int{3, 4}) {
		t.Errorf("Slice negative")
	}
	if len(lodash.Slice(s, 3, 1)) != 0 {
		t.Errorf("Slice empty")
	}
	if len(lodash.Tail([]int{1})) != 0 {
		t.Errorf("Tail single")
	}
}

func TestJoin(t *testing.T) {
	if lodash.Join([]int{1, 2, 3}, "-") != "1-2-3" {
		t.Errorf("Join")
	}
}

func TestDifferenceByWith(t *testing.T) {
	got := lodash.DifferenceBy(func(f float64) float64 { return lodash.Floor(f, 0) }, []float64{1.2, 2.3, 3.4}, []float64{2.9})
	if !lodash.IsEqual(got, []float64{1.2, 3.4}) {
		t.Errorf("DifferenceBy = %v", got)
	}
	eq := func(a, b int) bool { return a == b }
	if !lodash.IsEqual(lodash.DifferenceWith(eq, []int{1, 2, 3}, []int{2}), []int{1, 3}) {
		t.Errorf("DifferenceWith")
	}
}

func TestIntersectionByWith(t *testing.T) {
	got := lodash.IntersectionBy(func(n int) int { return n % 10 }, []int{1, 2, 3}, []int{12, 15})
	if !lodash.IsEqual(got, []int{2}) {
		t.Errorf("IntersectionBy = %v", got)
	}
	eq := func(a, b int) bool { return a == b }
	if !lodash.IsEqual(lodash.IntersectionWith(eq, []int{1, 2, 3}, []int{2, 3, 4}), []int{2, 3}) {
		t.Errorf("IntersectionWith")
	}
	if lodash.IntersectionBy(func(n int) int { return n }) != nil {
		t.Errorf("IntersectionBy empty")
	}
}

func TestUnionByWith(t *testing.T) {
	got := lodash.UnionBy(func(n int) int { return n % 10 }, []int{1, 2}, []int{12, 3})
	if !lodash.IsEqual(got, []int{1, 2, 3}) {
		t.Errorf("UnionBy = %v", got)
	}
	eq := func(a, b int) bool { return a == b }
	if !lodash.IsEqual(lodash.UnionWith(eq, []int{1, 2}, []int{2, 3}), []int{1, 2, 3}) {
		t.Errorf("UnionWith")
	}
}

func TestXor(t *testing.T) {
	if !lodash.IsEqual(lodash.Xor([]int{1, 2}, []int{2, 3}), []int{1, 3}) {
		t.Errorf("Xor")
	}
	got := lodash.XorBy(func(n int) int { return n % 10 }, []int{1, 2}, []int{12, 3})
	if !lodash.IsEqual(got, []int{1, 3}) {
		t.Errorf("XorBy = %v", got)
	}
	eq := func(a, b int) bool { return a == b }
	if !lodash.IsEqual(lodash.XorWith(eq, []int{1, 2}, []int{2, 3}), []int{1, 3}) {
		t.Errorf("XorWith")
	}
}

func TestSortedUniq(t *testing.T) {
	if !lodash.IsEqual(lodash.SortedUniq([]int{1, 1, 2, 3, 3}), []int{1, 2, 3}) {
		t.Errorf("SortedUniq")
	}
	got := lodash.SortedUniqBy([]float64{1.1, 1.2, 2.3}, func(f float64) int { return int(f) })
	if !lodash.IsEqual(got, []float64{1.1, 2.3}) {
		t.Errorf("SortedUniqBy = %v", got)
	}
}

func TestPullRemove(t *testing.T) {
	s := []int{1, 2, 3, 2, 1}
	if !lodash.IsEqual(lodash.PullAll(s, []int{2}), []int{1, 3, 1}) {
		t.Errorf("PullAll")
	}
	if !lodash.IsEqual(lodash.Pull(s, 1, 3), []int{2, 2}) {
		t.Errorf("Pull")
	}
	got := lodash.PullAllBy([]int{1, 2, 3}, []int{12}, func(n int) int { return n % 10 })
	if !lodash.IsEqual(got, []int{1, 3}) {
		t.Errorf("PullAllBy = %v", got)
	}
	if !lodash.IsEqual(lodash.PullAt(s, 0, -1), []int{2, 3, 2}) {
		t.Errorf("PullAt")
	}
	if !lodash.IsEqual(lodash.Remove(s, func(n int) bool { return n == 2 }), []int{1, 3, 1}) {
		t.Errorf("Remove")
	}
	// original not mutated
	if !lodash.IsEqual(s, []int{1, 2, 3, 2, 1}) {
		t.Errorf("mutation")
	}
}

func TestFlatMapDepth(t *testing.T) {
	got := lodash.FlatMap([]int{1, 2}, func(n int) []int { return []int{n, n} })
	if !lodash.IsEqual(got, []int{1, 1, 2, 2}) {
		t.Errorf("FlatMap = %v", got)
	}
	nested := []any{1, []any{2, []any{3}}}
	if !lodash.IsEqual(lodash.FlattenDepth[int](nested, 1), []int{1, 2}) {
		t.Errorf("FlattenDepth 1")
	}
	if !lodash.IsEqual(lodash.FlattenDepth[int](nested, 2), []int{1, 2, 3}) {
		t.Errorf("FlattenDepth 2")
	}
}

func TestTakeDropWhile(t *testing.T) {
	s := []int{1, 2, 3, 4, 1}
	lt3 := func(n int) bool { return n < 3 }
	if !lodash.IsEqual(lodash.TakeWhile(s, lt3), []int{1, 2}) {
		t.Errorf("TakeWhile")
	}
	if !lodash.IsEqual(lodash.DropWhile(s, lt3), []int{3, 4, 1}) {
		t.Errorf("DropWhile")
	}
	gt0 := func(n int) bool { return n < 3 }
	if !lodash.IsEqual(lodash.TakeRightWhile(s, gt0), []int{1}) {
		t.Errorf("TakeRightWhile")
	}
	if !lodash.IsEqual(lodash.DropRightWhile(s, gt0), []int{1, 2, 3, 4}) {
		t.Errorf("DropRightWhile")
	}
}

func TestZipWith(t *testing.T) {
	got := lodash.ZipWith([]int{1, 2, 3}, []int{10, 20}, func(a, b int) int { return a + b })
	if !lodash.IsEqual(got, []int{11, 22}) {
		t.Errorf("ZipWith = %v", got)
	}
	pairs := []lodash.Pair[int, int]{{First: 1, Second: 2}, {First: 3, Second: 4}}
	uz := lodash.UnzipWith(pairs, func(a, b int) int { return a * b })
	if !lodash.IsEqual(uz, []int{2, 12}) {
		t.Errorf("UnzipWith = %v", uz)
	}
}

func TestSortedIndex(t *testing.T) {
	s := []int{10, 20, 20, 30}
	if lodash.SortedIndex(s, 20) != 1 {
		t.Errorf("SortedIndex")
	}
	if lodash.SortedLastIndex(s, 20) != 3 {
		t.Errorf("SortedLastIndex")
	}
	if lodash.SortedIndexOf(s, 20) != 1 || lodash.SortedIndexOf(s, 99) != -1 {
		t.Errorf("SortedIndexOf")
	}
	if lodash.SortedLastIndexOf(s, 20) != 2 || lodash.SortedLastIndexOf(s, 99) != -1 {
		t.Errorf("SortedLastIndexOf")
	}
	if lodash.SortedIndexBy(s, 25, func(n int) int { return n }) != 3 {
		t.Errorf("SortedIndexBy")
	}
	if lodash.SortedLastIndexBy(s, 20, func(n int) int { return n }) != 3 {
		t.Errorf("SortedLastIndexBy")
	}
}

func TestIndexOfFromNone(t *testing.T) {
	s := []int{1, 2, 3, 2, 1}
	if lodash.IndexOfFrom(s, 2, 2) != 3 {
		t.Errorf("IndexOfFrom")
	}
	if lodash.IndexOfFrom(s, 9, 0) != -1 {
		t.Errorf("IndexOfFrom missing")
	}
	if lodash.LastIndexOfFrom(s, 2, 2) != 1 {
		t.Errorf("LastIndexOfFrom")
	}
	if lodash.LastIndexOfFrom(s, 1, -1) != 4 {
		t.Errorf("LastIndexOfFrom neg")
	}
	if lodash.None(s, func(n int) bool { return n > 5 }) != true {
		t.Errorf("None")
	}
	if lodash.None(s, func(n int) bool { return n == 2 }) != false {
		t.Errorf("None false")
	}
}

func TestForEachRight(t *testing.T) {
	var order []int
	lodash.ForEachRight([]int{1, 2, 3}, func(n int) { order = append(order, n) })
	if !lodash.IsEqual(order, []int{3, 2, 1}) {
		t.Errorf("ForEachRight = %v", order)
	}
}
