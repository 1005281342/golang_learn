package tree_case

import "fmt"

type TreeCase struct {
	Val   string
	Left  *TreeCase
	Right *TreeCase
}

func TreePrint(x *TreeCase) {

	if x == nil {
		return
	}

	fmt.Println(x.Val)
	TreePrint(x.Left)
	TreePrint(x.Right)
}
