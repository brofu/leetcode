package tree

import (
	"math"
)

type NodeWithParent struct {
	Val    int
	Left   *NodeWithParent
	Right  *NodeWithParent
	Parent *NodeWithParent
}

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

type TreeNodeWithSubtreeSum struct {
	Val   int
	Sum   int
	Left  *TreeNodeWithSubtreeSum
	Right *TreeNodeWithSubtreeSum
}

func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// GenerateTreeNodeFromSliceBFS construct the tree from slice with BFS
// leetcode style
func GenerateTreeNodeFromSliceBFS(numbers []int) *TreeNode {
	if len(numbers) == 0 {
		return nil
	}
	root := &TreeNode{numbers[0], nil, nil}
	q := []*TreeNode{root}
	numbers = numbers[1:]

	for len(q) > 0 {
		size := len(q)
		for idx := 0; idx < size; idx++ {
			node := q[idx]
			if 2*idx < len(numbers) {
				num := numbers[2*idx]
				if num != math.MaxInt {
					leftNode := &TreeNode{num, nil, nil}
					node.Left = leftNode
					q = append(q, leftNode)
				}
			}
			if 2*idx+1 < len(numbers) {
				num := numbers[2*idx+1]
				if num != math.MaxInt {
					rightNode := &TreeNode{num, nil, nil}
					node.Right = rightNode
					q = append(q, rightNode)
				}
			}
		}
		q = q[size:]

		if 2*size < len(numbers) {
			numbers = numbers[2*size:]
		} else {
			break
		}
	}

	return root
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

func GenerateTreeNodeWithParentFromSlice(numbers []int, index int, parent *NodeWithParent) *NodeWithParent {

	if index > len(numbers)-1 {
		return nil
	}

	number := numbers[index]
	if number == math.MaxInt {
		return nil
	}
	node := &NodeWithParent{
		Val:    number,
		Parent: parent,
	}

	node.Left = GenerateTreeNodeWithParentFromSlice(numbers, index*2+1, node)
	node.Right = GenerateTreeNodeWithParentFromSlice(numbers, index*2+2, node)

	return node
}

func GetNodeWithParentFromTree(root *NodeWithParent, key int) *NodeWithParent {
	if root == nil {
		return nil
	}

	if root.Val == key {
		return root
	}

	if res := GetNodeWithParentFromTree(root.Left, key); res != nil {
		return res
	}

	return GetNodeWithParentFromTree(root.Right, key)
}
func GetTreeNodeFromTree(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val == key {
		return root
	}

	if res := GetTreeNodeFromTree(root.Left, key); res != nil {
		return res
	}

	return GetTreeNodeFromTree(root.Right, key)
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

func GenerateSliceFromTreeNodeBFS(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	q := []*TreeNode{root}
	result := []int{}

	for len(q) > 0 {

		size := len(q)

		for idx := 0; idx < size; idx++ {
			node := q[idx]
			result = append(result, node.Val)
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		q = q[size:]
	}

	return result
}

func GenerateSliceFromTreeNodeBFSWithPadding(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	q := []*TreeNode{root}
	result := []int{}

	for len(q) > 0 {

		size := len(q)

		for idx := 0; idx < size; idx++ {
			node := q[idx]
			if node == nil {
				result = append(result, math.MaxInt)
				continue
			}

			result = append(result, node.Val)

			if node.Left == nil && node.Right == nil {
				continue
			}
			q = append(q, node.Left)
			q = append(q, node.Right)
		}
		q = q[size:]
	}
	if len(result) > 0 && result[len(result)-1] == math.MaxInt {
		result = result[:len(result)-1]
	}
	return result
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

func GetMinFromBST(root *TreeNode) int {
	for root.Left != nil {
		root = root.Left
	}
	return root.Val
}
