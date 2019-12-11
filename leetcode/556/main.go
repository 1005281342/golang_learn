package main

import (
	"fmt"
	"strconv"
)

const MAX = 4294967295

func nextGreaterElement(n int) int {
	if n >= MAX || n < 0 {
		return -1
	}
	sn := string(n)
	if len(sn) == 1 {
		return -1
	}
	fmt.Println("start")
	length := len(sn)
	tmp := make([]int, length, length)
	for i, v := range []rune(sn) {
		tmp[i] = int(v)
		//if i == 0 {
		//	continue
		//}
		//if v > sn[i-1] {
		//	return -1
		//}
	}
	fmt.Println(tmp)
	for i, v := range tmp {
		if i == 0 {
			continue
		}
		if v > tmp[i-1] {
			return -1
		}
	}

	tmp[length-1], tmp[length-2] = tmp[length-2], tmp[length-1]
	var a string
	for _, v := range tmp {
		a += string(v)
	}
	d, _ := strconv.Atoi(a)
	return d
}

func main() {
	nextGreaterElement(102)
}
