# lodash

A comprehensive, type-safe **functional utility library for Go**, inspired by
[lodash](https://lodash.com/). Built entirely on the Go standard library and
Go 1.24 generics — no third-party dependencies, no reflection on hot paths, no
cgo.

All functions are **pure and non-mutating** unless documented otherwise: they
return new slices or maps rather than modifying their arguments. Any randomness
is injected through a `*math/rand.Rand`, so behavior is fully deterministic and
testable.

## Install

```sh
go get github.com/malcolmston/lodash
```

Requires Go 1.24 or newer.

## Quick start

```go
package main

import (
	"fmt"

	lodash "github.com/malcolmston/lodash"
)

func main() {
	nums := []int{1, 2, 3, 4, 5, 6}

	evens := lodash.Filter(nums, func(n int) bool { return n%2 == 0 })
	doubled := lodash.Map(evens, func(n int) int { return n * 2 })
	total := lodash.Sum(doubled)

	fmt.Println(evens)   // [2 4 6]
	fmt.Println(doubled) // [4 8 12]
	fmt.Println(total)   // 24

	fmt.Println(lodash.CamelCase("Foo Bar-baz")) // fooBarBaz
	fmt.Println(lodash.SnakeCase("fooBarBaz"))   // foo_bar_baz
}
```

## Function reference

### Collections & slices

`Map`, `MapI`, `Filter`, `Reject`, `Reduce`, `ReduceRight`, `ForEach`,
`ForEachI`, `Find`, `FindLast`, `FindIndex`, `FindLastIndex`, `Every`, `Some`,
`Includes`, `IndexOf`, `LastIndexOf`, `GroupBy`, `KeyBy`, `CountBy`,
`Partition`, `Uniq`, `UniqBy`, `Chunk`, `Flatten`, `FlattenDeep`, `Compact`,
`Reverse`, `Zip`, `Unzip`, `Difference`, `Intersection`, `Union`, `Without`,
`Take`, `TakeRight`, `Drop`, `DropRight`, `Sample`, `SampleN`, `Shuffle`,
`SortBy`, `OrderBy`, `Concat`, `Fill`

### Math & numbers

`Sum`, `SumBy`, `Mean`, `MeanBy`, `Min`, `MinBy`, `Max`, `MaxBy`, `Clamp`,
`Range`, `RangeStep`, `InRange`

### Objects & maps

`Keys`, `SortedKeys`, `Values`, `Entries`, `FromEntries`, `Pick`, `PickBy`,
`Omit`, `OmitBy`, `MapKeys`, `MapValues`, `Invert`, `Merge`, `Assign`, `Get`,
`Has`

### Strings

`Words`, `CamelCase`, `PascalCase`, `SnakeCase`, `KebabCase`, `StartCase`,
`Capitalize`, `UpperFirst`, `LowerFirst`, `Pad`, `PadStart`, `PadEnd`,
`Truncate`, `Repeat`, `Deburr`, `Trim`, `TrimStart`, `TrimEnd`

### Functions

`Once`, `Memoize`, `MemoizeBy`, `NewDebouncer` (`Debouncer`), `NewThrottler`
(`Throttler`), `After`, `Before`

## Notes on randomness & timers

- `Sample`, `SampleN` and `Shuffle` take a `*math/rand.Rand` so results are
  deterministic when seeded — ideal for reproducible tests.
- `Debouncer` schedules work with a `time.Timer` on a background goroutine; call
  `Cancel` to abort a pending invocation.
- `Throttler` guards execution with an injectable clock (`now` field) so its
  timing logic can be tested without real sleeps.

## Development

```sh
go build ./...
go vet ./...
go test ./... -cover
gofmt -l .
golangci-lint run ./...
```

## License

MIT
