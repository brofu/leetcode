package tree

import (
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func GenerateBinaryTreeFromSlice(numbers []int, index int) *TreeNode {

	if index > len(numbers)-1 {
		return nil
	}

	number := numbers[index]
	if number == math.MaxInt {
		return nil
	}
	left := GenerateBinaryTreeFromSlice(numbers, index*2+1)
	right := GenerateBinaryTreeFromSlice(numbers, index*2+2)
	return &TreeNode{Val: number, Left: left, Right: right}
}

type Queue struct {
	elements []*TreeNode
}

func (q *Queue) Push(node *TreeNode) {
	q.elements = append(q.elements, node)
}

func (q *Queue) Pop() *TreeNode {
	n := q.elements[0]
	q.elements = q.elements[1:]
	return n
}

func (q *Queue) IsEmpty() bool {
	return len(q.elements) == 0
}

func BreadthFirstTraverse(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	res := []int{}
	q := &Queue{elements: []*TreeNode{}}
	q.Push(root)

	for !q.IsEmpty() {
		node := q.Pop()
		res = append(res, node.Val)
		if node.Left != nil {
			q.Push(node.Left)
		}
		if node.Right != nil {
			q.Push(node.Right)
		}
	}

	return res
}

func BreadthFirstTraverseRecursive(root *TreeNode) []int {

	if root == nil {
		return []int{}
	}

	res := []int{root.Val}

	nodes := []*TreeNode{root.Left, root.Right}

	BreadthFirstTraverseRecursiveWrapper(&res, nodes)
	return res
}

func BreadthFirstTraverseRecursiveWrapper(res *[]int, nodes []*TreeNode) {
	nextLayerNodes := []*TreeNode{}
	for _, node := range nodes {
		if node != nil {
			*res = append(*res, node.Val)
			nextLayerNodes = append(nextLayerNodes, node.Left)
			nextLayerNodes = append(nextLayerNodes, node.Right)
		}
	}
	if len(nextLayerNodes) != 0 {
		BreadthFirstTraverseRecursiveWrapper(res, nextLayerNodes)
	}
}
