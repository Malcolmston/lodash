package lodash

import (
	"reflect"
	"testing"
)

// The vectors in this file are transcribed from lodash's own QUnit suite
// (lodash/lodash test/test.js at v4.17.21). Each assertion pins this port to the
// exact output the upstream tests assert, so behavioral drift from the original
// library is caught as a test failure.

// TestParityDifference mirrors the lodash.difference family assertions
// (test.js: "lodash.difference and lodash.differenceBy...").
func TestParityDifference(t *testing.T) {
	if got := Difference([]int{2, 1}, []int{2, 3}); !reflect.DeepEqual(got, []int{1}) {
		t.Errorf("Difference([2 1],[2 3]) = %v, want [1]", got)
	}
	if got := Difference([]int{2, 1, 2, 3}, []int{3, 4}, []int{3, 2}); !reflect.DeepEqual(got, []int{1}) {
		t.Errorf("Difference([2 1 2 3],[3 4],[3 2]) = %v, want [1]", got)
	}
}

// TestParityIntersection mirrors the lodash.intersection family assertions.
func TestParityIntersection(t *testing.T) {
	if got := Intersection([]int{2, 1}, []int{2, 3}); !reflect.DeepEqual(got, []int{2}) {
		t.Errorf("Intersection([2 1],[2 3]) = %v, want [2]", got)
	}
	if got := Intersection([]int{2, 1, 2, 3}, []int{3, 4}, []int{3, 2}); !reflect.DeepEqual(got, []int{3}) {
		t.Errorf("Intersection([2 1 2 3],[3 4],[3 2]) = %v, want [3]", got)
	}
	if got := Intersection([]int{1, 1, 3, 2, 2}, []int{5, 2, 2, 1, 4}, []int{2, 1, 1}); !reflect.DeepEqual(got, []int{1, 2}) {
		t.Errorf("Intersection dedupe = %v, want [1 2]", got)
	}
}

// TestParityUnion mirrors the lodash.union family assertions.
func TestParityUnion(t *testing.T) {
	if got := Union([]int{2}, []int{1, 2}); !reflect.DeepEqual(got, []int{2, 1}) {
		t.Errorf("Union([2],[1 2]) = %v, want [2 1]", got)
	}
	if got := Union([]int{2}, []int{1, 2}, []int{2, 3}); !reflect.DeepEqual(got, []int{2, 1, 3}) {
		t.Errorf("Union([2],[1 2],[2 3]) = %v, want [2 1 3]", got)
	}
}

// TestParityXor mirrors the lodash.xor family assertions, including the
// odd/even occurrence semantics (a value in every array cancels out).
func TestParityXor(t *testing.T) {
	if got := Xor([]int{2, 1}, []int{2, 3}); !reflect.DeepEqual(got, []int{1, 3}) {
		t.Errorf("Xor([2 1],[2 3]) = %v, want [1 3]", got)
	}
	if got := Xor([]int{2, 1}, []int{2, 3}, []int{3, 4}); !reflect.DeepEqual(got, []int{1, 4}) {
		t.Errorf("Xor([2 1],[2 3],[3 4]) = %v, want [1 4]", got)
	}
	if got := Xor([]int{1, 2}, []int{2, 1}, []int{1, 2}); len(got) != 0 {
		t.Errorf("Xor([1 2],[2 1],[1 2]) = %v, want []", got)
	}
}

// TestParityUniq mirrors lodash.uniq / lodash.sortedUniq.
func TestParityUniq(t *testing.T) {
	if got := Uniq([]int{2, 1, 2}); !reflect.DeepEqual(got, []int{2, 1}) {
		t.Errorf("Uniq([2 1 2]) = %v, want [2 1]", got)
	}
	if got := Uniq([]int{1, 2, 3, 1, 2, 3}); !reflect.DeepEqual(got, []int{1, 2, 3}) {
		t.Errorf("Uniq = %v, want [1 2 3]", got)
	}
	if got := SortedUniq([]int{1, 1, 2}); !reflect.DeepEqual(got, []int{1, 2}) {
		t.Errorf("SortedUniq([1 1 2]) = %v, want [1 2]", got)
	}
}

// TestParityWithout mirrors lodash.without.
func TestParityWithout(t *testing.T) {
	if got := Without([]int{2, 1, 2, 3}, 1, 2); !reflect.DeepEqual(got, []int{3}) {
		t.Errorf("Without([2 1 2 3],1,2) = %v, want [3]", got)
	}
}

// TestParityCompact mirrors lodash.compact (falsey values are dropped; for the
// int specialization the sole falsey value is 0).
func TestParityCompact(t *testing.T) {
	if got := Compact([]int{0, 1, 0, 2, 0, 3}); !reflect.DeepEqual(got, []int{1, 2, 3}) {
		t.Errorf("Compact = %v, want [1 2 3]", got)
	}
	if got := Compact([]string{"", "a", "", "b"}); !reflect.DeepEqual(got, []string{"a", "b"}) {
		t.Errorf("Compact strings = %v, want [a b]", got)
	}
}

// TestParityChunk mirrors lodash.chunk (documented examples).
func TestParityChunk(t *testing.T) {
	if got := Chunk([]string{"a", "b", "c", "d"}, 2); !reflect.DeepEqual(got, [][]string{{"a", "b"}, {"c", "d"}}) {
		t.Errorf("Chunk(...,2) = %v", got)
	}
	if got := Chunk([]string{"a", "b", "c", "d"}, 3); !reflect.DeepEqual(got, [][]string{{"a", "b", "c"}, {"d"}}) {
		t.Errorf("Chunk(...,3) = %v", got)
	}
}

