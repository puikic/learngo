package demon_test

import (
	"fmt"
	"runtime"
	"testing"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func ChangeArr(arr []int) {
	arr[0] = 10010
}

func pathSum(root *TreeNode, targetSum int) [][]int {
	if root == nil {
		return [][]int{}
	}
	res := [][]int{}
	var travel func(node *TreeNode, sum int, path []int)
	travel = func(node *TreeNode, sum int, path []int) {
		path = append(path, node.Val)
		sum += node.Val
		if node.Left == nil && node.Right == nil { // 此时node为叶子节点
			if sum == targetSum {
				// temp := []int{}
				// copy(temp, path)
				res = append(res, path)
			}
			return
		}
		if node.Left != nil {
			travel(node.Left, sum, path)
		}
		if node.Right != nil {
			travel(node.Right, sum, path)
		}
	}
	travel(root, 0, []int{})
	return res
}

func TestDemo1(t *testing.T) {
	root := &TreeNode{Val: 5}
	root.Left = &TreeNode{Val: 4}
	root.Right = &TreeNode{Val: 8}
	root.Right.Left = &TreeNode{Val: 13}
	root.Right.Right = &TreeNode{Val: 4}
	root.Left.Left = &TreeNode{Val: 11}
	root.Left.Left.Left = &TreeNode{Val: 7}
	root.Left.Left.Right = &TreeNode{Val: 2}
	root.Right.Right.Left = &TreeNode{Val: 5}
	root.Right.Right.Right = &TreeNode{Val: 1}
	t.Logf("%v", pathSum(root, 22))
}

func TestF1(t *testing.T) {
	a := []int{1, 2, 3, 4}
	ChangeArr(a)
	fmt.Println(a[0])
}

func TestDemo2(t *testing.T) {
	t.Log(runtime.NumCPU())

	t.Log(runtime.NumGoroutine())
}
