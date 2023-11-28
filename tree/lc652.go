package tree

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {

	existMap := make(map[[3]int][2]int)

	res := []*TreeNode{}

	treeId := 1

	var dfs func(*TreeNode) int
	dfs = func(root *TreeNode) (id int) {

		if root == nil { // id empty node is 0
			return -201
		}

		leftId := dfs(root.Left)
		rightId := dfs(root.Right)

		key := [3]int{root.Val, leftId, rightId}

		val, exist := existMap[key]
		if exist {
			id = val[0]
			if val[1] == 1 {
				res = append(res, root)
				existMap[key] = [2]int{id, 2}
			}
		} else {
			id = treeId
			existMap[key] = [2]int{id, 1}
			treeId++
		}

		return
	}

	dfs(root)

	return res
}
