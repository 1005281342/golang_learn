package hello

import "fmt"

func add(a, b int) int {
	res := a + b
	return res
}

func main() {
	fmt.Println("你好")
	res := add(10, 20)
	go goroute_test(10, 20)
	fmt.Println("add(10,20) =", res)
	for i := 0; i < 100; i++ {
		go test_print(i)
	}
}
