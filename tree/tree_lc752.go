package tree

/*
KP
	1. BFS
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
