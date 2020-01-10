package main

import (
	"golang_learn/helper/math"
)

func main() {

}

// AVL Node
type TreeNode struct {
	value int
	level int
	left  *TreeNode
	right *TreeNode
}

type AVL TreeNode

// 获取树的高度
func (a *AVL) TreeHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return math.Max(a.TreeHeight(root.left), a.TreeHeight(root.right))
}

// 计算平衡因子。在进行以下操作时需要更新受影响的所有节点的高度
// 1。在插入节点时, 沿插入的路径更新节点的高度值
// 2。在删除节点时, 沿删除的路径更新节点的高度值
// ** 当平衡因子[绝对值]大于1时就会触发树的修正，或者说是树的再平衡操作
func (a *AVL) GetABSBalanceFactor(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return math.ABS(root.left.level - root.right.level)
}

// 树的平衡化操作 左旋与右旋

// 右旋
// 根节点C与其左孩子B相交换，若B存在右孩子节点时，则需要将其右孩子D与旋转后的节点C相连，成为C节点的左孩子
func (a *AVL) RotateRight(root *TreeNode) *TreeNode {

	left := root.left // 获取C节点的左孩子节点B

	root.left = left.right // 将将要被抛弃的节点连接为旋转后的root的左孩子节点
	left.right = root      // 调换父子关系

	left.level = math.Max(a.TreeHeight(left.left), a.TreeHeight(left.right)) + 1
	return left
}

// 左旋
func (a *AVL) RotateLeft(root *TreeNode) *TreeNode {

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
func (a *AVL) ReBalance() {

}
