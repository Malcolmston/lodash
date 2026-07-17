package lodash_test

import (
	"testing"

	lodash "github.com/malcolmston/lodash"
)

func TestSeqChain(t *testing.T) {
	var tapped []int
	result := lodash.Chain([]int{1, 2, 3, 4}).
		Thru(func(s []int) []int { return lodash.Filter(s, func(n int) bool { return n%2 == 0 }) }).
		Tap(func(s []int) { tapped = s }).
		Thru(func(s []int) []int { return lodash.Map(s, func(n int) int { return n * 10 }) }).
		Value()
	if !lodash.IsEqual(result, []int{20, 40}) {
		t.Errorf("Seq chain = %v", result)
	}
	if !lodash.IsEqual(tapped, []int{2, 4}) {
		t.Errorf("Seq tap = %v", tapped)
	}
}

func TestThruTapTypeChanging(t *testing.T) {
	sum := lodash.Thru(lodash.Chain([]int{1, 2, 3}), func(s []int) int {
		return lodash.Sum(s)
	})
	if sum.Value() != 6 {
		t.Errorf("Thru type-changing = %d", sum.Value())
	}
	var seen int
	out := lodash.Tap(lodash.Chain(42), func(v int) { seen = v })
	if out.Value() != 42 || seen != 42 {
		t.Errorf("Tap free func")
	}
}
