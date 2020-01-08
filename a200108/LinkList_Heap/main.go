package main

import "fmt"

//遍历大树
func PrintHQ(H *TreeNode) {
	if H == nil {
		return
	}
	PrintHQ(H.left)
	PrintHQ(H.right)
	fmt.Println(H.value, "  ")
}

func main() {
	H := NewPQ(3)
	H = H.Insert(2)
	H = H.Insert(1)
	H = H.Insert(0)
	H = H.Insert(4)
	H = H.Insert(14)
	H = H.Insert(44)
	H = H.Insert(10)

	PrintHQ(H)
	H, data := H.HeapPop()
	fmt.Println("min", data)
	H, data = H.HeapPop()
	fmt.Println("min", data)
}

type TreeNode struct {
	value int       // 数据值
	left  *TreeNode // 左节点
	right *TreeNode // 右节点
	level int       // 层级
}

// 开辟一个节点
func NewPQ(v int) *TreeNode {
	hq := new(TreeNode)
	hq.value = v
	return hq
}

func (HQ *TreeNode) Insert(data int) *TreeNode {
	insertNode := new(TreeNode)
	insertNode.value = data

	t := Merge(insertNode, HQ, false)
	return t
}

func (HQ *TreeNode) HeapPop() (*TreeNode, int) {
	if HQ == nil {
		return nil, -1 // 表示没有数据可以弹出
	}
	ans := HQ.value
	t := Merge(HQ.left, HQ.right, false)
	return t, ans
}

func Merge(left, right *TreeNode, reversed bool) *TreeNode {
	if left == nil {
		return right
	}
	if right == nil {
		return left
	}

	a, b := left.value, right.value
	if reversed {
		a, b = b, a
	}

	if a < b {
		return MergeSort(left, right)
	} else {
		return MergeSort(right, left)
	}
}

func MergeSort(left, right *TreeNode) *TreeNode {
	if left.left == nil {
		left.left = right
	} else {
		left.right = Merge(left.right, right, false)
		if left.left.level < left.right.level {
			left.left.level, left.right.level = left.right.level, left.left.level
		}
		left.level = left.right.level + 1
	}
	return left
}
