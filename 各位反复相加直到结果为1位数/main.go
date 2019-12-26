package main

import "fmt"

func main() {
	fmt.Println(addDigits(112))
}

func addDigits(num int) int {
	if num > 9 {
		return (num-1)%9 + 1
	}
	return num
}
