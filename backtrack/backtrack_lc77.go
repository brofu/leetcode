package backtrack

import "fmt"

func combine(n int, k int) [][]int {
	result := [][]int{}
	track := []int{}

	var bt func(int, int, int, []int, *[][]int)
	bt = func(n, k, start int, track []int, result *[][]int) {

		fmt.Println("flag", track)
		// base case
		if k == len(track) {
			temp := make([]int, k)
			//temp := []int{}
			copy(temp, track)
			*result = append(*result, temp)
			return
		}

		for i := start; i <= n; i++ {
			// choose
			track = append(track, i)
			// next layer
			bt(n, k, i+1, track, result)
			// cancel choose
			track = track[:len(track)-1]
		}
	}

	bt(n, k, 1, track, &result)
	return result
}

func combineV2(n int, k int) [][]int {
	result := [][]int{}
	track := []int{}

	var bt func(int, int, int, []int, *[][]int)
	bt = func(n, k, start int, track []int, result *[][]int) {

		fmt.Println("flag", track)

		// base case
		if k == len(track) {
			temp := make([]int, k)
			//temp := []int{}
			copy(temp, track)
			*result = append(*result, temp)
			return
		}

		for i := start; i <= n; i++ {
			// filter some cases
			if n-i+1 < k-len(track) { // the left number of numbers less than k - len(track)
				continue
			}
			// choose, next layer, and cancel choose
			//track = append(track, i)
			bt(n, k, i+1, append(track, i), result)
			//track = track[:len(track)-1]
		}
	}

	bt(n, k, 1, track, &result)
	return result
}

func combineV3(n int, k int) [][]int {

	result := make([][]int, 0)
	var bt func(start int, path []int)
	bt = func(start int, path []int) {

		//base case
		if len(path) == k {
			temp := make([]int, len(path))
			copy(temp, path)
			result = append(result, temp)
			return
		}

		for idx := start; idx <= n; idx++ {
			bt(idx+1, append(path, idx))
		}

	}

	bt(1, []int{})
	return result
}
