package tree

import (
	"strconv"
	"strings"
)

func restoreIpAddresses(s string) []string {

	result := []string{}
	solution := []string{}
	answer := [][]string{}

	if len(s) < 4 || len(s) > 12 {
		return result
	}

	backTrackIP(s, 4, solution, &answer)

	result = make([]string, len(answer))
	for index, solution := range answer {
		result[index] = strings.Join(solution, ".")
	}
	return result
}

func backTrackIP(s string, layer int, solution []string, answer *[][]string) {

	//base
	if layer == 1 {
		if !isValidIP(s) {
			return
		}

		temp := make([]string, len(solution)+1)
		copy(temp, solution)
		temp[len(solution)] = s
		*answer = append(*answer, temp)
		return
	}

	for _, slice := range []int{1, 2, 3} {
		if len(s) >= slice {
			target := s[:slice]

			if !isValidIP(target) {
				continue
			}

			//choose
			solution = append(solution, target)
			// next step
			backTrackIP(s[slice:], layer-1, solution, answer)

			// cancel choose
			solution = solution[:len(solution)-1]
		}
	}
}

// all the input is digital
func isValidIP(input string) bool {
	if len(input) == 0 || len(input) > 3 {
		return false
	}
	if len(input) > 1 && input[0] == '0' {
		return false
	}
	num, _ := strconv.Atoi(input)
	return num >= 0 && num <= 255
}
