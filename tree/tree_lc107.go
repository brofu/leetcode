package tree

/*
BFS

Key Points:

1. Insert into the head of list
*/
func levelOrderBottom(root *TreeNode) [][]int {
	result := [][]int{}
	if root == nil {
		return result
	}

	q := []*TreeNode{root}

	for len(q) > 0 {

		size := len(q)
		temp := make([]int, size)
		for idx := 0; idx < size; idx++ {
			node := q[idx]
			temp[idx] = node.Val
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		result = append([][]int{temp}, result...)
		q = q[size:]
	}

	return result
}

/*
BFS

Key Points:

1. Reverse the result finally
*/

func levelOrderBottomV2(root *TreeNode) [][]int {
	result := [][]int{}
	if root == nil {
		return result
	}

	q := []*TreeNode{root}

	for len(q) > 0 {

		size := len(q)
		temp := make([]int, size)
		for idx := 0; idx < size; idx++ {
			node := q[idx]
			temp[idx] = node.Val
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		result = append(result, temp)
		q = q[size:]
	}

	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return result
}
