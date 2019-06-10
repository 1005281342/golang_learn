package main

import (
	"fmt"
	"golang_learn/day_5/slack"
)

func main() {

	slack.Test_1()

	var ch chan int
	ch = make(chan int) // 如果不初始化会怎么样 https://github.com/developer-learning/reading-go/issues/392
	go func() {
		ch <- 1
	}()
	fmt.Println(<-ch)
}
