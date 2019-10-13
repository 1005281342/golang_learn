package main

import "fmt"

func main() {
	fmt.Println(OESort([]int{1, 5, 63, 2, 6, 723, 8436, 22, 6, 3, 2, 9}))
}

// 奇偶排序
func OESort(arr []int) []int {
	length := len(arr)
	if length <= 0 {
		return arr
	}
	isSorted := false
	for isSorted == false {
		isSorted = true
		// 对奇数位进行排序
		for i := 1; i < length-1; i += 2 {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				isSorted = false
			}
		}

		// 对偶数位进行排序
		for i := 0; i < length-1; i += 2 {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				isSorted = false
			}
		}
	}
	return arr
}
