package lodash

import (
	"reflect"
	"testing"
)

func TestUniqWith(t *testing.T) {
	type pt struct{ X, Y int }
	pts := []pt{{1, 2}, {1, 2}, {3, 4}, {3, 4}, {5, 6}}
	got := UniqWith(pts, func(a, b pt) bool { return a == b })
	want := []pt{{1, 2}, {3, 4}, {5, 6}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("UniqWith = %v, want %v", got, want)
	}
	// Case-insensitive strings.
	gotS := UniqWith([]string{"A", "a", "B", "b"}, func(a, b string) bool {
		return ToLower(a) == ToLower(b)
	})
	if !reflect.DeepEqual(gotS, []string{"A", "B"}) {
		t.Errorf("UniqWith strings = %v", gotS)
	}
}

func TestPullAllWith(t *testing.T) {
	got := PullAllWith(
		[]float64{1.0, 2.0, 3.0, 4.0},
		[]float64{2.0, 4.0},
		func(a, b float64) bool { return a == b },
	)
	if !reflect.DeepEqual(got, []float64{1.0, 3.0}) {
		t.Errorf("PullAllWith = %v", got)
	}
}

func TestZipObjectDeep(t *testing.T) {
	got := ZipObjectDeep([]string{"a.b[0].c", "a.b[1].d"}, []any{1, 2})
	v1, ok := Get(got, "a.b[0].c")
	if !ok || v1 != 1 {
		t.Errorf("ZipObjectDeep a.b[0].c = %v,%v", v1, ok)
	}
	v2, ok := Get(got, "a.b[1].d")
	if !ok || v2 != 2 {
		t.Errorf("ZipObjectDeep a.b[1].d = %v,%v", v2, ok)
	}
	// Missing value becomes nil.
	got2 := ZipObjectDeep([]string{"x.y"}, nil)
	if v, ok := Get(got2, "x.y"); !ok || v != nil {
		t.Errorf("ZipObjectDeep missing value = %v,%v", v, ok)
	}
}

func TestFlatMapDeepAndDepth(t *testing.T) {
	src := []any{1, 2, 3}
	deep := FlatMapDeep[int](src, func(v any) []any {
		n := v.(int)
		return []any{n, []any{n * 10, []any{n * 100}}}
	})
	want := []int{1, 10, 100, 2, 20, 200, 3, 30, 300}
	if !reflect.DeepEqual(deep, want) {
		t.Errorf("FlatMapDeep = %v, want %v", deep, want)
	}
	depth1 := FlatMapDepth[int]([]any{1, 2}, func(v any) []any {
		n := v.(int)
		return []any{[]any{n}}
	}, 1)
	// One level of flattening leaves the inner []any wrapping intact as []int? no:
	// mapped = [[1] [2]] (each is []any); depth 1 flattens one level -> 1,2.
	if !reflect.DeepEqual(depth1, []int{1, 2}) {
		t.Errorf("FlatMapDepth depth1 = %v", depth1)
	}
}

func TestFillRange(t *testing.T) {
	cases := []struct {
		in              []int
		val, start, end int
		want            []int
	}{
		{[]int{1, 2, 3, 4}, 0, 1, 3, []int{1, 0, 0, 4}},
		{[]int{1, 2, 3, 4}, 9, 0, 4, []int{9, 9, 9, 9}},
		{[]int{1, 2, 3, 4}, 5, -2, 4, []int{1, 2, 5, 5}},
		{[]int{1, 2, 3, 4}, 5, 2, 100, []int{1, 2, 5, 5}},
	}
	for _, c := range cases {
		got := FillRange(c.in, c.val, c.start, c.end)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("FillRange(%v,%d,%d,%d) = %v, want %v", c.in, c.val, c.start, c.end, got, c.want)
		}
	}
	// Original is not mutated.
	orig := []int{1, 2, 3}
	_ = FillRange(orig, 0, 0, 3)
	if !reflect.DeepEqual(orig, []int{1, 2, 3}) {
		t.Errorf("FillRange mutated input: %v", orig)
	}
}

func BenchmarkUniqWith(b *testing.B) {
	data := make([]int, 500)
	for i := range data {
		data[i] = i % 100
	}
	eq := func(a, c int) bool { return a == c }
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = UniqWith(data, eq)
	}
}
