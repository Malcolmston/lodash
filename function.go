package lodash

import (
	"sync"
	"time"
)

// Once returns a function that invokes fn at most once. Subsequent calls return
// the value produced by the first invocation. It is safe for concurrent use.
func Once[R any](fn func() R) func() R {
	var (
		once   sync.Once
		result R
	)
	return func() R {
		once.Do(func() {
			result = fn()
		})
		return result
	}
}

// Memoize returns a function that caches the result of fn for each distinct
// argument. The cache is safe for concurrent use.
func Memoize[K comparable, R any](fn func(K) R) func(K) R {
	var mu sync.Mutex
	cache := make(map[K]R)
	return func(k K) R {
		mu.Lock()
		defer mu.Unlock()
		if v, ok := cache[k]; ok {
			return v
		}
		v := fn(k)
		cache[k] = v
		return v
	}
}

// MemoizeBy is like Memoize but derives the cache key from arg via keyFn,
// allowing memoization over non-comparable argument types.
func MemoizeBy[A any, K comparable, R any](fn func(A) R, keyFn func(A) K) func(A) R {
	var mu sync.Mutex
	cache := make(map[K]R)
	return func(a A) R {
		k := keyFn(a)
		mu.Lock()
		defer mu.Unlock()
		if v, ok := cache[k]; ok {
			return v
		}
		v := fn(a)
		cache[k] = v
		return v
	}
}

// Debouncer wraps a function so that it only runs after calls have stopped for a
// specified quiet period. Each call to Call resets the timer; the wrapped
// function runs once, wait after the final call. Use Cancel to abort a pending
// invocation. Debouncer is safe for concurrent use and schedules the call on a
// separate goroutine via a time.Timer.
type Debouncer struct {
	mu    sync.Mutex
	wait  time.Duration
	timer *time.Timer
	fn    func()
}

// NewDebouncer creates a Debouncer that will invoke fn once calls have been
// quiet for wait.
func NewDebouncer(wait time.Duration, fn func()) *Debouncer {
	return &Debouncer{wait: wait, fn: fn}
}

// Call schedules the wrapped function to run after the quiet period, resetting
// any previously scheduled invocation.
func (d *Debouncer) Call() {
	d.mu.Lock()
	defer d.mu.Unlock()
	if d.timer != nil {
		d.timer.Stop()
	}
	d.timer = time.AfterFunc(d.wait, d.fn)
}

// Cancel stops any pending invocation. It reports whether a pending call was
// cancelled.
func (d *Debouncer) Cancel() bool {
	d.mu.Lock()
	defer d.mu.Unlock()
	if d.timer != nil {
		stopped := d.timer.Stop()
		d.timer = nil
		return stopped
	}
	return false
}

// Throttler wraps a function so that it runs at most once per interval. The
// first call runs immediately; subsequent calls within the interval are
// suppressed. Throttler is safe for concurrent use.
type Throttler struct {
	mu       sync.Mutex
	interval time.Duration
	last     time.Time
	fn       func()
	now      func() time.Time // injectable clock for deterministic tests
}

// NewThrottler creates a Throttler that permits fn to run at most once per
// interval.
func NewThrottler(interval time.Duration, fn func()) *Throttler {
	return &Throttler{interval: interval, fn: fn, now: time.Now}
}

// Call runs the wrapped function if at least interval has elapsed since the last
// successful run (or if this is the first call). It reports whether the function
// was invoked.
func (t *Throttler) Call() bool {
	t.mu.Lock()
	now := t.now()
	if !t.last.IsZero() && now.Sub(t.last) < t.interval {
		t.mu.Unlock()
		return false
	}
	t.last = now
	fn := t.fn
	t.mu.Unlock()
	fn()
	return true
}

// After returns a function that only invokes fn beginning with the nth call. The
// first n-1 calls do nothing. It is safe for concurrent use.
func After[R any](n int, fn func() R) func() (R, bool) {
	var (
		mu    sync.Mutex
		count int
	)
	return func() (R, bool) {
		mu.Lock()
		count++
		ready := count >= n
		mu.Unlock()
		if ready {
			return fn(), true
		}
		var zero R
		return zero, false
	}
}

// Before returns a function that invokes fn only for the first n-1 calls. From
// the nth call onward it returns the result of the last successful invocation.
// It is safe for concurrent use.
func Before[R any](n int, fn func() R) func() R {
	var (
		mu    sync.Mutex
		count int
		last  R
	)
	return func() R {
		mu.Lock()
		defer mu.Unlock()
		if count < n-1 {
			count++
			last = fn()
		}
		return last
	}
}
