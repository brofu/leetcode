package tree

import (
	"fmt"
	"strconv"
	"strings"
)

type Codec struct {
	holder    string
	delimiter string
}

func Constructor() Codec {
	return Codec{
		holder:    "#",
		delimiter: "*",
	}
}

/*
DFS + Sub Task

1. Construct the left child
2. Construct the right child
3. Construct current node

Key Point:

1. The delimiter can NOT be "-", since there would be negative number
2. How to restore the tree, with a pre-order array with pointer "nil".
3. How to find the bounder of left child and right child
4. O(n)
*/
func (this *Codec) serialize(root *TreeNode) string {

	nodes := []*TreeNode{}

	var traverse func(*TreeNode)
	traverse = func(root *TreeNode) {
		if root == nil {
			nodes = append(nodes, nil)
			return
		}
		nodes = append(nodes, root)
		traverse(root.Left)
		traverse(root.Right)
	}
	traverse(root)

	values := make([]string, len(nodes))
	for idx, node := range nodes {
		if node == nil {
			values[idx] = this.holder
		} else {
			values[idx] = fmt.Sprintf("%d", node.Val)
		}
	}

	//fmt.Println(strings.Join(values, this.delimiter))
	return strings.Join(values, this.delimiter)
}

// Deserializes your encoded data to tree.

func (this *Codec) deserialize(data string) *TreeNode {
	values := strings.Split(data, this.delimiter)

	// the root node of the sub tree, and the index of the last right node of the sub tree
	var traverse func([]string, int) (*TreeNode, int)

	traverse = func(values []string, start int) (*TreeNode, int) {
		if values[start] == this.holder {
			return nil, start
		}

		left, idx := traverse(values, start+1)
		right, idx := traverse(values, idx+1)

		val, _ := strconv.ParseInt(values[start], 10, 32)

		root := &TreeNode{
			Val:   int(val),
			Left:  left,
			Right: right,
		}

		return root, idx
	}

	root, _ := traverse(values, 0)

	return root
}

/*
BFS

Key Point:

1. The delimiter can NOT be "-", since there would be negative number
2. Use SAME approach to serialize and deserialize the nodes
3. O(n)
*/

func (this *Codec) serializeBFS(root *TreeNode) string {

	q := []*TreeNode{root}
	nodes := []*TreeNode{}

	for len(q) > 0 {
		size := len(q)
		for i := 0; i < size; i++ {
			cur := q[i]
			nodes = append(nodes, cur)
			if cur != nil {
				q = append(q, cur.Left)
				q = append(q, cur.Right)
			}
		}
		q = q[size:]
	}

	values := make([]string, len(nodes))
	for idx, node := range nodes {
		if node == nil {
			values[idx] = this.holder
		} else {
			values[idx] = fmt.Sprintf("%d", node.Val)
		}
	}

	//fmt.Println(strings.Join(values, this.delimiter))
	return strings.Join(values, this.delimiter)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserializeBFS(data string) *TreeNode {

	values := strings.Split(data, this.delimiter)

	if values[0] == this.holder {
		return nil
	}

	rootVal, _ := strconv.ParseInt(values[0], 10, 32)
	root := &TreeNode{
		Val: int(rootVal),
	}
	q := []*TreeNode{root}

	index := 1
	for len(q) > 0 {
		size := len(q)
		for idx := 0; idx < size; idx++ {
			parent := q[idx]

			leftVale := values[index]
			index += 1
			if leftVale != this.holder {
				val, _ := strconv.ParseInt(leftVale, 10, 32)
				leftNode := &TreeNode{
					Val: int(val),
				}
				parent.Left = leftNode
				q = append(q, leftNode)
			}

			rightValue := values[index]
			index += 1
			if rightValue != this.holder {
				val, _ := strconv.ParseInt(rightValue, 10, 32)
				rightNode := &TreeNode{
					Val: int(val),
				}
				parent.Right = rightNode
				q = append(q, rightNode)
			}
		}
		q = q[size:]
	}
	return root
}
