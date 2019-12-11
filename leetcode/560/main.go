package main

func main() {

}

func subarraySum(nums []int, k int) int {
	length := len(nums)
	st := make([]int, length, length)
	st[0] = nums[0]
	cnt := 0
	if nums[0] == k {
		cnt += 1
	}
	for i := 1; i < length; i++ {
		st[i] = st[i-1] + nums[i]
		if nums[i] == k {
			cnt += 1
		}
	}
	for i := 0; i < length-1; i++ {
		for j := i + 1; j < length; j++ {
			if st[j]-st[i]+nums[i] == k {
				cnt += 1
			}
		}
	}
	return cnt
}
