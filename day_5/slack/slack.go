package slack

import (
	"fmt"
)

func Test_1() {
	a := []int(nil)
	var b []int
	c := make([]int, 0)

	fmt.Printf("%p \n", a)
	fmt.Printf("%v \n", a == nil)

	fmt.Printf("%p \n", b)
	fmt.Printf("%v \n", b == nil)

	fmt.Printf("%p \n", c)
	fmt.Printf("%v \n", c == nil)
}
