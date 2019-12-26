package main

import "fmt"

func main() {
	fmt.Println(HashCode("123456"))
}

func HashCode(s string) int64 {
	var mask int64
	mask = (1 << 32) - 1
	var h int64
	for _, c := range s {
		h = (h << 5 & mask) | (h >> 27)
		h += int64(c)
	}
	return h
}
