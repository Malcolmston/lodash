package lodash

import (
	"reflect"
	"testing"
)

func TestSetSeqPipeline(t *testing.T) {
	got := ChainSet([]int{1, 1, 2, 2, 3, 0}).
		Uniq().     // 1,2,3,0
		Compact().  // 1,2,3
		Without(2). // 1,3
		Value()
	if !reflect.DeepEqual(got, []int{1, 3}) {
		t.Errorf("SetSeq pipeline = %v, want [1 3]", got)
	}
}

func TestSetSeqCombinators(t *testing.T) {
	base := ChainSet([]int{1, 2, 3, 4})
	if got := base.Union([]int{4, 5, 6}).Value(); !reflect.DeepEqual(got, []int{1, 2, 3, 4, 5, 6}) {
		t.Errorf("Union = %v", got)
	}
	if got := base.Intersection([]int{2, 4, 6}).Value(); !reflect.DeepEqual(got, []int{2, 4}) {
		t.Errorf("Intersection = %v", got)
	}
	if got := base.Difference([]int{2, 4}).Value(); !reflect.DeepEqual(got, []int{1, 3}) {
		t.Errorf("Difference = %v", got)
	}
	if got := base.Filter(func(n int) bool { return n > 2 }).Value(); !reflect.DeepEqual(got, []int{3, 4}) {
		t.Errorf("Filter = %v", got)
	}
	if got := base.Reverse().Value(); !reflect.DeepEqual(got, []int{4, 3, 2, 1}) {
		t.Errorf("Reverse = %v", got)
	}
}

func TestSetSeqQueries(t *testing.T) {
	s := ChainSet([]string{"a", "b", "c"})
	if !s.Includes("b") || s.Includes("z") {
		t.Error("Includes failed")
	}
	if s.IndexOf("c") != 2 || s.IndexOf("z") != -1 {
		t.Errorf("IndexOf failed: %d", s.IndexOf("c"))
	}
	if s.Size() != 3 {
		t.Errorf("Size = %d", s.Size())
	}
}

func BenchmarkSetSeqPipeline(b *testing.B) {
	data := make([]int, 1000)
	for i := range data {
		data[i] = i % 200
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ChainSet(data).Uniq().Without(0, 1, 2).Value()
	}
}
