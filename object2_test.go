package lodash_test

import (
	"sort"
	"testing"

	lodash "github.com/malcolmston/lodash"
)

func TestToPairsFromPairs(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2}
	pairs := lodash.ToPairs(m)
	sort.Slice(pairs, func(i, j int) bool { return pairs[i].First < pairs[j].First })
	if len(pairs) != 2 || pairs[0].First != "a" || pairs[0].Second != 1 {
		t.Errorf("ToPairs = %v", pairs)
	}
	back := lodash.FromPairs(pairs)
	if !lodash.IsEqual(back, m) {
		t.Errorf("FromPairs = %v", back)
	}
}

func TestZipObject(t *testing.T) {
	got := lodash.ZipObject([]string{"a", "b", "c"}, []int{1, 2})
	want := map[string]int{"a": 1, "b": 2, "c": 0}
	if !lodash.IsEqual(got, want) {
		t.Errorf("ZipObject = %v", got)
	}
}

func TestFindKey(t *testing.T) {
	m := map[string]int{"a": 5}
	k, ok := lodash.FindKey(m, func(_ string, v int) bool { return v == 5 })
	if !ok || k != "a" {
		t.Errorf("FindKey = %v %v", k, ok)
	}
	if _, ok := lodash.FindKey(m, func(_ string, v int) bool { return v == 99 }); ok {
		t.Errorf("FindKey no match")
	}
}

func TestInvertBy(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 1}
	got := lodash.InvertBy(m, func(v int) int { return v })
	sort.Strings(got[1])
	if !lodash.IsEqual(got[1], []string{"a", "c"}) || !lodash.IsEqual(got[2], []string{"b"}) {
		t.Errorf("InvertBy = %v", got)
	}
}

func TestTransformForOwn(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	sum := lodash.Transform(m, func(acc int, v int, _ string) int { return acc + v }, 0)
	if sum != 6 {
		t.Errorf("Transform sum = %d", sum)
	}
	total := 0
	lodash.ForOwn(m, func(_ string, v int) { total += v })
	if total != 6 {
		t.Errorf("ForOwn total = %d", total)
	}
}

func TestMergeWithAssignWith(t *testing.T) {
	a := map[string]int{"x": 1, "y": 2}
	b := map[string]int{"y": 3, "z": 4}
	merged := lodash.MergeWith(func(e, i int) int { return e + i }, a, b)
	want := map[string]int{"x": 1, "y": 5, "z": 4}
	if !lodash.IsEqual(merged, want) {
		t.Errorf("MergeWith = %v", merged)
	}
	assigned := lodash.AssignWith(a, func(d, s int) int { return s }, b)
	if assigned["y"] != 3 || assigned["z"] != 4 || assigned["x"] != 1 {
		t.Errorf("AssignWith = %v", assigned)
	}
	// inputs unchanged
	if a["y"] != 2 {
		t.Errorf("AssignWith mutated input")
	}
}

func TestDefaults(t *testing.T) {
	dst := map[string]int{"a": 1}
	got := lodash.Defaults(dst, map[string]int{"a": 99, "b": 2}, map[string]int{"b": 100, "c": 3})
	want := map[string]int{"a": 1, "b": 2, "c": 3}
	if !lodash.IsEqual(got, want) {
		t.Errorf("Defaults = %v", got)
	}
	if dst["b"] != 0 && len(dst) != 1 {
		t.Errorf("Defaults mutated input")
	}
}

func TestDefaultsDeep(t *testing.T) {
	dst := map[string]any{"a": map[string]any{"x": 1}}
	src := map[string]any{"a": map[string]any{"x": 99, "y": 2}, "b": 3}
	got := lodash.DefaultsDeep(dst, src)
	if v, _ := lodash.Get(got, "a.x"); v != 1 {
		t.Errorf("DefaultsDeep should keep existing, got %v", v)
	}
	if v, _ := lodash.Get(got, "a.y"); v != 2 {
		t.Errorf("DefaultsDeep should add nested, got %v", v)
	}
	if v, _ := lodash.Get(got, "b"); v != 3 {
		t.Errorf("DefaultsDeep top-level = %v", v)
	}
}

func TestSize(t *testing.T) {
	if lodash.Size([]int{1, 2, 3}) != 3 || lodash.Size("abc") != 3 ||
		lodash.Size(map[string]int{"a": 1}) != 1 || lodash.Size(5) != 0 || lodash.Size(nil) != 0 {
		t.Errorf("Size failed")
	}
}
