package tree

/**
KP
	1.	Use array and coding rule to present a tree
*/

func pathSumIV(nums []int) int {
	sum := 0
	if len(nums) == 0 {
		return sum
	}
	nodeMap := make(map[int]int)

	for _, num := range nums {
		val := num % 10
		code := num / 10
		nodeMap[code] = val
	}

	var traverse func(int, int, int)

	rootCode := nums[0] / 10
	traverse = func(rootCode int, path, depth int) {

		rootVal, exists := nodeMap[rootCode]
		if !exists {
			return
		}

		rootLocation := rootCode % 10

		leftCode := (depth+1)*10 + rootLocation*2 - 1
		rightCode := (depth+1)*10 + rootLocation*2

		_, existLeft := nodeMap[leftCode]
		_, existRight := nodeMap[rightCode]

		if !existLeft && !existRight {
			sum += (path + rootVal)
			return
		}

		traverse(leftCode, path+rootVal, depth+1)
		traverse(rightCode, path+rootVal, depth+1)
	}

	traverse(rootCode, 0, 1)
	return sum
}
