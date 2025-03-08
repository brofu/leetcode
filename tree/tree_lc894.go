package tree

/*
Sub Tasks

Key Points

 1. If n is an EVEN number, then, no possible to generate a FBT
 2. How to abstract to problem, to get the Sub Tasks (left and right child)? ==> split N into 2 parts
    a. If n = 3, then, 1 left and 1 right
    b. If n = 5, then, 1 left, 3 right, and 3 left, 1 right
    c. And so on...
    d. And for each <left, right> pair, there maybe different combines...
*/

func allPossibleFBT(n int) []*TreeNode {

	if n%2 == 0 {
		return []*TreeNode{}
	}

	var buildFBT func(int) []*TreeNode

	buildFBT = func(n int) []*TreeNode {

		if n == 1 {
			return []*TreeNode{{Val: 0}}
		}

		trees := []*TreeNode{}

		for i := 1; i < n; i += 2 {
			j := n - i - 1

			leftTrees := buildFBT(i)
			rightTrees := buildFBT(j)

			for _, leftTree := range leftTrees {
				for _, rightTree := range rightTrees {
					trees = append(trees, &TreeNode{0, leftTree, rightTree})
				}
			}
		}

		return trees
	}
	return buildFBT(n)
}
