// File example_test.go
package sort_test

import (
	"fmt"
	"sort"
)

func ExampleInts() {
	s := []int{5, 2, 6, 3, 1, 4} // unsorted
	sort.Ints(s)
	fmt.Println(s)
	// Output: [1 2 3 4 5 6]
}
