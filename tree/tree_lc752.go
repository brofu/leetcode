package tree

import (
	"strconv"
)

/*
KP
	1. BFS
	2. ~100ms
*/
func openLock(deadends []string, target string) int {

	init := "0000"

	deadendsMap := make(map[string]struct{})
	for _, str := range deadends {
		deadendsMap[str] = struct{}{}
		if str == init {
			return -1
		}
	}

	occured := map[string]struct{}{}
	occured[init] = struct{}{}

	queue := []string{init}
	result := 0

	for len(queue) > 0 {
		sz := len(queue)

		for i := 0; i < sz; i += 1 {
			str := queue[i]
			if str == target { // get the target
				return result
			}

			neighbours := getNeighbour(str)

			for _, neighbour := range neighbours {
				if _, ok := deadendsMap[neighbour]; ok {
					continue
				}
				if _, ok := occured[neighbour]; ok {
					continue
				}

				queue = append(queue, neighbour)
				occured[neighbour] = struct{}{}
			}
		}
		queue = queue[sz:]
		result += 1
	}

	return -1
}

func getNeighbour(str string) [8]string {

	var result [8]string

	var helper func(input byte) (byte, byte)

	helper = func(input byte) (byte, byte) {
		if input == '9' {
			return '8', '0'
		}
		if input == '0' {
			return '9', '1'
		}
		return input - 1, input + 1
	}

	for index, s := range str {
		pre, next := helper(byte(s))
		result[2*index] = string(str[:index]) + string(pre) + string(str[index+1:])
		result[2*index+1] = string(str[:index]) + string(next) + string(str[index+1:])
	}

	return result
}

/*
KP
	1. BFS
	2. Optimization. < 10ms
		* store markers with array
		* generate neighbors with integer approach

*/
func openLockV2(deadends []string, target string) int {

	result := 0
	init := "0000"

	// possible values
	// 0 not occurred
	// 1 occurred
	// 2 dead end
	markers := [10000]int{}

	for _, str := range deadends {
		if str == target && str == init {
			return -1
		}
		number, _ := strconv.Atoi(str)
		markers[number] = 2
	}

	initNumber, _ := strconv.Atoi(init)
	targetNumber, _ := strconv.Atoi(target)

	markers[initNumber] = 1

	queue := make([]int, 1)
	queue[0] = initNumber

	for len(queue) > 0 {
		sz := len(queue)
		index := 0
		for ; index < sz; index += 1 {
			number := queue[index]

			if number == targetNumber {
				return result
			}

			nexts := openLockHelperV2(number)
			for _, next := range nexts {
				if markers[next] == 0 {
					queue = append(queue, next)
					markers[next] = 1
				}
			}
		}
		queue = queue[index:]
		result += 1
	}

	return -1
}

func openLockHelperV2(number int) [8]int {

	result := [8]int{}
	changeFactors := [4]int{1000, 100, 10, 1}

	for i, factor := range changeFactors {
		d := (number / factor) % 10
		for j, mul := range [2]int{-1, 1} {
			z := d + mul
			if z == -1 {
				z = 9
			}
			if z == 10 {
				z = 0
			}
			next := number + (z-d)*factor
			result[2*i+j] = next
		}
	}
	return result
}
