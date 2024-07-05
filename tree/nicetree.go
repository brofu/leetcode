package tree

/**
check if a NICE binary tree contains a value
NICE tree means
1. root.val = 0
2. if it has left child, the val is 2 * root.val + 1
2. if it has right child, the val is 2 * root.right + 2
*/
func Query(root *TreeNode, target int) bool {

	track := make([]int, 0)
	parent := target

	for parent != 0 {
		track = append(track, parent)
		if parent%2 == 0 {
			parent = (parent - 2) / 2
		} else {
			parent = (parent - 1) / 2
		}
	}
	var index int
	for index = len(track) - 1; index >= 0 && root != nil; index-- {
		if track[index] == 2*root.Val+1 {
			root = root.Left
		} else {
			root = root.Right
		}
	}
	return index == -1 && root != nil
}
