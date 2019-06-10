package switch_test

import (
	"fmt"
	"math/rand"
)

func Case_1() {
	n := rand.Intn(100)
	for {
		var input int
		var state = false
		fmt.Scanf("%d\n", &input)
		switch {
		case input == n:
			fmt.Println("猜对了")
			state = true
		case input < n:
			fmt.Println("小了")
		default:
			fmt.Println("大了")
		}
		if state {
			break
		}
	}

}
