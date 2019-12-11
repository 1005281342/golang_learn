package main

func nextGreaterElements(nums []int) []int {
	length := len(nums)
	tmp := append(nums, nums...)
	res := make([]int, length, length)
	for i, v := range nums {
		for j := i + 1; j <= i+length; j++ {
			if v < tmp[j] {
				res[i] = tmp[j]
			}
		}
	}
	return res
}

func main() {

}
