package tree

func diameterOfBinaryTree(root *TreeNode) int {
	_, diameter := diameterOfBinaryTreeWraper(root)
	return diameter
}

func diameterOfBinaryTreeWraper(root *TreeNode) (int, int) {
	if root == nil {
		return 0, 0
	}

	leftDepth, leftDiameter := diameterOfBinaryTreeWraper(root.Left)
	rightDepth, rightDiameter := diameterOfBinaryTreeWraper(root.Right)

	currentDiameter := leftDepth + rightDepth
	currentDiameter = MaxInt(MaxInt(leftDiameter, rightDiameter), currentDiameter)

	return (MaxInt(leftDepth, rightDepth)) + 1, currentDiameter
}

func diameterOfBinaryTreeV2(root *TreeNode) int {
	maxDiameter := 0
	traverseDiameterOfBinaryTree(root, &maxDiameter)
	return maxDiameter
}

func traverseDiameterOfBinaryTree(root *TreeNode, maxDiameter *int) int {
	if root == nil {
		return 0
	}
	left := traverseDiameterOfBinaryTree(root.Left, maxDiameter)
	right := traverseDiameterOfBinaryTree(root.Right, maxDiameter)
	myDiameter := left + right
	if myDiameter > *maxDiameter {
		*maxDiameter = myDiameter
	}
	return MaxInt(left, right) + 1
}
