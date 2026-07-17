package lodash

import (
	"math/rand"
	"reflect"
	"testing"
)

func eqInts(t *testing.T, got, want []int) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func TestMap(t *testing.T) {
	got := Map([]int{1, 2, 3}, func(x int) int { return x * 2 })
	eqInts(t, got, []int{2, 4, 6})
	strs := Map([]int{1, 2}, func(x int) string {
		if x == 1 {
			return "a"
		}
		return "b"
	})
	if !reflect.DeepEqual(strs, []string{"a", "b"}) {
		t.Fatalf("map to string failed: %v", strs)
	}
}

func TestMapI(t *testing.T) {
	got := MapI([]int{5, 6, 7}, func(v, i int) int { return v + i })
	eqInts(t, got, []int{5, 7, 9})
}

func TestFilterReject(t *testing.T) {
	even := func(x int) bool { return x%2 == 0 }
	eqInts(t, Filter([]int{1, 2, 3, 4}, even), []int{2, 4})
	eqInts(t, Reject([]int{1, 2, 3, 4}, even), []int{1, 3})
}

func TestReduce(t *testing.T) {
	sum := Reduce([]int{1, 2, 3, 4}, func(a, c int) int { return a + c }, 0)
	if sum != 10 {
		t.Fatalf("reduce sum = %d", sum)
	}
	concat := ReduceRight([]string{"a", "b", "c"}, func(a, c string) string { return a + c }, "")
	if concat != "cba" {
		t.Fatalf("reduceRight = %q", concat)
	}
}

func TestForEach(t *testing.T) {
	sum := 0
	ForEach([]int{1, 2, 3}, func(x int) { sum += x })
	if sum != 6 {
		t.Fatalf("foreach sum = %d", sum)
	}
	idxSum := 0
	ForEachI([]int{9, 9, 9}, func(_, i int) { idxSum += i })
	if idxSum != 3 {
		t.Fatalf("foreachi idx = %d", idxSum)
	}
}

func TestFind(t *testing.T) {
	v, ok := Find([]int{1, 2, 3}, func(x int) bool { return x > 1 })
	if !ok || v != 2 {
		t.Fatalf("find = %d,%v", v, ok)
	}
	_, ok = Find([]int{1, 2}, func(x int) bool { return x > 5 })
	if ok {
		t.Fatal("find should miss")
	}
	lv, ok := FindLast([]int{1, 2, 3, 2}, func(x int) bool { return x == 2 })
	if !ok || lv != 2 {
		t.Fatalf("findlast = %d", lv)
	}
	if FindIndex([]int{1, 2, 3}, func(x int) bool { return x == 3 }) != 2 {
		t.Fatal("findindex")
	}
	if FindIndex([]int{1}, func(x int) bool { return x == 9 }) != -1 {
		t.Fatal("findindex miss")
	}
	if FindLastIndex([]int{2, 2, 2}, func(x int) bool { return x == 2 }) != 2 {
		t.Fatal("findlastindex")
	}
	if FindLastIndex([]int{1}, func(x int) bool { return x == 9 }) != -1 {
		t.Fatal("findlastindex miss")
	}
}

func TestEverySome(t *testing.T) {
	if !Every([]int{2, 4}, func(x int) bool { return x%2 == 0 }) {
		t.Fatal("every")
	}
	if Every([]int{2, 3}, func(x int) bool { return x%2 == 0 }) {
		t.Fatal("every should fail")
	}
	if !Every([]int{}, func(int) bool { return false }) {
		t.Fatal("every empty must be true")
	}
	if !Some([]int{1, 2}, func(x int) bool { return x == 2 }) {
		t.Fatal("some")
	}
	if Some([]int{}, func(int) bool { return true }) {
		t.Fatal("some empty must be false")
	}
}

