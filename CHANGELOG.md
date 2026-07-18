# Changelog

All notable changes to this project are documented in this file. The format is
based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/) and this
project adheres to [Semantic Versioning](https://semver.org/).

## [Unreleased]

- Upstream-parity: synced lodash v4.17.21 test vectors into `*_parity_test.go`; fixed `StartCase` to preserve interior letter case, made the case compounders (`CamelCase`/`PascalCase`/`SnakeCase`/`KebabCase`/`StartCase`/`LowerCase`/`UpperCase`) deburr letters and strip contraction apostrophes like lodash, and rewrote `Deburr` to cover the full Latin-1/Latin Extended-A tables, ligatures, and combining marks.

## [0.3.0] - 2026-07-18

### Added

Expanded the library from **269** to **354** exported identifiers (+85, a ~32%
increase), closing further gaps toward lodash. All additions are
standard-library-only, generics-based, fully documented, gofmt-clean and
covered by deterministic known-answer tests (plus benchmarks for the
performance-sensitive routines).

- **Lang predicates & conversions** (`predicates.go`): `IsFunction`,
  `IsRegExp`, `IsDate`, `IsBuffer`, `IsSafeInteger`, `IsLength`, `IsNull`,
  `IsUndefined`, `IsObject`, `IsArrayLike`, `IsArrayLikeObject`, `ToLength`,
  `IsEqualWith`, `IsMatchWith`.
- **Object-as-collection helpers** (`collectionmap.go`), mirroring lodash's
  collection methods that also accept objects: `EveryMap`, `SomeMap`,
  `NoneMap`, `FindMap`, `FilterMap`, `RejectMap`, `MapToSlice`, `ReduceMap`,
  `PartitionMap`, `GroupByMap`, `IncludesValue`, `MinByMap`, `MaxByMap`.
- **Array gaps** (`arraymore.go`): `UniqWith`, `PullAllWith`, `ZipObjectDeep`,
  `FlatMapDeep`, `FlatMapDepth`, `FillRange`.
- **Function combinators** (`functionmore.go`): `CurryRight`, `CurryRight3`,
  `Curry4`, `Rest`, `Delay`, `Defer`.
- **Deep-object writes** (`objectmore.go`): `SetWith`, `UpdateWith`.
- **Utilities & strings** (`utilmore.go`, `stringsmore.go`): `MatchesProperty`,
  `ParseInt`, `ParseFloat`.
- **Fluent chaining** (`chainslice.go`, `chainset.go`): `SliceSeq`/`ChainSlice`
  for same-type slice pipelines (`Filter`, `Reject`, `Reverse`, `Take`,
  `TakeRight`, `TakeWhile`, `Drop`, `DropRight`, `DropWhile`, `Tail`, `Initial`,
  `Slice`, `Concat`, `ForEach`, `Tap`, `Shuffle`, `Head`, `Last`, `Nth`,
  `Sample`, `Size`, `IsEmpty`, `Join`, `Chunk`) and `SetSeq`/`ChainSet` for
  comparable-element set operations (`Uniq`, `Compact`, `Without`, `Union`,
  `Intersection`, `Difference`, `Filter`, `Reverse`, `Includes`, `IndexOf`,
  `Size`).

## [0.2.0] - 2026-07-17

### Added

Expanded the library from **98** to **250** exported functions (~80% parity
with lodash), all documented, gofmt-clean and covered by deterministic tests
(95% statement coverage). New standard-library-only, generics-based helpers:

- **Lang & value** (`lang.go`): `Clone`, `CloneDeep`, `IsEqual`, `Eq`,
  `IsEmpty`, `IsMatch`, `IsNil`, `IsPlainObject`, `IsString`, `IsBool`,
  `IsNumber`, `IsInteger`, `IsSlice`, `IsMap`, `IsError`, `IsObjectLike`,
  `IsNaN`, `IsFinite`, `CastArray`, `ToArray`, `DefaultTo`, `Gt`, `Gte`, `Lt`,
  `Lte`, `ToNumber`, `ToInteger`, `ToFinite`, `ToSafeInteger`, `ToString`,
  `Conforms`, `ConformsTo`.
- **Deep-path objects** (`path.go`, `object2.go`): `Get`/`Has` upgraded to
  dotted **and** bracket paths plus `GetOr`, `At`, `Result`, `Set`, `Unset`,
  `Update`, `ToPath`; `MergeWith`, `AssignWith`, `Defaults`, `DefaultsDeep`,
  `ToPairs`, `FromPairs`, `ZipObject`, `FindKey`, `InvertBy`, `Transform`,
  `ForOwn`, `Size`.
- **Function combinators** (`combinators.go`): `Curry`, `Curry3`, `Partial`,
  `PartialRight`, `Flow`, `FlowRight`, `Compose`, `Cond`, `Over`, `OverEvery`,
  `OverSome`, `OverArgs`, `Negate`, `Ary`, `Unary`, `Flip`, `Rearg`, `Spread`,
  `NthArg`, `Wrap`.
- **Array set-ops & variants** (`setops.go`): `Head`, `First`, `Last`, `Nth`,
  `Initial`, `Tail`, `Slice`, `Join`, `DifferenceBy`, `DifferenceWith`,
  `IntersectionBy`, `IntersectionWith`, `UnionBy`, `UnionWith`, `Xor`, `XorBy`,
  `XorWith`, `SortedUniq`, `SortedUniqBy`, `PullAll`, `PullAllBy`, `Pull`,
  `PullAt`, `Remove`, `FlatMap`, `FlattenDepth`, `TakeWhile`, `TakeRightWhile`,
  `DropWhile`, `DropRightWhile`, `ZipWith`, `UnzipWith`, `SortedIndex(By/Of)`,
  `SortedLastIndex(By/Of)`, `IndexOfFrom`, `LastIndexOfFrom`, `None`,
  `ForEachRight`.
- **Strings** (`strings2.go`): `Escape`, `Unescape`, `EscapeRegExp`, `Template`
  (interpolate / escape / evaluate delimiters), `LowerCase`, `UpperCase`,
  `ToLower`, `ToUpper`, `StartsWith`, `EndsWith`, `Replace`, `Split`.
- **Number & math** (`number.go`): `Random` and `RandomFloat` (seeded via
  `*math/rand.Rand`), `Round`, `Floor`, `Ceil` (precision-aware), `Add`,
  `Subtract`, `Multiply`, `Divide`, `RangeRight`.
- **Sequences** (`seq.go`): `Chain` returning a lazy `Seq[T]` wrapper with a
  `Value()` terminator plus `Thru`/`Tap` (methods and type-changing functions).
- **Utilities** (`util.go`): `Identity`, `Constant`, `Noop`, `Times`,
  `UniqueID`, `Property`, `PropertyOf`, `Matches`, `Attempt`, `StubArray`,
  `StubObject`, `StubString`, `StubTrue`, `StubFalse`.

### Changed

- `Get` and `Has` now parse bracket notation (e.g. `a.b[0].c`) and can traverse
  `[]any` slices in addition to nested `map[string]any` values. Their existing
  dotted-path behaviour is unchanged.

## [0.1.0]

### Added

- Initial release with 98 collection, math, object, string and function
  helpers, plus a standard-library documentation generator and React docs site.
