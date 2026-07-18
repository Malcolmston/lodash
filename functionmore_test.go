package lodash

import (
	"sync"
	"testing"
	"time"
)

func TestCurryRight(t *testing.T) {
	sub := func(a, b int) int { return a - b }
	// CurryRight takes the last argument (b) first.
	if got := CurryRight(sub)(3)(10); got != 7 {
		t.Errorf("CurryRight = %d, want 7", got)
	}
}

func TestCurryRight3(t *testing.T) {
	f := func(a, b, c string) string { return a + b + c }
	// Supplies c, then b, then a.
	if got := CurryRight3(f)("c")("b")("a"); got != "abc" {
		t.Errorf("CurryRight3 = %q, want abc", got)
	}
}

func TestCurry4(t *testing.T) {
	f := func(a, b, c, d int) int { return a*1000 + b*100 + c*10 + d }
	if got := Curry4(f)(1)(2)(3)(4); got != 1234 {
		t.Errorf("Curry4 = %d, want 1234", got)
	}
}

func TestRest(t *testing.T) {
	sum := func(xs []int) int {
		total := 0
		for _, x := range xs {
			total += x
		}
		return total
	}
	variadic := Rest(sum)
	if got := variadic(1, 2, 3, 4); got != 10 {
		t.Errorf("Rest = %d, want 10", got)
	}
	if got := variadic(); got != 0 {
		t.Errorf("Rest() = %d, want 0", got)
	}
}

func TestDelay(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)
	fired := false
	Delay(func() { fired = true; wg.Done() }, time.Millisecond)
	wg.Wait()
	if !fired {
		t.Error("Delay did not fire")
	}
}

func TestDelayCancel(t *testing.T) {
	fired := make(chan struct{}, 1)
	timer := Delay(func() { fired <- struct{}{} }, 50*time.Millisecond)
	if !timer.Stop() {
		t.Skip("timer already fired; timing-dependent")
	}
	select {
	case <-fired:
		t.Error("Delay fired after cancel")
	case <-time.After(20 * time.Millisecond):
		// expected: nothing fired
	}
}

func TestDefer(t *testing.T) {
	done := make(chan int, 1)
	Defer(func() { done <- 42 })
	select {
	case v := <-done:
		if v != 42 {
			t.Errorf("Defer value = %d", v)
		}
	case <-time.After(time.Second):
		t.Error("Defer did not run")
	}
}
