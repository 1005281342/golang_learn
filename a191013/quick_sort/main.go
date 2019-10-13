package main

import (
	"fmt"
	"math/rand"
)

func main() {
	res := QuickSort([]int{2, 5, 16, 2, 6, 7, 2, 1, 2, 4, 5, 3})
	fmt.Println(res)
}

func QuickSort(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr
	} else {
		// low, mid, high
		tmp := arr[rand.Intn(length)]
		//tmp := arr[time.Now().Unix() % int64(length)]
		low := make([]int, 0, 0)
		mid := make([]int, 0, 0)
		high := make([]int, 0, 0)
		//mid = append(mid, tmp)

		for i := 0; i < length; i++ {
			if arr[i] < tmp {
				low = append(low, arr[i])
			} else if arr[i] > tmp {
				high = append(high, arr[i])
			} else {
				mid = append(mid, arr[i])
			}
		}
		low, high = QuickSort(low), QuickSort(high)
		myArr := append(append(low, mid...), high...)
		return myArr
	}

}
