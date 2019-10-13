package main

import "fmt"

func main() {
	fmt.Println(MergeSort([]int{1, 5, 63, 2, 6, 723, 8436, 22, 6, 3, 2, 9}))
}

func Merge(a, b []int) []int {
	la := len(a)
	if la == 0 {
		return b
	}
	lb := len(b)
	if lb == 0 {
		return a
	}
	c := make([]int, 0, 0)
	ia, ib := 0, 0
	for ia < la && ib < lb {
		if a[ia] < b[ib] {
			c = append(c, a[ia])
			ia++
		} else if a[ia] > b[ib] {
			c = append(c, b[ib])
			ib++
		} else {
			c = append(c, a[ia])
			c = append(c, b[ib])
			ia++
			ib++
		}
	}
	if ia < la {
		c = append(c, a[ia:]...)
	}
	if ib < lb {
		c = append(c, b[ib:]...)
	}
	return c
}

func MergeSort(arr []int) []int {

	length := len(arr)
	if length <= 1 {
		return arr
	}
	mid := length / 2
	return Merge(MergeSort(arr[:mid]), MergeSort(arr[mid:]))
}
