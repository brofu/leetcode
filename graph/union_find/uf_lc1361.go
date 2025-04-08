package uf

/*
Key Points

# Validate Tree by Graph with direction

1. Root node has 0 indegree and all the others have 1 indegree
2. There is and ONLY is ONE root
3. Connected (Connected Component number is 1)
*/

func validateBinaryTreeNodes(n int, leftChild []int, rightChild []int) bool {

	indegree := make([]int, n)
	for _, num := range leftChild {
		if num != -1 {
			indegree[num]++
		}
	}

	for _, num := range rightChild {
		if num != -1 {
			indegree[num]++
		}
	}

	root := -1
	for idx, ind := range indegree {
		if ind == 0 {
			if root != -1 { // more than 1 node has indegree NOT equal 0
				return false
			}
			root = idx
		}
		if ind > 1 { // there is node has indegree larger that 1
			return false
		}
	}

	if root == -1 {
		return false // has NO root
	}

	uf := NewPathCompactedUF(n)

	for idx, num := range leftChild {
		if num != -1 {
			uf.Connect(idx, num)
		}
	}

	for idx, num := range rightChild {
		if num != -1 {
			uf.Connect(idx, num)
		}
	}

	return uf.Count() == 1
}
