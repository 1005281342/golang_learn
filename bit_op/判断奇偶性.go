package main

import "fmt"

func main() {
	res := isO(10)
	fmt.Println(res)
}

func isO(num int) bool {
	return num&1 == 0
}
