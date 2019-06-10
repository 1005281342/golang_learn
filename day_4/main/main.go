package main

import (
	"fmt"
	"golang_learn/day_4/fib"
	"golang_learn/day_4/sort_test"
)

func main() {

	// 初始化数组
	var a0 [5]int = [5]int{1, 2, 3} // int类型默认值为0
	var a1 = [5]int{1, 2, 3}
	var a2 = [...]int{1, 2, 3, 4, 50}
	var s0 = [5]string{3: "hello", 4: "tom"} // string类型默认值是空字符串
	fmt.Println(a0)
	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(s0)
	fib.Fib(10)

	sort_test.SortCase()
	sort_test.SearchCase()
}
