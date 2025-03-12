package tree

/*
Sub Tasks

Key Points:

 1. Why sub task? Need to check children tree first.
 2. Record the forest during the traverse
 3. The `root` node need to check out side the recursion.
*/
func delNodes(root *TreeNode, to_delete []int) []*TreeNode {

	valueMap := make(map[int]struct{})
	for _, val := range to_delete {
		valueMap[val] = struct{}{}
	}

	result := []*TreeNode{}
	var traverse func(*TreeNode) *TreeNode

	traverse = func(root *TreeNode) *TreeNode {

		if root == nil {
			return root
		}

		if root.Left == nil && root.Right == nil { // leaf node
			if _, exists := valueMap[root.Val]; exists { // need to delete
				return nil
			}
			return root
		}

		root.Left = traverse(root.Left)
		root.Right = traverse(root.Right)

		if _, exists := valueMap[root.Val]; exists { // need to delete
			if root.Left != nil { // collect the forest
				result = append(result, root.Left)
			}
			if root.Right != nil { // collect the forest
				result = append(result, root.Right)
			}
			return nil
		}
		return root

	}

	root = traverse(root)
	if root != nil { // check the `root` node outside the recursion
		result = append(result, root)
	}

	return result
}

/*
Sub Tasks
Key Points:
 1. Key idea: if a node need to delete, and has NO parent, then, it's a root of a new tree.
*/

func delNodesV2(root *TreeNode, to_delete []int) []*TreeNode {

	valueMap := make(map[int]struct{})
	for _, val := range to_delete {
		valueMap[val] = struct{}{}
	}

	result := []*TreeNode{}
	var traverse func(*TreeNode, bool) *TreeNode

	traverse = func(root *TreeNode, hasParent bool) *TreeNode {

		if root == nil {
			return root
		}

		needDelete := false
		if _, exists := valueMap[root.Val]; exists {
			needDelete = true
		}
		if !hasParent && !needDelete {
			result = append(result, root)
		}

		root.Left = traverse(root.Left, !needDelete)
		root.Right = traverse(root.Right, !needDelete)

		if needDelete {
			return nil
		}
		return root
	}

	traverse(root, false)
	return result
}
