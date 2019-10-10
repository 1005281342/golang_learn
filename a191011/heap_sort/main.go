package main

import "fmt"

func main() {
	fmt.Println(HeapSortMax([]int{3, 1, 56, 22}, 4))
	fmt.Println(HeapSort([]int{3, 1, 56, 22, 55, 12, 46523}))
}

func HeapSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		j := len(arr) - i
		HeapSortMax(arr, j)
		fmt.Println(arr)
		arr[0], arr[j-1] = arr[j-1], arr[0]
		fmt.Println("swap: ", arr)
	}
	return arr
}

func HeapSortMax(arr []int, length int) []int {

	if length <= 1 {
		return arr
	} else {
		d := length/2 - 1
		for i := d; i >= 0; i-- {
			left := 2*i + 1
			right := 2*i + 2
			if left < length && arr[left] > arr[i] {
				arr[i], arr[left] = arr[left], arr[i]
			}
			if right < length && arr[right] > arr[i] {
				arr[i], arr[right] = arr[right], arr[i]
			}
		}
		return arr
	}
}
