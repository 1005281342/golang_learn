package main

import "fmt"

const (
	_ = 1 << (10 * iota)
	KB
	MB
	GB
)

type color byte //自定义类型
const (
	red color = iota //指定常量类型
	yellow
	blue
)

// test 函数参数是byte类型
func test(c byte) {
	println(c)
}

// printColor 函数参数是color类型
func printColor(c color) {
	println(c)
}

func main() {
	fmt.Println(KB, MB, GB)
	printColor(yellow)
	printColor(45) // 45并未超出color类型取值范围
	//var c byte = 100
	//test(c)
	////错误：cannot use c (type byte) as type color in argument to printColor
	//printColor(c)
	const cc color = 100
	//const cc color = 256 // constant 256 overflows color
	//错误：cannot use cc (type byte) as type color in argument to printColor
	printColor(cc)
}
