package main

import (
	array "GO/go-dsa/array"
	"fmt"
)

func main() {
	r := array.UnmarkedSumArray([]int{1, 12, 12, 4, 14, 1, 12, 1}, [][]int{
		{1, 2},
		{5, 4},
		{4, 0},
		{0, 1},
		{0, 3},
	})
	fmt.Printf("%v\n", r)
}