func TestIncludesIndex(t *testing.T) {
	if !Includes([]string{"a", "b"}, "b") {
		t.Fatal("includes")
	}
	if Includes([]string{"a"}, "z") {
		t.Fatal("includes miss")
	}
	if IndexOf([]int{5, 6, 7}, 6) != 1 {
		t.Fatal("indexof")
	}
	if IndexOf([]int{5}, 9) != -1 {
		t.Fatal("indexof miss")
	}
	if LastIndexOf([]int{1, 2, 1}, 1) != 2 {
		t.Fatal("lastindexof")
	}
	if LastIndexOf([]int{1}, 9) != -1 {
		t.Fatal("lastindexof miss")
	}
}

func TestGroupKeyCount(t *testing.T) {
	g := GroupBy([]int{1, 2, 3, 4}, func(x int) string {
		if x%2 == 0 {
			return "even"
		}
		return "odd"
	})
	if !reflect.DeepEqual(g["even"], []int{2, 4}) || !reflect.DeepEqual(g["odd"], []int{1, 3}) {
		t.Fatalf("groupby = %v", g)
	}
	k := KeyBy([]string{"aa", "bb", "cc"}, func(s string) byte { return s[0] })
	if k['a'] != "aa" || k['c'] != "cc" {
		t.Fatalf("keyby = %v", k)
	}
	c := CountBy([]int{1, 1, 2, 3, 3, 3}, func(x int) int { return x })
	if c[1] != 2 || c[3] != 3 {
		t.Fatalf("countby = %v", c)
	}
}

func TestPartition(t *testing.T) {
	tr, fa := Partition([]int{1, 2, 3, 4}, func(x int) bool { return x > 2 })
	eqInts(t, tr, []int{3, 4})
	eqInts(t, fa, []int{1, 2})
}

func TestUniq(t *testing.T) {
	eqInts(t, Uniq([]int{1, 1, 2, 3, 3, 3}), []int{1, 2, 3})
	got := UniqBy([]int{1, 2, 3, 4}, func(x int) int { return x % 2 })
	eqInts(t, got, []int{1, 2})
}

func TestChunk(t *testing.T) {
	got := Chunk([]int{1, 2, 3, 4, 5}, 2)
	want := [][]int{{1, 2}, {3, 4}, {5}}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("chunk = %v", got)
	}
	if !reflect.DeepEqual(Chunk([]int{}, 3), [][]int{}) {
		t.Fatal("chunk empty")
	}
}

func TestChunkPanics(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Fatal("expected panic")
		}
	}()
	Chunk([]int{1}, 0)
}

func TestFlatten(t *testing.T) {
	eqInts(t, Flatten([][]int{{1, 2}, {3}, {}, {4, 5}}), []int{1, 2, 3, 4, 5})
	nested := []any{1, []any{2, []any{3, 4}}, 5}
	eqInts(t, FlattenDeep[int](nested), []int{1, 2, 3, 4, 5})
}

func TestCompactReverse(t *testing.T) {
	eqInts(t, Compact([]int{0, 1, 0, 2, 3, 0}), []int{1, 2, 3})
	if !reflect.DeepEqual(Compact([]string{"", "a", "", "b"}), []string{"a", "b"}) {
		t.Fatal("compact strings")
	}
	eqInts(t, Reverse([]int{1, 2, 3}), []int{3, 2, 1})
}

func TestZipUnzip(t *testing.T) {
	pairs := Zip([]int{1, 2, 3}, []string{"a", "b"})
	if len(pairs) != 2 || pairs[0].First != 1 || pairs[1].Second != "b" {
		t.Fatalf("zip = %v", pairs)
	}
	a, b := Unzip(pairs)
	eqInts(t, a, []int{1, 2})
	if !reflect.DeepEqual(b, []string{"a", "b"}) {
		t.Fatalf("unzip b = %v", b)
	}
}

func TestSetOps(t *testing.T) {
	eqInts(t, Difference([]int{1, 2, 3, 4}, []int{2, 4}), []int{1, 3})
	eqInts(t, Intersection([]int{1, 2, 3}, []int{2, 3, 4}, []int{0, 2, 3}), []int{2, 3})
	eqInts(t, Union([]int{1, 2}, []int{2, 3}, []int{3, 4}), []int{1, 2, 3, 4})
	eqInts(t, Without([]int{1, 2, 3, 2, 1}, 2), []int{1, 3, 1})
	if Intersection[int]() != nil {
		t.Fatal("intersection empty")
	}
}

