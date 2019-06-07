package main

import "fmt"

func main() {
	sum, avg := add_avg(10, 20)
	fmt.Println(sum, avg)
}

func add_avg(a, b int) (int, int) {
	sum := a + b
	avg := sum >> 1
	return sum, avg
}
