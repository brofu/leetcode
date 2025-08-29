package bfs

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*

KP

* TC:
1. Constructor. O(N) N is the node number in the tree
2. Insert. O(1)
3. Get_root. O(1)

* SC:
1. O(N/2), since the node with left AND right child would be removed from the queue.
2. Over all, it's O(N)
*/

type CBTInserter struct {
	root *TreeNode
	q    []*TreeNode
}

func Constructor(root *TreeNode) CBTInserter {
	this := CBTInserter{
		root: root,
	}

	q := []*TreeNode{root}

	for len(q) > 0 {
		current := q[0]
		if current.Left == nil {
			break
		}
		q = append(q, current.Left)

		if current.Right == nil {
			break
		}
		q = append(q, current.Right)
		q = q[1:]
	}

	this.q = q
	return this
}

func (this *CBTInserter) Insert(val int) int {

	newNode := &TreeNode{
		Val: val,
	}

	q := this.q
	parent := q[0]

	q = append(q, newNode)

	if q[0].Left == nil {
		q[0].Left = newNode
	} else {
		q[0].Right = newNode
		q = q[1:]
	}

	this.q = q

	return parent.Val
}

func (this *CBTInserter) Get_root() *TreeNode {
	return this.root
}
