package maps

import (
	"fmt"
	"sort"
)

func SortKeys(m map[int]string) []int {
	keys := []int{}
	for k := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	return keys
}

func PrintSorted(m map[int]string) {
	keys := SortKeys(m)
	for _, v := range keys {
		fmt.Printf("%v ", m[v])
	}
}
