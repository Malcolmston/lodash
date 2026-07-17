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

Over **250 functions** across seven categories — roughly 80% of lodash's API,
adapted idiomatically to Go generics.

### Collections & slices

`Map`, `MapI`, `Filter`, `Reject`, `Reduce`, `ReduceRight`, `ForEach`,
`ForEachI`, `ForEachRight`, `Find`, `FindLast`, `FindIndex`, `FindLastIndex`,
`Every`, `Some`, `None`, `Includes`, `IndexOf`, `LastIndexOf`, `IndexOfFrom`,
`LastIndexOfFrom`, `GroupBy`, `KeyBy`, `CountBy`, `Partition`, `Uniq`, `UniqBy`,
`SortedUniq`, `SortedUniqBy`, `Chunk`, `Flatten`, `FlattenDeep`, `FlattenDepth`,
`FlatMap`, `Compact`, `Reverse`, `Zip`, `ZipWith`, `Unzip`, `UnzipWith`,
`ZipObject`, `Difference`, `DifferenceBy`, `DifferenceWith`, `Intersection`,
`IntersectionBy`, `IntersectionWith`, `Union`, `UnionBy`, `UnionWith`, `Xor`,
`XorBy`, `XorWith`, `Without`, `Pull`, `PullAll`, `PullAllBy`, `PullAt`,
`Remove`, `Take`, `TakeRight`, `TakeWhile`, `TakeRightWhile`, `Drop`,
`DropRight`, `DropWhile`, `DropRightWhile`, `Head`, `First`, `Last`, `Nth`,
`Initial`, `Tail`, `Slice`, `Join`, `Sample`, `SampleN`, `Shuffle`, `SortBy`,
`OrderBy`, `SortedIndex`, `SortedIndexBy`, `SortedIndexOf`, `SortedLastIndex`,
`SortedLastIndexBy`, `SortedLastIndexOf`, `Concat`, `Fill`

### Math & numbers

`Sum`, `SumBy`, `Mean`, `MeanBy`, `Min`, `MinBy`, `Max`, `MaxBy`, `Add`,
`Subtract`, `Multiply`, `Divide`, `Clamp`, `Round`, `Floor`, `Ceil`, `Random`,
`RandomFloat`, `Range`, `RangeStep`, `RangeRight`, `InRange`

### Objects & maps

`Keys`, `SortedKeys`, `Values`, `Entries`, `FromEntries`, `ToPairs`,
`FromPairs`, `ZipObject`, `Pick`, `PickBy`, `Omit`, `OmitBy`, `MapKeys`,
`MapValues`, `Invert`, `InvertBy`, `Merge`, `MergeWith`, `Assign`, `AssignWith`,
`Defaults`, `DefaultsDeep`, `Transform`, `ForOwn`, `FindKey`, `Size`, `Get`,
`GetOr`, `Has`, `At`, `Result`, `Set`, `Unset`, `Update`, `ToPath`

Deep-path helpers (`Get`, `Set`, `Unset`, `Update`, `Has`, `At`, `Result`)
accept dotted **and** bracket notation over `map[string]any`, e.g.
`Get(obj, "a.b[0].c")`.

### Strings

`Words`, `CamelCase`, `PascalCase`, `SnakeCase`, `KebabCase`, `StartCase`,
`LowerCase`, `UpperCase`, `ToLower`, `ToUpper`, `Capitalize`, `UpperFirst`,
`LowerFirst`, `Pad`, `PadStart`, `PadEnd`, `Truncate`, `Repeat`, `Deburr`,
`Trim`, `TrimStart`, `TrimEnd`, `Escape`, `Unescape`, `EscapeRegExp`, `Replace`,
`Split`, `StartsWith`, `EndsWith`, `Template`

### Lang & value

`Clone`, `CloneDeep`, `IsEqual`, `Eq`, `IsEmpty`, `IsMatch`, `IsNil`,
`IsPlainObject`, `IsString`, `IsBool`, `IsNumber`, `IsInteger`, `IsSlice`,
`IsMap`, `IsError`, `IsObjectLike`, `IsNaN`, `IsFinite`, `CastArray`, `ToArray`,
`DefaultTo`, `Gt`, `Gte`, `Lt`, `Lte`, `ToNumber`, `ToInteger`, `ToFinite`,
`ToSafeInteger`, `ToString`, `Conforms`, `ConformsTo`

### Function combinators

`Curry`, `Curry3`, `Partial`, `PartialRight`, `Flow`, `FlowRight`, `Compose`,
`Cond`, `Over`, `OverEvery`, `OverSome`, `OverArgs`, `Negate`, `Ary`, `Unary`,
`Flip`, `Rearg`, `Spread`, `NthArg`, `Wrap`, `Once`, `Memoize`, `MemoizeBy`,
`NewDebouncer` (`Debouncer`), `NewThrottler` (`Throttler`), `After`, `Before`

### Utilities & sequences

`Identity`, `Constant`, `Noop`, `Times`, `UniqueID`, `Property`, `PropertyOf`,
`Matches`, `Attempt`, `StubArray`, `StubObject`, `StubString`, `StubTrue`,
`StubFalse`, `Chain` (a lazy `Seq` wrapper with `Value`, `Thru` and `Tap`)

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
