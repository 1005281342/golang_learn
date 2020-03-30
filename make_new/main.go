package main

import "fmt"

var ch = make(chan int, 10)

func main() {
	ch <- 200
	x, ok := <-ch
	close(ch)
	fmt.Println(x, ok)
	x, ok = <-ch
	fmt.Println(x, ok)

	a := new([]int)
	fmt.Println(a)
	fmt.Printf("%p \n", a)
	*a = append(*a, 1)
	fmt.Println(a)
	fmt.Printf("%p \n", a)

	bb := new(map[string]int)
	fmt.Println(*bb)
	fmt.Printf("%p \n", bb)
	//bb["a"] = 123
	//fmt.Println(bb)
	//d := *bb
	//d["22"] = 1
	//fmt.Println(d)
}
