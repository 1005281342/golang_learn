package main

import (
	"fmt"
	"golang_learn/day_2/add"
	a2 "golang_learn/day_2/add2"
	"golang_learn/day_2/modify"
	"golang_learn/day_2/test_time"
)

func main() {
	fmt.Println(add.Name)
	fmt.Println(add.Age)

	fmt.Println(a2.Age)
	fmt.Println(a2.Name)

	fmt.Println(a2.Age)
	fmt.Println(a2.Name)

	test_time.OddEven()

	// 值传递与引用传递
	a := 100
	fmt.Println(a)
	modify.Modify(a)
	fmt.Println(a)
	modify.Modify_1(&a)
	fmt.Println(a)

	c, d := 10, 5
	c, d = d, c
	fmt.Println(c, d)

	modify.Swap2(&c, &d)
	fmt.Println(c, d)
}
