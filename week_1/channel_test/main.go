package main

import (
	"fmt"
	"golang_learn/week_1/package_1"
)

func add(a, b int, c chan int) {

	c <- a + b
}

func main() {

	pipe := make(chan int, 1)
	go add(10, 20, pipe)
	sum := <-pipe
	fmt.Println(sum)
	go package_1.Goroute(100, 100, pipe)
	sum = <-pipe
	fmt.Println(sum)
}

// 这种方法不好
//// 声明
//var pipe chan int
//
//func add(a, b int)  {
//
//	pipe <- a+b
//}
//
//
//func main() {
//
//	// 创建
//	pipe = make(chan int, 1)
//	go add(10, 20)
//
//	sum :=<- pipe
//	fmt.Println(sum)
//}
