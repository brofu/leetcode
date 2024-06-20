package tree

import (
	"math"
)

func pathInZigZagTree(label int) []int {
	result := []int{}

	for label >= 1 {
		result = append(result, label)
		label /= 2
		depth := int(math.Log2(float64(label))) //in math, the label should > 0 as input of log functions. But looks like not like this.
		minInDepth := int(math.Pow(float64(2), float64(depth)))
		maxInDepth := minInDepth*2 - 1
		//fmt.Println("flag", depth, minInDepth, maxInDepth)

		label = maxInDepth - (label - minInDepth)
	}

	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
	return result
}
