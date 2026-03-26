package main

import (
	hashtable "GO/go-dsa/hashtable"
	array "GO/go-dsa/array"
)

func main() {
	k := xorAfterQueries([]int{780}, [][]int{
		{0, 0, 1, 13},
		{0, 0, 1, 17},
		{0, 0, 1, 9},
		{0, 0, 1, 18},
		{0, 0, 1, 16},
		{0, 0, 1, 6},
		{0, 0, 1, 4},
		{0, 0, 1, 11},
		{0, 0, 1, 7},
		{0, 0, 1, 18},
		{0, 0, 1, 8},
		{0, 0, 1, 15},
		{0, 0, 1, 12},
	})
	fmt.Println(k)}
