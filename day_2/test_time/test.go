package test_time

import (
	"fmt"
	"time"
)

const (
	_ = iota
	Man
	Female
)

func OddEven() {
	if time.Now().Unix()%Female == 0 {
		fmt.Println(Female)
	} else {
		fmt.Println(Man)
	}
}
