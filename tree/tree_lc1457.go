package tree

/*
Traverse + DFS

Key Point

	-How to check if a path is a `pseudo-palindromic`?
	-If there are 0 or at most 1 number occurs ODD times in the path
		- How to check
		- Map, record the occur times of each time. Space Complexity O(N)
		- Bit operation. Space Complexity O(1)
			- set bit of a number. mask = 1 << node.Val
			- if the number happen with EVEN times, mask ^ mask == 0, else 1
			- if there is at most 1 number occurs ODD times, then, there would at least ONE 1, in `mask`
			- check mask & (mask -1)
*/

func pseudoPalindromicPaths(root *TreeNode) int {

	var traverse func(*TreeNode, int)

	sum := 0
	traverse = func(root *TreeNode, mask int) {
		if root == nil {
			return
		}

		mask ^= 1 << root.Val
		if root.Left == nil && root.Right == nil {
			if mask&(mask-1) == 0 {
				sum++
			}
		}

		traverse(root.Left, mask)
		traverse(root.Right, mask)
	}

	traverse(root, 0)

	return sum
}
