package tree

func preorderTraversal(root *TreeNode) []int {

	if root == nil {
		return []int{}
	}

	array := []int{root.Val}
	array = append(array, preorderTraversal(root.Left)...)
	array = append(array, preorderTraversal(root.Right)...)
	return array
}

var res []int

func preorderTraversalV2(root *TreeNode) []int {
	traverse(root)
	return res
}

func traverse(root *TreeNode) {
	if root == nil {
		return
	}
	res = append(res, root.Val)
	traverse(root.Left)
	traverse(root.Right)
}
