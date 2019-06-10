package main

import (
	"fmt"
	"golang_learn/day_3/big_num_add"
	"golang_learn/day_3/defer_test"
	"golang_learn/day_3/strings_operation"
)

func main() {
	s := "www.baidu.com"
	h := "http://"
	e := "/"
	s = strings_operation.HasPrefixTest(s, h)
	fmt.Println(s)
	s = strings_operation.HasSuffixTest(s, e)
	fmt.Println(s)

	fmt.Println(strings_operation.IndexTest(s, e))
	fmt.Println(strings_operation.LastIndexTest(s, e))

	strings_operation.AtoiTest("12a3")

	//switch_test.Case_1()

	bb := 10
	defer_test.DeferCase(&bb)

	res := big_num_add.Multi("11112222333", "3265744821")
	fmt.Println(res)
	res = big_num_add.BigNumberAdd("11112222333", "3265744821")
	fmt.Println(res)
}

// 在golang中switch不需要break，满足条件执行完毕会自动跳出，需要强制执行下面的分支则需要使用 fallthrough
// goto 通过设置label的方式 实现 for 循环

// defer
