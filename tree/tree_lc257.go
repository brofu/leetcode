package tree

import (
	"fmt"
	"strings"
)

/*
Traverse + DFS
*/
func binaryTreePaths(root *TreeNode) []string {

	var res []string
	var path []string

	var traverse func(*TreeNode, *[]string, *[]string)

	traverse = func(root *TreeNode, res, path *[]string) {

		// base case 1
		if root == nil {
			return
		}

		// base case 2
		if root.Left == nil && root.Right == nil {
			*res = append(*res, strings.Join(append(*path, fmt.Sprintf("%d", root.Val)), "->"))
			return
		}

		// choose
		*path = append(*path, fmt.Sprintf("%d", root.Val))

		// next layer
		traverse(root.Left, res, path)
		traverse(root.Right, res, path)

		// trackback
		*path = (*path)[:len(*path)-1]
	}
	traverse(root, &res, &path)
	return res
}