// TestParityDropTake mirrors lodash.drop / dropRight / take / takeRight using
// the shared array [1,2,3].
func TestParityDropTake(t *testing.T) {
	arr := []int{1, 2, 3}
	if got := Drop(arr, 2); !reflect.DeepEqual(got, []int{3}) {
		t.Errorf("Drop([1 2 3],2) = %v, want [3]", got)
	}
	if got := DropRight(arr, 2); !reflect.DeepEqual(got, []int{1}) {
		t.Errorf("DropRight([1 2 3],2) = %v, want [1]", got)
	}
	if got := Take(arr, 2); !reflect.DeepEqual(got, []int{1, 2}) {
		t.Errorf("Take([1 2 3],2) = %v, want [1 2]", got)
	}
	if got := TakeRight(arr, 2); !reflect.DeepEqual(got, []int{2, 3}) {
		t.Errorf("TakeRight([1 2 3],2) = %v, want [2 3]", got)
	}
}

// TestParityHeadLastNth mirrors lodash.head/first, lodash.last and lodash.nth.
func TestParityHeadLastNth(t *testing.T) {
	arr := []string{"a", "b", "c", "d"}
	if v, ok := Head(arr); !ok || v != "a" {
		t.Errorf("Head = %q,%v, want a,true", v, ok)
	}
	if v, ok := Last(arr); !ok || v != "d" {
		t.Errorf("Last = %q,%v, want d,true", v, ok)
	}
	if v, ok := Nth(arr, 1); !ok || v != "b" {
		t.Errorf("Nth(1) = %q,%v, want b,true", v, ok)
	}
	// lodash: _.nth(array, -2) === 'c'
	if v, ok := Nth(arr, -2); !ok || v != "c" {
		t.Errorf("Nth(-2) = %q,%v, want c,true", v, ok)
	}
	if _, ok := Head([]int{}); ok {
		t.Errorf("Head([]) ok = true, want false")
	}
}

// TestParityInitialTail mirrors lodash.initial / lodash.tail.
func TestParityInitialTail(t *testing.T) {
	arr := []int{1, 2, 3}
	if got := Initial(arr); !reflect.DeepEqual(got, []int{1, 2}) {
		t.Errorf("Initial([1 2 3]) = %v, want [1 2]", got)
	}
	if got := Tail(arr); !reflect.DeepEqual(got, []int{2, 3}) {
		t.Errorf("Tail([1 2 3]) = %v, want [2 3]", got)
	}
}

// TestParityFlatten mirrors lodash.flatten (single level) and flattenDeep.
func TestParityFlatten(t *testing.T) {
	if got := Flatten([][]int{{1, 2}, {3, 4}}); !reflect.DeepEqual(got, []int{1, 2, 3, 4}) {
		t.Errorf("Flatten = %v, want [1 2 3 4]", got)
	}
	nested := []any{1, []any{2, []any{3, []any{4}}, 5}}
	if got := FlattenDeep[int](nested); !reflect.DeepEqual(got, []int{1, 2, 3, 4, 5}) {
		t.Errorf("FlattenDeep = %v, want [1 2 3 4 5]", got)
	}
}

// TestParityFlattenDepth mirrors lodash.flattenDepth on [1,[2,[3,[4]],5]].
func TestParityFlattenDepth(t *testing.T) {
	array := []any{1, []any{2, []any{3, []any{4}}, 5}}
	// depth 1 -> [1, 2, [3, [4]], 5]
	got1 := FlattenDepth[any](array, 1)
	want1 := []any{1, 2, []any{3, []any{4}}, 5}
	if !reflect.DeepEqual(got1, want1) {
		t.Errorf("FlattenDepth(depth 1) = %v, want %v", got1, want1)
	}
	// depth 2 -> [1, 2, 3, [4], 5]
	got2 := FlattenDepth[any](array, 2)
	want2 := []any{1, 2, 3, []any{4}, 5}
	if !reflect.DeepEqual(got2, want2) {
		t.Errorf("FlattenDepth(depth 2) = %v, want %v", got2, want2)
	}
}

// TestParityZipFromPairs mirrors lodash.zip, lodash.unzip, lodash.fromPairs and
// lodash.zipObject.
func TestParityZipFromPairs(t *testing.T) {
	pairs := Zip([]string{"a", "b"}, []int{1, 2})
	want := []Pair[string, int]{{First: "a", Second: 1}, {First: "b", Second: 2}}
	if !reflect.DeepEqual(pairs, want) {
		t.Errorf("Zip = %v, want %v", pairs, want)
	}
	ks, vs := Unzip(pairs)
	if !reflect.DeepEqual(ks, []string{"a", "b"}) || !reflect.DeepEqual(vs, []int{1, 2}) {
		t.Errorf("Unzip = %v,%v", ks, vs)
	}
	m := FromPairs(pairs)
	if !reflect.DeepEqual(m, map[string]int{"a": 1, "b": 2}) {
		t.Errorf("FromPairs = %v", m)
	}
	zo := ZipObject([]string{"a", "b"}, []int{1, 2})
	if !reflect.DeepEqual(zo, map[string]int{"a": 1, "b": 2}) {
		t.Errorf("ZipObject = %v", zo)
	}
}

// TestParityIndexOf mirrors lodash.indexOf.
func TestParityIndexOf(t *testing.T) {
	arr := []int{1, 2, 3, 1, 2, 3}
	if got := IndexOf(arr, 2); got != 1 {
		t.Errorf("IndexOf(2) = %d, want 1", got)
	}
	if got := IndexOf(arr, 4); got != -1 {
		t.Errorf("IndexOf(4) = %d, want -1", got)
	}
}
