package lodash_test

import (
	"testing"

	lodash "github.com/malcolmston/lodash"
)

func TestNegate(t *testing.T) {
	isEven := func(n int) bool { return n%2 == 0 }
	isOdd := lodash.Negate(isEven)
	if isOdd(2) || !isOdd(3) {
		t.Errorf("Negate")
	}
}

func TestCurryPartial(t *testing.T) {
	add := func(a, b int) int { return a + b }
	if lodash.Curry(add)(2)(3) != 5 {
		t.Errorf("Curry")
	}
	add3 := func(a, b, c int) int { return a + b + c }
	if lodash.Curry3(add3)(1)(2)(3) != 6 {
		t.Errorf("Curry3")
	}
	if lodash.Partial(add, 10)(5) != 15 {
		t.Errorf("Partial")
	}
	sub := func(a, b int) int { return a - b }
	if lodash.PartialRight(sub, 3)(10) != 7 {
		t.Errorf("PartialRight")
	}
	if lodash.Flip(sub)(3, 10) != 7 {
		t.Errorf("Flip")
	}
}

func TestFlow(t *testing.T) {
	inc := func(n int) int { return n + 1 }
	double := func(n int) int { return n * 2 }
	if lodash.Flow(inc, double)(3) != 8 { // (3+1)*2
		t.Errorf("Flow")
	}
	if lodash.FlowRight(inc, double)(3) != 7 { // (3*2)+1
		t.Errorf("FlowRight")
	}
	if lodash.Compose(inc, double)(3) != 7 {
		t.Errorf("Compose")
	}
	if lodash.Flow[int]()(5) != 5 {
		t.Errorf("Flow empty identity")
	}
}

func TestCond(t *testing.T) {
	classify := lodash.Cond(
		lodash.CondPair[int, string]{Pred: func(n int) bool { return n < 0 }, Result: func(int) string { return "neg" }},
		lodash.CondPair[int, string]{Pred: func(n int) bool { return n == 0 }, Result: func(int) string { return "zero" }},
	)
	if v, _ := classify(-5); v != "neg" {
		t.Errorf("Cond neg")
	}
	if v, ok := classify(0); v != "zero" || !ok {
		t.Errorf("Cond zero")
	}
	if _, ok := classify(5); ok {
		t.Errorf("Cond no match should be false")
	}
}

func TestOver(t *testing.T) {
	fns := lodash.Over(func(n int) int { return n + 1 }, func(n int) int { return n * 2 })
	if !lodash.IsEqual(fns(3), []int{4, 6}) {
		t.Errorf("Over")
	}
	pos := func(n int) bool { return n > 0 }
	even := func(n int) bool { return n%2 == 0 }
	if !lodash.OverEvery(pos, even)(4) || lodash.OverEvery(pos, even)(3) {
		t.Errorf("OverEvery")
	}
	if !lodash.OverSome(pos, even)(3) || lodash.OverSome(pos, even)(-3) {
		t.Errorf("OverSome")
	}
}

func TestVariadicCombinators(t *testing.T) {
	sum := func(ns ...int) int {
		total := 0
		for _, n := range ns {
			total += n
		}
		return total
	}
	if lodash.Ary(sum, 2)(1, 2, 3, 4) != 3 {
		t.Errorf("Ary")
	}
	if lodash.Unary(sum)(5) != 5 {
		t.Errorf("Unary")
	}
	if lodash.Spread(sum)([]int{1, 2, 3}) != 6 {
		t.Errorf("Spread")
	}
	// Rearg reorders: call with args reordered [2,0]
	first := func(ns ...int) int {
		if len(ns) > 0 {
			return ns[0]
		}
		return -1
	}
	if lodash.Rearg(first, 2, 0)(10, 20, 30) != 30 {
		t.Errorf("Rearg")
	}
	doubled := lodash.OverArgs(sum, func(n int) int { return n * 10 })
	if doubled(1, 2) != 12 { // 10 + 2
		t.Errorf("OverArgs = %d", doubled(1, 2))
	}
	if lodash.NthArg[int](1)(7, 8, 9) != 8 {
		t.Errorf("NthArg")
	}
	if lodash.NthArg[int](-1)(7, 8, 9) != 9 {
		t.Errorf("NthArg neg")
	}
	if lodash.NthArg[int](9)(1) != 0 {
		t.Errorf("NthArg out of range")
	}
}

func TestWrap(t *testing.T) {
	wrapped := lodash.Wrap("!", func(suffix string, name string) string {
		return "Hi " + name + suffix
	})
	if wrapped("Bob") != "Hi Bob!" {
		t.Errorf("Wrap")
	}
}
