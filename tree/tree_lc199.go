package tree

/*
Traverse + BFS
*/

func rightSideView(root *TreeNode) []int {

	res := []int{}
	if root == nil {
		return res
	}

	q := []*TreeNode{root}

	for len(q) > 0 {
		size := len(q)
		for idx := 0; idx < size; idx++ {
			node := q[idx]
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		res = append(res, q[size-1].Val)
		q = q[size:]
	}

	return res
}

func rightSideViewDFS(root *TreeNode) []int {

	temp := make(map[int]int)
	deepth := 0
	var traverse func(*TreeNode)

	traverse = func(root *TreeNode) {
		if root == nil {
			return
		}
		deepth++
		temp[deepth] = root.Val

		traverse(root.Left)
		traverse(root.Right)
		deepth--
	}

	traverse(root)
	res := make([]int, len(temp))
	for idx := 0; idx < len(res); idx++ {
		res[idx] = temp[idx+1]
	}
	return res
}
