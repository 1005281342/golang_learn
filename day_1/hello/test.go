package hello

import "fmt"

func goroute_test(a, b int) int {
	return a + b
}

func test_print(a int) {
	fmt.Println(a)
}

func test_channel_int() {
	pipe := make(chan int, 3)
	pipe <- 1
	pipe <- 2
}

func test_channel_string() {
	pipe := make(chan string, 3)
	pipe <- "a"
	pipe <- "b"
}
