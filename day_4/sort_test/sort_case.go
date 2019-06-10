package sort_test

import (
	"fmt"
	"sort"
)

func SortCase() {
	var a = [...]int{1, 2, 58, 88, 4}
	sort.Ints(a[:])
	fmt.Println(a)
}

func SearchCase() {
	var a = [...]int{1, 2, 58, 88, 4}
	sort.Ints(a[:])
	index := sort.SearchInts(a[:], 88)
	fmt.Println(index)
}
