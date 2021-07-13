package maps

import (
	"fmt"
	"sort"
)

func SortKeys(mapa map[int]string) []int {
	keys := []int{}
	for k := range mapa {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	return keys
}

func PrintSorted(mapa map[int]string) {
	keys := SortKeys(mapa)
	for _, v := range keys {
		fmt.Printf("%v ", mapa[v])
	}
}
