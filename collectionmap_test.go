package lodash

import (
	"sort"
	"testing"
)

func cmSorted(s []int) []int {
	out := append([]int(nil), s...)
	sort.Ints(out)
	return out
}

func cmEqInts(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestEverySomeNoneMap(t *testing.T) {
	m := map[string]int{"a": 2, "b": 4, "c": 6}
	even := func(v int, _ string) bool { return v%2 == 0 }
	if !EveryMap(m, even) {
		t.Error("EveryMap = false")
	}
	if EveryMap(m, func(v int, _ string) bool { return v > 3 }) {
		t.Error("EveryMap should be false")
	}
	if !SomeMap(m, func(v int, _ string) bool { return v == 4 }) {
		t.Error("SomeMap = false")
	}
	if SomeMap(m, func(v int, _ string) bool { return v == 5 }) {
		t.Error("SomeMap should be false")
	}
	if !NoneMap(m, func(v int, _ string) bool { return v == 5 }) {
		t.Error("NoneMap = false")
	}
	if !EveryMap(map[string]int{}, even) {
		t.Error("EveryMap(empty) should be true")
	}
	if SomeMap(map[string]int{}, even) {
		t.Error("SomeMap(empty) should be false")
	}
}

func TestFindFilterRejectMap(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
	v, ok := FindMap(m, func(v int, _ string) bool { return v == 3 })
	if !ok || v != 3 {
		t.Errorf("FindMap = %d,%v", v, ok)
	}
	if _, ok := FindMap(m, func(v int, _ string) bool { return v == 99 }); ok {
		t.Error("FindMap should miss")
	}
	got := cmSorted(FilterMap(m, func(v int, _ string) bool { return v%2 == 0 }))
	if !cmEqInts(got, []int{2, 4}) {
		t.Errorf("FilterMap = %v", got)
	}
	got = cmSorted(RejectMap(m, func(v int, _ string) bool { return v%2 == 0 }))
	if !cmEqInts(got, []int{1, 3}) {
		t.Errorf("RejectMap = %v", got)
	}
}

func TestMapToSliceReduceMap(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	doubled := cmSorted(MapToSlice(m, func(v int, _ string) int { return v * 2 }))
	if !cmEqInts(doubled, []int{2, 4, 6}) {
		t.Errorf("MapToSlice = %v", doubled)
	}
	sum := ReduceMap(m, func(acc, v int, _ string) int { return acc + v }, 0)
	if sum != 6 {
		t.Errorf("ReduceMap = %d", sum)
	}
}

func TestPartitionGroupIncludesMap(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
	yes, no := PartitionMap(m, func(v int, _ string) bool { return v%2 == 0 })
	if !cmEqInts(cmSorted(yes), []int{2, 4}) || !cmEqInts(cmSorted(no), []int{1, 3}) {
		t.Errorf("PartitionMap = %v / %v", yes, no)
	}
	groups := GroupByMap(m, func(v int) string {
		if v%2 == 0 {
			return "even"
		}
		return "odd"
	})
	if !cmEqInts(cmSorted(groups["even"]), []int{2, 4}) {
		t.Errorf("GroupByMap even = %v", groups["even"])
	}
	if !cmEqInts(cmSorted(groups["odd"]), []int{1, 3}) {
		t.Errorf("GroupByMap odd = %v", groups["odd"])
	}
	if !IncludesValue(m, 3) || IncludesValue(m, 99) {
		t.Error("IncludesValue failed")
	}
}

func TestMinMaxByMap(t *testing.T) {
	m := map[string]int{"a": 3, "b": 1, "c": 2}
	mn, ok := MinByMap(m, func(v int) int { return v })
	if !ok || mn != 1 {
		t.Errorf("MinByMap = %d,%v", mn, ok)
	}
	mx, ok := MaxByMap(m, func(v int) int { return v })
	if !ok || mx != 3 {
		t.Errorf("MaxByMap = %d,%v", mx, ok)
	}
	if _, ok := MinByMap(map[string]int{}, func(v int) int { return v }); ok {
		t.Error("MinByMap(empty) should be false")
	}
}
