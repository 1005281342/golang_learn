package main

import "fmt"

func sum_list(n int) {

	for i := 0; i <= n; i++ {
		//fmt.Println(i,"+", 5-i, "= 5")
		fmt.Printf("%d+%d=%d\n", i, n-i, n)
	}
}

func main() {

	sum_list(10)
}
