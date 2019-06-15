package main

import (
	"fmt"
	"golang_learn/day_5/slack"
	"golang_learn/day_5/sort_case"
	"golang_learn/day_5/tree_case"
)

func main() {

	slack.Test_1()

	var ch chan int
	ch = make(chan int) // 如果不初始化会怎么样 https://github.com/developer-learning/reading-go/issues/392
	go func() {
		ch <- 1
	}()
	fmt.Println(<-ch)

	b := [...]int{4, 8, 5, 2, 1, 9}
	//sort_case.Bsort(b[:])
	//sort_case.Ssort(b[:])
	sort_case.Isort(b[:])
	fmt.Println(b)

	l2 := tree_case.TreeCase{
		Val: "l2",
	}

	r2 := tree_case.TreeCase{
		Val: "r2",
	}

	l1 := tree_case.TreeCase{
		Val:   "l1",
		Left:  &l2,
		Right: &r2,
	}

	s0 := tree_case.TreeCase{
		Val:  "0",
		Left: &l1,
	}

	tree_case.TreePrint(&s0)
}
