package tree

func permute(nums []int) [][]int {

	result := [][]int{}
	track := []int{}
	used := make(map[int]struct{}, len(nums))

	backtrack(nums, track, used, &result)
	return result
}

func backtrack(nums, track []int, used map[int]struct{}, answer *[][]int) {

	if len(track) == len(nums) {
		temp := make([]int, len(track))
		//temp := []int{}
		//n := copy(temp, track)
		copy(temp, track)
		*answer = append(*answer, temp)
		return
	}

	for _, num := range nums {
		if _, ok := used[num]; ok {
			continue
		}

		used[num] = struct{}{}
		track = append(track, num)

		backtrack(nums, track, used, answer)

		track = track[:len(track)-1]
		delete(used, num)
	}
}
