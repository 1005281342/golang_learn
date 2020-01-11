package main

import (
	"fmt"
	"golang_learn/helper/math"
)

func main() {
	helper := NewTreeNode(0)
	root := NewTreeNode(10)

	root = helper.Insert(root, 1)
	fmt.Printf("%+v", root)
}

// TreeNode Node
type TreeNode struct {
	value int
	level int
	left  *TreeNode
	right *TreeNode
}

func NewTreeNode(value int) *TreeNode {
	node := new(TreeNode)
	node.value = value
	//node.level = 0
	return node
}

// 获取树的高度
func (a *TreeNode) TreeHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return math.Max(a.TreeHeight(root.left), a.TreeHeight(root.right))
}

// 计算平衡因子。在进行以下操作时需要更新受影响的所有节点的高度
// 1。在插入节点时, 沿插入的路径更新节点的高度值
// 2。在删除节点时, 沿删除的路径更新节点的高度值
func (a *TreeNode) GetBalanceFactor(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.left != nil && root.right != nil {
		return root.left.level - root.right.level
	}

	if root.left == nil && root.right != nil {
		return -root.right.value
	}

	return 0
}

// 树的平衡化操作 左旋与右旋

// 右旋
// 根节点C与其左孩子B相交换，若B存在右孩子节点时，则需要将其右孩子D与旋转后的节点C相连，成为C节点的左孩子
func (a *TreeNode) RotateRight(root *TreeNode) *TreeNode {

	left := root.left // 获取C节点的左孩子节点B

	root.left = left.right // 将将要被抛弃的节点连接为旋转后的root的左孩子节点
	left.right = root      // 调换父子关系

	left.level = math.Max(a.TreeHeight(left.left), a.TreeHeight(left.right)) + 1
	return left
}

// 左旋
func (a *TreeNode) RotateLeft(root *TreeNode) *TreeNode {

	right := root.right

	root.right = right.left
	right.left = root // 调换父子关系

	right.level = math.Max(a.TreeHeight(right.left), a.TreeHeight(right.right)) + 1
	return right
}

/*
需要平衡的4种情况
LL
RR
LR
RL
*/
// 平衡化操作
func (a *TreeNode) ReBalance(root *TreeNode) *TreeNode {
	factor := a.GetBalanceFactor(root)
	//if math.ABS(factor) <= 1 {
	//	return root
	//}
	if factor > 1 && a.GetBalanceFactor(root.left) > 0 { // LL
		return a.RotateRight(root)
	} else if factor > 1 && a.GetBalanceFactor(root.left) <= 0 { // LR
		root.left = a.RotateLeft(root.left)
		return a.RotateRight(root)
	} else if factor < -1 && a.GetBalanceFactor(root.right) <= 0 { // RR
		return a.RotateLeft(root)
	} else if factor < -1 && a.GetBalanceFactor(root.right) > 0 { // RL
		root.right = a.RotateRight(root.right)
		return a.RotateLeft(root)
	} else {
		return root
	}
}

// 插入
func (a *TreeNode) Insert(root *TreeNode, value int) *TreeNode {
	if root == nil {
		root = NewTreeNode(value)
	} else if root.value == value {
		return root // 该值已存在 无法继续插入
	} else if value > root.value { // 在右边插入
		if root.right != nil {
			a.Insert(root.right, value)
		} else {
			root.right = NewTreeNode(value)
		}
	} else if value < root.value { // 在左边插入
		if root.left != nil {
			a.Insert(root.left, value)
		} else {
			root.left = NewTreeNode(value)
		}
	}
	return a.ReBalance(root)
}

// 删除
func (a *TreeNode) Delete(root *TreeNode, value int) *TreeNode {
	if root == nil {
		return root
	}
	if root.value == value { // 要移除的是根节点时
		if root.right != nil { // 存在右子树时，选择其右子树最小节点作为根节点，进行调整
			root.value = root.right.value // 将右子树的值给根节点再从右子树中移除节点
			root.right = a.Delete(root.right, root.value)
		} else {
			root = root.left // 只存在左子树时
		}
	} else if value > root.value {
		root = a.Delete(root.right, value)
	} else if value < root.value {
		root = a.Delete(root.left, value)
	}
	return a.ReBalance(root)
}
