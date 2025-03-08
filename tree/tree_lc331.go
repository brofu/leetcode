package tree

import (
	"strings"
)

/*
Sub Tasks. DFS

Key Points
 1. Actually, the tree is constructed.
 2. Too many edged cases need to consider
    a. "1" : No enough nil node case 1
    b. "1,#": No enough nil node case 2
    c. "1,#,#,#,2": to many nil node

3. No a good enough idea
*/
func isValidSerializationSubTask(preorder string) bool {

	list := strings.Split(preorder, ",")
	var traverse func([]string, int) int

	traverse = func(list []string, index int) int {

		if index == -1 {
			return -1
		}

		if index >= len(list) { // no enough nil node
			return -1
		}

		if list[index] == "#" {
			return index
		}

		idx := traverse(list, index+1)
		if idx == -1 {
			return idx
		}
		idx = traverse(list, idx+1)

		return idx
	}

	result := traverse(list, 0)

	if result == -1 {
		return false
	}

	return result == len(list)-1 // too many nil node
}

/*
Key Point

 1. 就是利用二叉树节点和边的关系。 每个非空的二叉树节点都会产生两条边，并消耗一条边；而每个空节点只会消耗一条边	
 2. Reduce the space complexity
*/
func isValidSerialization(preorder string) bool {

	slot := 1

	for idx := 0; idx < len(preorder); {
		if preorder[idx] == ',' {
			idx++
			continue
		}
		slot--
		if slot < 0 {
			return false
		}
		if preorder[idx] != '#' {
			slot += 2
		}
		for idx < len(preorder) && preorder[idx] != ',' { // why need this? case: "9,#,92,#,#"
			idx++
		}
	}
	return slot == 0
}
