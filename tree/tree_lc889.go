package tree

import "fmt"

/*
Sub Task

Key Point
  - Need to check if current node has ONLY 1 child, if yes, put it on the left or the right. Both is OK
*/

func constructFromPrePostSubTaskPV1(preorder []int, postorder []int) *TreeNode {

	var traverse func([]int, []int) *TreeNode

	traverse = func(preorder, postorder []int) *TreeNode {
		if len(preorder) == 0 || len(postorder) == 0 || len(preorder) != len(postorder) {
			return nil
		}

		cur := preorder[0]
		root := &TreeNode{
			Val: cur,
		}

		// leaf node
		if len(preorder) == 1 {
			return root
		}

		leftVal := preorder[1]
		rightVal := postorder[len(postorder)-2]

		// only 1 child
		if leftVal == rightVal {
			root.Left = traverse(preorder[1:], postorder[:len(postorder)-1])
			// or put the right location. Both are ok
			// root.Right = traverse(preorder[1:], postorder[:len(postorder)-1])
			return root
		}

		// 2 child
		indexLeft := -1
		for idx, v := range postorder {
			if v == leftVal {
				indexLeft = idx
				break
			}
		}
		indexRight := -1
		for idx, v := range preorder {
			if v == rightVal {
				indexRight = idx
				break
			}
		}

		root.Left = traverse(preorder[1:indexRight], postorder[:indexLeft+1])
		root.Right = traverse(preorder[indexRight:], postorder[indexLeft+1:len(postorder)-1])

		return root
	}

	return traverse(preorder, postorder)
}

/*
Sub Task

Key Points
 1. May Use HashMap to reduce the complexity
 2. Edged case: only 1 child
*/
func constructFromPrePostSubTaskPV2(preorder []int, postorder []int) *TreeNode {

	var traverse func([]int, []int) *TreeNode

	traverse = func(preorder, postorder []int) *TreeNode {

		if len(preorder) == 0 {
			return nil
		}

		root := &TreeNode{
			Val: preorder[0],
		}
		if len(preorder) == 1 {
			return root
		}

		// only 1 child
		if postorder[len(postorder)-2] == preorder[1] {
			root.Left = traverse(preorder[1:], postorder[:len(postorder)-1])
			return root
		}

		rightValIndex := -1
		for idx, num := range preorder {
			if num == postorder[len(postorder)-2] {
				rightValIndex = idx
				break
			}
		}
		leftValIndex := -1
		for idx, num := range postorder {
			if num == preorder[1] {
				leftValIndex = idx
				break
			}
		}
		root.Left = traverse(preorder[1:rightValIndex], postorder[0:leftValIndex+1])
		root.Right = traverse(preorder[rightValIndex:], postorder[leftValIndex+1:len(postorder)-1])
		return root
	}

	return traverse(preorder, postorder)
}

/*
Sub Task + Using the array index (instead of slice)

Key Points
 1. May Use HashMap to reduce the complexity
 2. Edged case: only 1 child
*/

func constructFromPrePostSubTaskPV3(preorder []int, postorder []int) *TreeNode {
	preMap := make(map[int]int)
	for idx, num := range preorder {
		preMap[num] = idx
	}
	postMap := make(map[int]int)
	for idx, num := range postorder {
		postMap[num] = idx
	}

	return traversePrePostSubTaskPV3(preorder, postorder, 0, len(preorder)-1, 0, len(postorder)-1, preMap, postMap)
}

func traversePrePostSubTaskPV3(preorder, postorder []int, preStart, preEnd, postStart, postEnd int, preMap, postMap map[int]int) *TreeNode {

	fmt.Println(preStart, preEnd, postStart, postEnd)
	if preStart > preEnd {
		return nil
	}

	root := &TreeNode{
		Val: preorder[preStart],
	}

	if preStart == preEnd {
		return root
	}

	leftVal := preorder[preStart+1]
	rightVal := postorder[postEnd-1]

	// only 1 child
	if leftVal == rightVal {
		root.Left = traversePrePostSubTaskPV3(preorder, postorder, preStart+1, preEnd, postStart, postEnd-1, preMap, postMap)
		return root
	}

	root.Left = traversePrePostSubTaskPV3(preorder, postorder, preStart+1, preMap[rightVal]-1, postStart, postMap[leftVal], preMap, postMap)
	root.Right = traversePrePostSubTaskPV3(preorder, postorder, preMap[rightVal], preEnd, postMap[leftVal]+1, postEnd-1, preMap, postMap)
	return root
}
