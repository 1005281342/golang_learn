package main

import "fmt"

/*
将数组对半划分子数组，进行外部排序。
当子数组不可划分时。会逐层弹出栈。
*/

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
	// 进行外排
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
	// 将余下的数据追加过去
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
