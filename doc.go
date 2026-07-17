// Package lodash is a comprehensive, type-safe functional utility library for
// Go, inspired by the JavaScript library lodash. It is implemented entirely
// with the Go standard library and leans on Go 1.24 generics so that helpers
// stay type-safe without reflection.
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
//	Find, FindLast, FindIndex, FindLastIndex, Every, Some, Includes,
//	IndexOf, LastIndexOf, GroupBy, KeyBy, CountBy, Partition, Uniq, UniqBy,
//	Chunk, Flatten, FlattenDeep, Compact, Reverse, Zip, Unzip, Difference,
//	Intersection, Union, Without, Take, TakeRight, Drop, DropRight, Sample,
//	SampleN, Shuffle, SortBy, OrderBy, Concat, Fill.
//
// # Math and numbers
//
// Aggregate and constrain numeric data:
//
//	Sum, SumBy, Mean, MeanBy, Min, MinBy, Max, MaxBy, Clamp, Range,
//	RangeStep, InRange.
//
// # Objects and maps
//
// Inspect and reshape maps, including nested map[string]any access:
//
//	Keys, SortedKeys, Values, Entries, FromEntries, Pick, PickBy, Omit,
//	OmitBy, MapKeys, MapValues, Invert, Merge, Assign, Get, Has.
//
// # Strings
//
// Case conversion and formatting helpers:
//
//	Words, CamelCase, PascalCase, SnakeCase, KebabCase, StartCase,
//	Capitalize, UpperFirst, LowerFirst, Pad, PadStart, PadEnd, Truncate,
//	Repeat, Deburr, Trim, TrimStart, TrimEnd.
//
// # Functions
//
// Higher-order function wrappers. Once and Memoize are fully deterministic;
// Debouncer and Throttler manage real timers and clocks and document their
// concurrency behavior on each type:
//
//	Once, Memoize, MemoizeBy, NewDebouncer, NewThrottler, After, Before.
package lodash
