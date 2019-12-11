package main

import (
	"fmt"
	"unsafe"
)

func main() {

	const x = unsafe.Sizeof(0)
	const y = len([8]byte{})
	fmt.Println(x, y)
}
