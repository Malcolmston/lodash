// Package lodash is a comprehensive, type-safe functional utility library for
// Go, inspired by the JavaScript library lodash. It is implemented entirely
// with the Go standard library and leans on Go 1.24 generics so that helpers
// stay type-safe without external dependencies.
//
// All functions are pure and non-mutating unless explicitly documented
// otherwise: they return new slices or maps rather than modifying their
// arguments. Any randomness is injected via a *math/rand.Rand so that behavior
// is fully deterministic and testable.
//
// # Collections and slices
//
// Transform, query and reshape slices:
//
//	Map, MapI, Filter, Reject, Reduce, ReduceRight, ForEach, ForEachI,
//	ForEachRight, Find, FindLast, FindIndex, FindLastIndex, Every, Some, None,
//	Includes, IndexOf, LastIndexOf, IndexOfFrom, LastIndexOfFrom, GroupBy,
//	KeyBy, CountBy, Partition, Uniq, UniqBy, SortedUniq, SortedUniqBy, Chunk,
//	Flatten, FlattenDeep, FlattenDepth, FlatMap, Compact, Reverse, Zip, ZipWith,
//	Unzip, UnzipWith, ZipObject, Difference, DifferenceBy, DifferenceWith,
//	Intersection, IntersectionBy, IntersectionWith, Union, UnionBy, UnionWith,
//	Xor, XorBy, XorWith, Without, Pull, PullAll, PullAllBy, PullAt, Remove,
//	Take, TakeRight, TakeWhile, TakeRightWhile, Drop, DropRight, DropWhile,
//	DropRightWhile, Head, First, Last, Nth, Initial, Tail, Slice, Join, Sample,
//	SampleN, Shuffle, SortBy, OrderBy, SortedIndex, SortedIndexBy, SortedIndexOf,
//	SortedLastIndex, SortedLastIndexBy, SortedLastIndexOf, Concat, Fill.
//
// # Math and numbers
//
// Aggregate, constrain and compute over numeric data:
//
//	Sum, SumBy, Mean, MeanBy, Min, MinBy, Max, MaxBy, Add, Subtract, Multiply,
//	Divide, Clamp, Round, Floor, Ceil, Random, RandomFloat, Range, RangeStep,
//	RangeRight, InRange.
//
// # Objects and maps
//
// Inspect and reshape maps, including nested map[string]any access with dotted
// and bracket paths:
//
//	Keys, SortedKeys, Values, Entries, FromEntries, ToPairs, FromPairs,
//	ZipObject, Pick, PickBy, Omit, OmitBy, MapKeys, MapValues, Invert, InvertBy,
//	Merge, MergeWith, Assign, AssignWith, Defaults, DefaultsDeep, Transform,
//	ForOwn, FindKey, Size, Get, GetOr, Has, At, Result, Set, Unset, Update,
//	ToPath.
//
// # Strings
//
// Case conversion, formatting, escaping and templating:
//
//	Words, CamelCase, PascalCase, SnakeCase, KebabCase, StartCase, LowerCase,
//	UpperCase, ToLower, ToUpper, Capitalize, UpperFirst, LowerFirst, Pad,
//	PadStart, PadEnd, Truncate, Repeat, Deburr, Trim, TrimStart, TrimEnd,
//	Escape, Unescape, EscapeRegExp, Replace, Split, StartsWith, EndsWith,
//	Template.
//
// # Lang and value helpers
//
// Clone, compare and interrogate arbitrary values:
//
//	Clone, CloneDeep, IsEqual, Eq, IsEmpty, IsMatch, IsNil, IsPlainObject,
//	IsString, IsBool, IsNumber, IsInteger, IsSlice, IsMap, IsError,
//	IsObjectLike, IsNaN, IsFinite, CastArray, ToArray, DefaultTo, Gt, Gte, Lt,
//	Lte, ToNumber, ToInteger, ToFinite, ToSafeInteger, ToString, Conforms,
//	ConformsTo.
//
// # Function combinators
//
// Compose, adapt and constrain functions:
//
//	Curry, Curry3, Partial, PartialRight, Flow, FlowRight, Compose, Cond, Over,
//	OverEvery, OverSome, OverArgs, Negate, Ary, Unary, Flip, Rearg, Spread,
//	NthArg, Wrap, Once, Memoize, MemoizeBy, NewDebouncer, NewThrottler, After,
//	Before.
//
// # Utilities and sequences
//
// General-purpose helpers and a lazy value wrapper for method chaining:
//
//	Identity, Constant, Noop, Times, UniqueID, Property, PropertyOf, Matches,
//	Attempt, StubArray, StubObject, StubString, StubTrue, StubFalse, Chain
//	(with Value, Thru and Tap).
package lodash
