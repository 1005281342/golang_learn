package main

import "fmt"

func main() {
	res := InsertSort([]int{1, 4, 23, 3, 6, 7})
	fmt.Println(res)
}

func InsertSort(nums []int) []int {

	ans := []int{nums[0]}
	for i := 1; i < len(nums); i++ {
		flag := true
		for j := 0; j < len(ans); j++ {
			if nums[i] < ans[j] {
				tmp := []int{nums[i]}
				for _, v := range ans[j:] {
					tmp = append(tmp, v)
				}
				ans = append(ans[:j], tmp...)
				flag = false
				break
			}
		}
		if flag {
			ans = append(ans, nums[i])
		}
	}
	return ans
}
