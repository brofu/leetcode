package dc

import "strconv"

/*
KP:
	1. memo to avoid duplicate calculation.
	2. How to process base case. No need to check a number
	3. 对于运算表达式相关的问题，一般都会涉及括号以及优先级的问题，常用的技巧是分治算法，明确递归函数的定义，让递归函数去处理括号。	


TC:
	1. Recursive depth: O(K-1). K is the number of operator
	2. At each recursive layer, need to process the left and right result combination, which is j*k (as well as append), and the worst case should be (N/2)^2 ==> oh, that's wrong
	3. So, we have O(K-1)*(j*k), which is O(K-1)*(N/2)^2
	4. But in real case, it would be much less than that:
		* (N/2)^2 only happens when the depth is lgN
		* there is memo
	5. And, there should K splits. So, overall is O(K*(K-1)*(N/2)^2)

*/
func diffWaysToCompute(expression string) []int {

	memo := make(map[string][]int)
	return caculateRecursive(expression, memo)
}

func caculateRecursive(expression string, memo map[string][]int) []int {

	if results, exists := memo[expression]; exists {
		return results
	}

	results := make([]int, 0)

	for idx, operx := range expression {
		if operx != '+' && operx != '-' && operx != '*' {
			continue
		}

		left := caculateRecursive(expression[0:idx], memo)
		right := caculateRecursive(expression[idx+1:], memo)

		for _, l := range left {
			for _, r := range right {
				if operx == '+' {
					results = append(results, l+r)
				}
				if operx == '-' {
					results = append(results, l-r)
				}
				if operx == '*' {
					results = append(results, l*r)
				}
			}
		}
	}

	//base case
	if len(results) == 0 {
		result, _ := strconv.Atoi(expression)
		results = []int{result}
	}

	// record memo
	memo[expression] = results

	return results
}
