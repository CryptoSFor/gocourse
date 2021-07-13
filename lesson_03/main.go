package main

import (
	"fmt"
	"storages/arrays"
	"storages/maps"
	"storages/slices"
)

func main() {
	array := [5]float64{1, 2, 3, 4, 5}
	slice1 := []string{"one", "two", "a"}
	slice2 := []int64{1, 2, 3, 4, 5}
	mapa := map[int]string{10: "aa", 0: "bb", 500: "cc"}

	fmt.Println(arrays.AverageValue(array))
	fmt.Println(slices.MaxLength(slice1))
	fmt.Println(slices.Reverse(slice2))
	maps.PrintSorted(mapa)

}
