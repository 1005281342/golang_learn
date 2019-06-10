package defer_test

import "fmt"

func DeferCase(a *int) {
	defer fmt.Println(*a)
	*a = 100
}
