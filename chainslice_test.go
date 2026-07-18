package lodash

import (
	"math/rand"
	"reflect"
	"sort"
	"testing"
)

func TestSliceSeqPipeline(t *testing.T) {
	got := ChainSlice([]int{5, 4, 3, 2, 1, 0}).
		Filter(func(n int) bool { return n%2 == 0 }). // 4,2,0
		Reverse().                                    // 0,2,4
		Drop(1).                                      // 2,4
		Value()
	if !reflect.DeepEqual(got, []int{2, 4}) {
		t.Errorf("pipeline = %v, want [2 4]", got)
	}
}

func TestSliceSeqTakeDrop(t *testing.T) {
	s := ChainSlice([]int{1, 2, 3, 4, 5})
	if got := s.Take(2).Value(); !reflect.DeepEqual(got, []int{1, 2}) {
		t.Errorf("Take = %v", got)
	}
	if got := s.TakeRight(2).Value(); !reflect.DeepEqual(got, []int{4, 5}) {
		t.Errorf("TakeRight = %v", got)
	}
	if got := s.DropRight(2).Value(); !reflect.DeepEqual(got, []int{1, 2, 3}) {
		t.Errorf("DropRight = %v", got)
	}
	if got := s.TakeWhile(func(n int) bool { return n < 3 }).Value(); !reflect.DeepEqual(got, []int{1, 2}) {
		t.Errorf("TakeWhile = %v", got)
	}
	if got := s.DropWhile(func(n int) bool { return n < 3 }).Value(); !reflect.DeepEqual(got, []int{3, 4, 5}) {
		t.Errorf("DropWhile = %v", got)
	}
	if got := s.Tail().Value(); !reflect.DeepEqual(got, []int{2, 3, 4, 5}) {
		t.Errorf("Tail = %v", got)
	}
	if got := s.Initial().Value(); !reflect.DeepEqual(got, []int{1, 2, 3, 4}) {
		t.Errorf("Initial = %v", got)
	}
	if got := s.Slice(1, 3).Value(); !reflect.DeepEqual(got, []int{2, 3}) {
		t.Errorf("Slice = %v", got)
	}
}

func TestSliceSeqTerminators(t *testing.T) {
	s := ChainSlice([]int{10, 20, 30})
	if h, ok := s.Head(); !ok || h != 10 {
		t.Errorf("Head = %d,%v", h, ok)
	}
	if l, ok := s.Last(); !ok || l != 30 {
		t.Errorf("Last = %d,%v", l, ok)
	}
	if n, ok := s.Nth(-1); !ok || n != 30 {
		t.Errorf("Nth(-1) = %d,%v", n, ok)
	}
	if s.Size() != 3 {
		t.Errorf("Size = %d", s.Size())
	}
	if s.IsEmpty() {
		t.Error("IsEmpty should be false")
	}
	if !ChainSlice([]int{}).IsEmpty() {
		t.Error("IsEmpty should be true")
	}
	if got := s.Join("-"); got != "10-20-30" {
		t.Errorf("Join = %q", got)
	}
	if got := s.Concat([]int{40}).Value(); !reflect.DeepEqual(got, []int{10, 20, 30, 40}) {
		t.Errorf("Concat = %v", got)
	}
	if got := s.Chunk(2); !reflect.DeepEqual(got, [][]int{{10, 20}, {30}}) {
		t.Errorf("Chunk = %v", got)
	}
}

func TestSliceSeqForEachTapShuffleSample(t *testing.T) {
	seen := 0
	ChainSlice([]int{1, 2, 3}).ForEach(func(n int) { seen += n })
	if seen != 6 {
		t.Errorf("ForEach sum = %d", seen)
	}
	var tapped []int
	ChainSlice([]int{1, 2}).Tap(func(s []int) { tapped = s })
	if !reflect.DeepEqual(tapped, []int{1, 2}) {
		t.Errorf("Tap = %v", tapped)
	}
	rng := rand.New(rand.NewSource(1))
	shuffled := ChainSlice([]int{1, 2, 3, 4, 5}).Shuffle(rng).Value()
	cp := append([]int(nil), shuffled...)
	sort.Ints(cp)
	if !reflect.DeepEqual(cp, []int{1, 2, 3, 4, 5}) {
		t.Errorf("Shuffle lost elements: %v", shuffled)
	}
	if _, ok := ChainSlice([]int{7}).Sample(rng); !ok {
		t.Error("Sample of non-empty failed")
	}
}
