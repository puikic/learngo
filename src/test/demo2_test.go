package demo_test

import "testing"

var pre *TreeNode = nil
var left bool = false
var right bool = false

//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

func isValidBST(root *TreeNode) bool {
	// 递归法, 中序
	if root == nil {
		return true
	}
	left = isValidBST(root.Left)
	if pre != nil && pre.Val >= root.Val {
		return false
	}
	pre = root
	right = isValidBST(root.Right)
	return left && right
}

func TestIsValidBST(t *testing.T) {
	root := &TreeNode{}
	root.Left = &TreeNode{Val: -1}
	root.Val = 0
	t.Log(isValidBST(root))
}
