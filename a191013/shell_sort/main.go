package main

import "fmt"

func main() {

	fmt.Println(ShellSort([]int{1, 5, 63, 2, 6, 723, 8436, 22, 6, 3, 2, 9}))
}

// 一般用在并发场合

// 步长收缩
func ShellSort(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr
	}
	gap := length / 2 // 步长
	for gap > 0 {
		for i := 0; i < gap; i++ {
			ShellSortStep(arr, i, gap)
		}
		gap /= 2
	}
	return arr
}

// 插入排序的变种
func ShellSortStep(nums []int, start, gap int) {
	length := len(nums)
	for i := start; i < length; i += gap {
		j := i
		for j > gap && nums[j] < nums[j-gap] {
			nums[j], nums[j-gap] = nums[j-gap], nums[j]
			j -= gap
		}
		fmt.Println(gap, nums)
	}
}
