package main

import "math"

func main() {

}

const MAX = 1000000007

func findBestValue(arr []int, target int) int {
	length := len(arr)
	if length == 1 {
		return target
	}
	tv := target / length
	var tmp = target * length
	var ans int
	mm := target / 2
	for i := 0; i < mm; i++ {
		cnt := 0
		tt := tv + i
		if tt > 100000 {
			continue
		}
		for _, v := range arr {
			if v > tt {
				cnt += tt
			} else {
				cnt += v
			}
		}
		tm := int(math.Abs(float64(cnt) - float64(target)))
		if tm < tmp {
			tmp = tm
			ans = tt
		}
	}

	for i := -1; i > -mm; i-- {
		cnt := 0
		tt := tv + i
		if tt < 1 {
			continue
		}
		for _, v := range arr {
			if v > tt {
				cnt += tt
			} else {
				cnt += v
			}
		}
		tm := int(math.Abs(float64(cnt) - float64(target)))
		if tm < tmp {
			tmp = tm
			ans = tt
		}
	}

	return ans
}

func replaceElements(arr []int) []int {
	ans := make([]int, len(arr))
	for i := 0; i < len(arr)-1; i++ {
		var maxNum int
		for j := i + 1; j < len(arr); j++ {
			if arr[j] > maxNum {
				maxNum = arr[j]
			}
		}
		ans[i] = maxNum

	}
	ans[len(arr)-1] = -1
	return ans
}
