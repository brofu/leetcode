package tree

func constructFromPrePost(preorder []int, postorder []int) *TreeNode {

	if len(preorder) == 0 {
		return nil
	}

	dummy := &TreeNode{
		Val: -1,
	}

	p := dummy
	index := 0

	for ; index < len(preorder); index++ {
		if preorder[index] != postorder[len(postorder)-(index+1)] {
			break
		}
		p.Left = &TreeNode{
			Val: preorder[index],
		}
		p = p.Left
	}

	preorder = preorder[index:]

	if len(preorder) == 0 {
		return dummy.Left
	}

	postorder = postorder[:len(postorder)-index]

	leftVal := preorder[0]
	rightVal := postorder[len(postorder)-1]

	leftIndexInPostOrder := FindIndex(postorder, leftVal)
	rightIndexInPreOrder := FindIndex(preorder, rightVal)

	p.Left = constructFromPrePost(preorder[:rightIndexInPreOrder], postorder[:leftIndexInPostOrder+1])
	p.Right = constructFromPrePost(preorder[rightIndexInPreOrder:], postorder[leftIndexInPostOrder+1:len(postorder)])

	return dummy.Left
}
