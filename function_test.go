package lodash

import (
	"sync"
	"testing"
	"time"
)

func TestOnce(t *testing.T) {
	calls := 0
	f := Once(func() int {
		calls++
		return 42
	})
	for i := 0; i < 5; i++ {
		if f() != 42 {
			t.Fatal("once value")
		}
	}
	if calls != 1 {
		t.Fatalf("once called %d times", calls)
	}
}

func TestOnceConcurrent(t *testing.T) {
	var calls int
	var mu sync.Mutex
	f := Once(func() int {
		mu.Lock()
		calls++
		mu.Unlock()
		return 1
	})
	var wg sync.WaitGroup
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			f()
		}()
	}
	wg.Wait()
	if calls != 1 {
		t.Fatalf("once concurrent called %d times", calls)
	}
}

func TestMemoize(t *testing.T) {
	calls := 0
	square := Memoize(func(x int) int {
		calls++
		return x * x
	})
	first := square(4)
	cached := square(4) // served from cache
	if first != 16 || cached != 16 || square(5) != 25 {
		t.Fatal("memoize values")
	}
	if calls != 2 {
		t.Fatalf("memoize calls = %d", calls)
	}
}

func TestMemoizeBy(t *testing.T) {
	calls := 0
	f := MemoizeBy(func(p person) int {
		calls++
		return p.Age * 2
	}, func(p person) string { return p.Name })
	a := person{"Ada", 30}
	if f(a) != 60 || f(person{"Ada", 99}) != 60 {
		t.Fatal("memoizeby uses key")
	}
	if calls != 1 {
		t.Fatalf("memoizeby calls = %d", calls)
	}
}

func TestAfterBefore(t *testing.T) {
	f := After(3, func() string { return "go" })
	if _, ok := f(); ok {
		t.Fatal("after 1")
	}
	if _, ok := f(); ok {
		t.Fatal("after 2")
	}
	if v, ok := f(); !ok || v != "go" {
		t.Fatal("after 3")
	}

	calls := 0
	b := Before(3, func() int {
		calls++
		return calls
	})
	if b() != 1 || b() != 2 || b() != 2 || b() != 2 {
		t.Fatal("before values")
	}
	if calls != 2 {
		t.Fatalf("before calls = %d", calls)
	}
}

func TestThrottleDeterministic(t *testing.T) {
	calls := 0
	th := NewThrottler(100*time.Millisecond, func() { calls++ })
	// Inject a controllable clock.
	now := time.Unix(0, 0)
	th.now = func() time.Time { return now }

	if !th.Call() { // first call runs
		t.Fatal("throttle first")
	}
	if th.Call() { // within interval, suppressed
		t.Fatal("throttle suppress")
	}
	now = now.Add(150 * time.Millisecond)
	if !th.Call() { // interval elapsed, runs again
		t.Fatal("throttle after interval")
	}
	if calls != 2 {
		t.Fatalf("throttle calls = %d", calls)
	}
}

func TestDebounce(t *testing.T) {
	var mu sync.Mutex
	calls := 0
	d := NewDebouncer(30*time.Millisecond, func() {
		mu.Lock()
		calls++
		mu.Unlock()
	})
	// Rapid burst: only the final call should fire.
	for i := 0; i < 5; i++ {
		d.Call()
		time.Sleep(5 * time.Millisecond)
	}
	time.Sleep(60 * time.Millisecond)
	mu.Lock()
	got := calls
	mu.Unlock()
	if got != 1 {
		t.Fatalf("debounce fired %d times, want 1", got)
	}
}

func TestDebounceCancel(t *testing.T) {
	calls := 0
	d := NewDebouncer(30*time.Millisecond, func() { calls++ })
	d.Call()
	if !d.Cancel() {
		t.Fatal("cancel should report pending")
	}
	time.Sleep(50 * time.Millisecond)
	if calls != 0 {
		t.Fatalf("debounce cancel fired %d", calls)
	}
	if d.Cancel() {
		t.Fatal("cancel with nothing pending")
	}
}
