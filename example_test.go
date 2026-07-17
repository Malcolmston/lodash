package lodash_test

import (
	"fmt"

	lodash "github.com/malcolmston/lodash"
)

// Example demonstrates chaining a few of the collection helpers together.
func Example() {
	nums := []int{1, 2, 3, 4, 5, 6}
	evens := lodash.Filter(nums, func(n int) bool { return n%2 == 0 })
	doubled := lodash.Map(evens, func(n int) int { return n * 2 })
	total := lodash.Sum(doubled)
	fmt.Println(evens)
	fmt.Println(doubled)
	fmt.Println(total)
	// Output:
	// [2 4 6]
	// [4 8 12]
	// 24
}

// ExampleGroupBy shows grouping elements by a computed key.
func ExampleGroupBy() {
	words := []string{"apple", "banana", "avocado", "cherry", "blueberry"}
	groups := lodash.GroupBy(words, func(s string) byte { return s[0] })
	fmt.Println(groups['a'])
	fmt.Println(groups['b'])
	// Output:
	// [apple avocado]
	// [banana blueberry]
}

// ExampleCamelCase demonstrates string case conversion.
func ExampleCamelCase() {
	fmt.Println(lodash.CamelCase("Foo Bar-baz"))
	fmt.Println(lodash.SnakeCase("fooBarBaz"))
	// Output:
	// fooBarBaz
	// foo_bar_baz
}
