package interviews

import (
	"strconv"
	"strings"
)

func solutionBCG(capbility []int, logs []string) int {

	p := make([]int, len(capbility))
	t := make([]int, len(capbility))
	closed := make([]bool, len(capbility))
	ti := 0

	for _, log := range logs {

		//
		if ti == len(capbility)-1 && (t[ti] == capbility[ti] || closed[ti] == true) {
			for i := 0; i < len(t); i++ {
				t[i] = 0
			}
			for i := 0; i < len(t); i++ {
				closed[i] = false
			}
			ti = 0
		}

		if strings.HasPrefix(log, "CLOSE") {
			c := strings.Split(log, " ")[1]
			n, _ := strconv.ParseInt(c, 10, 64)
			closed[n] = true
			continue
		}
		if t[ti] >= capbility[ti] {
			ti += 1
			for closed[ti] {
				ti += 1
			}
		}
		t[ti] += 1
		p[ti] += 1
	}

	result := 0
	m := 0
	for idx, v := range p {
		if m <= v {
			m = v
			result = idx
		}
	}

	return result
}
