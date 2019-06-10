package fib

import "fmt"

func Fib(n int) {
	if n <= 2 {
		n = 2
	}
	//var res [n]int64  // non-constant array bound n
	res := make([]int64, n) // 实例化动态数组要使用make
	res[0] = 1
	fmt.Println(res[0])
	res[1] = 1
	fmt.Println(res[0])
	for i := 2; i < n; i++ {
		res[i] = res[i-1] + res[i-2]
		fmt.Println(res[i])
	}
}
