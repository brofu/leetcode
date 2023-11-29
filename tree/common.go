package tree

import (
	"math"
)

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

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
			//fmt.Println(node, node.Val)
			*res = append(*res, node.Val)
			nextLayerNodes = append(nextLayerNodes, node.Left)
			nextLayerNodes = append(nextLayerNodes, node.Right)
		}
	}
	if len(nextLayerNodes) != 0 {
		BreadthFirstTraverseRecursiveWrapper(res, nextLayerNodes)
	}
}
func GenerateSliceFromTreeNode(root *TreeNode) []int {
	return BreadthFirstTraverseRecursive(root)
}

func GenerateNodeFromSlice(numbers []int, index int) *Node {

	if index > len(numbers)-1 {
		return nil
	}

	number := numbers[index]
	if number == math.MaxInt {
		return nil
	}
	left := GenerateNodeFromSlice(numbers, index*2+1)
	right := GenerateNodeFromSlice(numbers, index*2+2)
	return &Node{Val: number, Left: left, Right: right}
}

// GenerateSlinceFromNode generate the slice from a linked perfect binary tree
func GenerateSlinceFromNode(root *Node) []int {

	res := []int{}
	if root == nil {
		return res
	}

	pointer := root

	for pointer != nil {
		res = append(res, pointer.Val)
		pointer = pointer.Next
	}

	res = append(res, math.MaxInt)

	res = append(res, GenerateSlinceFromNode(root.Left)...)

	return res
}

func Pow(a, b int) int {
	mul := 1
	for i := 0; i < b; i++ {
		mul *= a
	}
	return mul
}

func FindIndex(nums []int, num int) int {
	for index, val := range nums {
		if val == num {
			return index
		}
	}

	return -1
}

func FindMax(nums []int) (index, val int) {

	index, val = 0, nums[0]

	for i, v := range nums {
		if val < v {
			index, val = i, v
		}
	}

	return
}

func NewBinaryIndexTree(length int) BinaryIndexedTree {
	return make(BinaryIndexedTree, length+1)
}

type BinaryIndexedTree []int

func (BIT BinaryIndexedTree) Add(index, val int) {
	i := index + 1
	for i < len(BIT) {
		BIT[i] += val
		i += i & (-i)
	}
}

func (BIT BinaryIndexedTree) Sum(index int) int {
	sum := 0
	for i := index + 1; i > 0 && i < len(BIT); i -= i & (-i) {
		sum += BIT[i]
	}
	return sum
}

type BinaryIndexedTreeReverse []int

func NewBinaryIndexTreeReverse(length int) BinaryIndexedTreeReverse {
	return make(BinaryIndexedTreeReverse, length+1)
}

func (BITR BinaryIndexedTreeReverse) Add(index, val int) {
	for i := index + 1; i > 0; i -= i & (-i) {
		BITR[i] += val
	}
}

func (BITR BinaryIndexedTreeReverse) Sum(index int) int {
	sum := 0
	for i := index + 1; i < len(BITR); i += i & (-i) {
		sum += BITR[i]
	}
	return sum
}
