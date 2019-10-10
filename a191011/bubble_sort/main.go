package main

import "fmt"

func main() {
	res := BubbleSort([]int{1, 4, 6, 2, 7, 3, 8, 22, 5})
	fmt.Println(res)
}

func BubbleSort(nums []int) []int {
	for i := 0; i < len(nums)-1; i++ {
		state := true
		for j := 0; j < len(nums)-1-i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				state = false
			}
		}
		if state {
			return nums
		}
		fmt.Println(nums)
	}
	return nums
}
