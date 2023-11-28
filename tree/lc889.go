package tree

import "fmt"

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

func constructFromPrePostV2(preorder []int, postorder []int) *TreeNode {
	return constructFromPrePostV2Wrapper(preorder, 0, len(preorder)-1, postorder, 0)
}

func constructFromPrePostV2Wrapper(preorder []int, prestart, preend int, postorder []int, poststart int) *TreeNode {

	fmt.Println(preorder, prestart, preend, postorder, poststart)
	if prestart > preend {
		return nil
	}

	node := &TreeNode{
		Val: preorder[prestart],
	}

	if prestart == preend {
		return node
	}

	leftVal := preorder[prestart+1]
	leftIndexInPostOrder := FindIndex(postorder, leftVal)

	leftSize := leftIndexInPostOrder + 1 - poststart

	node.Left = constructFromPrePostV2Wrapper(preorder, prestart+1, prestart+leftSize, postorder, poststart)
	node.Right = constructFromPrePostV2Wrapper(preorder, prestart+leftSize+1, preend, postorder, leftIndexInPostOrder+1)

	return node
}
