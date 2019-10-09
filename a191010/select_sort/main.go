package main

import (
	"fmt"
	"log"
	"strings"
)

func main() {
	res := SelectSort([]int{1, 4, 23, 3, 6, 7})
	log.Printf("选择排序结果是 %v", res)
	arr := SelectSortStringArr([]string{"aa", "a", "bbb", "b", "bb", "c"})
	arr2 := SelectSortStringArrByCompare([]string{"aa", "a", "bbb", "b", "bb", "c"})
	log.Printf("选择排序结果是 %v", arr)
	log.Printf("选择排序2结果是 %v", arr2)
}

func GreaterThan(a, b string) bool {
	al := len(a)
	bl := len(b)
	if al > bl { // "aa" > "a"
		for i := 0; i < bl; i++ {
			if int(a[i]) >= int(b[i]) {
				return true
			}
		}
		return false
	} else {
		for i := 0; i < al; i++ {
			if int(b[i]) < int(a[i]) {
				return true
			}
		}
		return false
	}
}

func SelectSortStringArr(arr []string) []string {

	length := len(arr)
	for i := 0; i < length-1; i++ {
		fmt.Printf("setp %v, %v", i, arr)
		fmt.Println()
		for j := i + 1; j < length; j++ {
			fmt.Println(GreaterThan(arr[i], arr[j]))
			if GreaterThan(arr[i], arr[j]) {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return arr
}

func SelectSortStringArrByCompare(arr []string) []string {

	length := len(arr)
	for i := 0; i < length-1; i++ {
		fmt.Printf("setp %v, %v", i, arr)
		fmt.Println()
		for j := i + 1; j < length; j++ {
			if strings.Compare(arr[i], arr[j]) > 0 {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return arr
}

func SelectSort(nums []int) []int {

	length := len(nums)
	for i := 0; i < length-1; i++ {
		fmt.Printf("setp %v, %v", i, nums)
		fmt.Println()
		for j := i + 1; j < length; j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
	return nums
}
