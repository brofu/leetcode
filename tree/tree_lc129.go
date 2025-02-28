package tree

/*
Traverse + DFS
*/

func sumNumbers(root *TreeNode) int {

	var resultList []int
	var valueList []int

	var traverse func(*TreeNode, *[]int, *[]int)

	traverse = func(root *TreeNode, resultList, valueList *[]int) {

		// base case 1
		if root == nil {
			return
		}

		// base case 2
		if root.Left == nil && root.Right == nil {
			sum := 0
			for _, val := range append(*valueList, root.Val) {
				sum = sum*10 + val
			}
			*resultList = append(*resultList, sum)
			return
		}

		// pre order location
		*valueList = append(*valueList, root.Val)

		traverse(root.Left, resultList, valueList)
		traverse(root.Right, resultList, valueList)

		// post order location
		*valueList = (*valueList)[:len(*valueList)-1]

		return
	}

	traverse(root, &resultList, &valueList)

	sum := 0

	for _, result := range resultList {
		sum += result
	}

	return sum
}

/*
Traverse + DFS

The brief version.
If just pass append(valueList, root.Val) to next layer,
Then, no need to call valueList = valueList[:len(valueList)-1] at the `post-order` location
*/

func sumNumbersV2(root *TreeNode) int {
	var resultList []int

	var traverse func(*TreeNode, *[]int, []int)

	traverse = func(root *TreeNode, resultList *[]int, valueList []int) {

		// base case 1
		if root == nil {
			return
		}

		// base case 2
		if root.Left == nil && root.Right == nil {
			sum := 0
			for _, val := range append(valueList, root.Val) {
				sum = sum*10 + val
			}
			*resultList = append(*resultList, sum)
			return
		}

		// pre order location
		// actually, it's `valueList = append(valueList, root.Val)`

		traverse(root.Left, resultList, append(valueList, root.Val))
		traverse(root.Right, resultList, append(valueList, root.Val))

		// post order location
		// No need. if we don't update valueList

		return
	}

	traverse(root, &resultList, []int{})

	sum := 0

	for _, result := range resultList {
		sum += result
	}

	return sum
}
