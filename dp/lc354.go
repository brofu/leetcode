package dp

import "sort"

func maxEnvelopes(envelopes [][]int) int {

	sort.Slice(envelopes, func(i, j int) bool {
		return envelopes[i][0] < envelopes[j][0] || envelopes[i][0] == envelopes[j][0] && envelopes[i][1] > envelopes[j][1]
	})

	heightNums := make([]int, len(envelopes))
	for i := 0; i < len(envelopes); i++ {
		heightNums[i] = envelopes[i][1]
	}

	return lengthOfLIS(heightNums)
}
