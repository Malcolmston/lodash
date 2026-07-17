// Library content for the lodash documentation site. Mirrors the shape used by
// the malcolmston/go landing site's data.ts so the sibling sites stay in sync.
export interface Lib {
  id: string; name: string; icon: string; accent: string; pkg: string; node: string;
  repo: string; docs: string; tagline: string; blurb: string; tags: string[];
  features: string[]; node_code: string; go_code: string; integrate: string;
}

export const NODE_ACCENT = '#8cc84b';

export const LODASH: Lib = {
  id:"lodash", name:"lodash", icon:'<i class="fa-solid fa-toolbox"></i>', accent:"#f2a33c",
  pkg:"github.com/malcolmston/lodash", node:"lodash/lodash",
  repo:"https://github.com/malcolmston/lodash", docs:"https://malcolmston.github.io/lodash/",
  tagline:"Type-safe, lodash-style functional utilities for Go.",
  blurb:"A comprehensive, type-safe functional utility library for Go, inspired by JavaScript's lodash and built "+
    "entirely on the standard library and Go 1.24 generics — no third-party dependencies, no cgo, no reflection "+
    "on hot paths. Roughly 98 helpers span collections and slices (Map/Filter/Reduce/GroupBy/Uniq/Chunk/"+
    "Difference/Union/SortBy), math (Sum/Mean/Min/Max/Clamp/Range), objects and nested maps (Keys/Pick/Omit/"+
    "Invert/Merge/Get), string casing (CamelCase/SnakeCase/KebabCase/Deburr) and function wrappers (Once/Memoize/"+
    "Debouncer/Throttler). Every function is pure and non-mutating unless documented otherwise: helpers return new "+
    "slices or maps rather than modifying their arguments. Any randomness is injected through a *math/rand.Rand, so "+
    "Sample, SampleN and Shuffle stay fully deterministic and testable, while Debouncer and Throttler take "+
    "injectable timers and clocks. The import path is github.com/malcolmston/lodash and the package is named lodash.",
  tags:["generics","pure/non-mutating","zero dependencies","collections","math","objects","strings","functions"],
  features:[
    "Collections — <code>Map</code>, <code>Filter</code>, <code>Reduce</code>, <code>GroupBy</code>, <code>KeyBy</code>, <code>CountBy</code>, <code>Partition</code>, <code>ForEach</code>",
    "Querying — <code>Find</code>, <code>FindIndex</code>, <code>Every</code>, <code>Some</code>, <code>Includes</code>, <code>IndexOf</code>, <code>LastIndexOf</code>",
    "Set &amp; shape ops — <code>Uniq</code>, <code>UniqBy</code>, <code>Chunk</code>, <code>Flatten</code>, <code>FlattenDeep</code>, <code>Difference</code>, <code>Intersection</code>, <code>Union</code>, <code>Zip</code>, <code>Unzip</code>",
    "Ordering &amp; slicing — <code>SortBy</code>, <code>OrderBy</code>, <code>Reverse</code>, <code>Take</code>, <code>TakeRight</code>, <code>Drop</code>, <code>DropRight</code>",
    "Randomness (seeded) — <code>Sample</code>, <code>SampleN</code>, <code>Shuffle</code> over an injected <code>*math/rand.Rand</code>",
    "Math &amp; numbers — <code>Sum</code>, <code>SumBy</code>, <code>Mean</code>, <code>Min</code>, <code>Max</code>, <code>Clamp</code>, <code>Range</code>, <code>RangeStep</code>, <code>InRange</code>",
    "Objects &amp; maps — <code>Keys</code>, <code>Values</code>, <code>Entries</code>, <code>Pick</code>, <code>Omit</code>, <code>MapKeys</code>, <code>MapValues</code>, <code>Invert</code>, <code>Merge</code>, <code>Assign</code>",
    "Nested access — <code>Get</code> and <code>Has</code> resolve dot-separated paths into <code>map[string]any</code> trees",
    "String casing — <code>CamelCase</code>, <code>PascalCase</code>, <code>SnakeCase</code>, <code>KebabCase</code>, <code>StartCase</code>, <code>Capitalize</code>, <code>Deburr</code>",
    "String shaping — <code>Words</code>, <code>Pad</code>, <code>PadStart</code>, <code>PadEnd</code>, <code>Truncate</code>, <code>Repeat</code>, <code>Trim</code>",
    "Function wrappers — <code>Once</code>, <code>Memoize</code>, <code>MemoizeBy</code>, <code>After</code>, <code>Before</code>",
    "Rate control — <code>NewDebouncer</code> (<code>Debouncer</code>) and <code>NewThrottler</code> (<code>Throttler</code>) with injectable timers and clocks",
    "Type-safe via Go 1.24 generics — no reflection on hot paths, no <code>interface{}</code> juggling at call sites",
    "Zero dependencies — pure Go standard library, nothing to audit but the toolchain"
  ],
  node_code:
`const _ = require("lodash");

const nums = [1, 2, 3, 4, 5, 6];
const evens = _.filter(nums, n => n % 2 === 0);
const doubled = _.map(evens, n => n * 2);
const total = _.sum(doubled);   // 24

_.camelCase("Foo Bar-baz");     // "fooBarBaz"
_.snakeCase("fooBarBaz");       // "foo_bar_baz"`,
  go_code:
`import lodash "github.com/malcolmston/lodash"

nums := []int{1, 2, 3, 4, 5, 6}
evens := lodash.Filter(nums, func(n int) bool { return n%2 == 0 })
doubled := lodash.Map(evens, func(n int) int { return n * 2 })
total := lodash.Sum(doubled) // 24

lodash.CamelCase("Foo Bar-baz") // "fooBarBaz"
lodash.SnakeCase("fooBarBaz")   // "foo_bar_baz"`,
  integrate:
`<span class="tok-c">// Group a slice of records by a derived key, then reduce each bucket.</span>
type User struct{ Name, City string; Age int }
users := []User{{"Ada", "London", 36}, {"Linus", "Helsinki", 54}, {"Grace", "London", 45}}
byCity := lodash.GroupBy(users, func(u User) string { return u.City })
oldest, _ := lodash.MaxBy(byCity["London"], func(u User) int { return u.Age })

<span class="tok-c">// Sort and de-duplicate without mutating the input, then take the top two.</span>
ages := lodash.Map(users, func(u User) int { return u.Age })
ranked := lodash.Reverse(lodash.SortBy(lodash.Uniq(ages), func(a int) int { return a }))
top := lodash.Take(ranked, 2)

<span class="tok-c">// Reshape a map: keep only some keys, then invert it.</span>
scores := map[string]int{"ada": 36, "linus": 54, "grace": 45}
picked := lodash.Pick(scores, "ada", "grace")
byScore := lodash.Invert(picked)

<span class="tok-c">// Reach into a nested map[string]any config with a dotted path.</span>
cfg := map[string]any{"db": map[string]any{"host": "localhost", "port": 5432}}
host, ok := lodash.Get(cfg, "db.host")`
};