func TestTakeDrop(t *testing.T) {
	eqInts(t, Take([]int{1, 2, 3, 4}, 2), []int{1, 2})
	eqInts(t, Take([]int{1}, 5), []int{1})
	eqInts(t, Take([]int{1}, -1), []int{})
	eqInts(t, TakeRight([]int{1, 2, 3, 4}, 2), []int{3, 4})
	eqInts(t, Drop([]int{1, 2, 3, 4}, 2), []int{3, 4})
	eqInts(t, DropRight([]int{1, 2, 3, 4}, 2), []int{1, 2})
	eqInts(t, Drop([]int{1, 2}, 5), []int{})
	eqInts(t, DropRight([]int{1, 2}, -3), []int{1, 2})
	eqInts(t, TakeRight([]int{1, 2}, -1), []int{})
}

func TestSampleShuffle(t *testing.T) {
	rng := rand.New(rand.NewSource(42))
	v, ok := Sample([]int{1, 2, 3, 4, 5}, rng)
	if !ok {
		t.Fatal("sample ok")
	}
	if !Includes([]int{1, 2, 3, 4, 5}, v) {
		t.Fatalf("sample out of range: %d", v)
	}
	_, ok = Sample([]int{}, rng)
	if ok {
		t.Fatal("sample empty")
	}

	// Deterministic given a fixed seed.
	r1 := rand.New(rand.NewSource(7))
	r2 := rand.New(rand.NewSource(7))
	s := []int{1, 2, 3, 4, 5, 6}
	if !reflect.DeepEqual(Shuffle(s, r1), Shuffle(s, r2)) {
		t.Fatal("shuffle not deterministic for same seed")
	}
	// Original not mutated.
	eqInts(t, s, []int{1, 2, 3, 4, 5, 6})

	sn := SampleN([]int{1, 2, 3, 4, 5}, 3, rand.New(rand.NewSource(1)))
	if len(sn) != 3 || len(Uniq(sn)) != 3 {
		t.Fatalf("sampleN not unique: %v", sn)
	}
	all := SampleN([]int{1, 2}, 9, rand.New(rand.NewSource(1)))
	if len(all) != 2 {
		t.Fatalf("sampleN overflow: %v", all)
	}
	if len(SampleN([]int{1, 2}, -1, rand.New(rand.NewSource(1)))) != 0 {
		t.Fatal("sampleN negative")
	}
}

type person struct {
	Name string
	Age  int
}

func TestSortByOrderBy(t *testing.T) {
	people := []person{{"Bob", 30}, {"Alice", 25}, {"Carol", 30}}
	byAge := SortBy(people, func(p person) int { return p.Age })
	if byAge[0].Name != "Alice" {
		t.Fatalf("sortby = %v", byAge)
	}
	// Original untouched.
	if people[0].Name != "Bob" {
		t.Fatal("sortby mutated input")
	}

	ordered := OrderBy(people,
		[]func(a, b person) int{
			func(a, b person) int { return a.Age - b.Age },
			func(a, b person) int {
				switch {
				case a.Name < b.Name:
					return -1
				case a.Name > b.Name:
					return 1
				default:
					return 0
				}
			},
		},
		[]Order{Desc, Asc},
	)
	// Age desc, then name asc: Bob(30), Carol(30), Alice(25)
	if ordered[0].Name != "Bob" || ordered[1].Name != "Carol" || ordered[2].Name != "Alice" {
		t.Fatalf("orderby = %v", ordered)
	}
}

func TestConcatFill(t *testing.T) {
	eqInts(t, Concat([]int{1}, []int{2, 3}, []int{4}), []int{1, 2, 3, 4})
	eqInts(t, Fill(7, 3), []int{7, 7, 7})
	eqInts(t, Fill(1, -2), []int{})
}
