package backtrack

import "math"

/**
KP.
	1. Setup correct connection with `backtrack`. Migrate the debts of current person (with index `cur`) to the next person (with index `cur+1`). Similar to `Dynamic Programming`?
	2. Migrate the debt to the next person, so we have
		a.	bt(cur) = 1 + bt(cur + 1) and if
		b.  leftDebts[cur] == 0, then bt(cur) = bt(cur + 1)
	3. Two import pruning:
		a.  if leftDebts[cur]*leftDebts[cur+1] > 0, can skip the debt migration
		b.  if leftDebts[cur]+leftDets[cur+1] == 0, can break. Since if the min transactions of bt(cur+1) is N, then bt(cur) must be N+1. (The pruning happens after the deeper layer traverse). Similiar to `Dynamic Programming`?

*/
func minTransfers(transactions [][]int) int {

	debts := make([]int, 24)
	for _, tx := range transactions {
		debts[tx[0]] -= tx[2]
		debts[tx[1]] += tx[2]
	}

	leftDebts := []int{}
	for _, debt := range debts {
		if debt != 0 {
			leftDebts = append(leftDebts, debt)
		}
	}

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	var bt func(int) int
	bt = func(pos int) int {

		// base case
		if pos == len(leftDebts) {
			return 0
		}

		cost := math.MaxInt
		current := leftDebts[pos]

		// choose, next layer, and cancel choose
		if current == 0 {
			return bt(pos + 1)
		}

		for i := pos + 1; i < len(leftDebts); i++ {

			// pruning
			if current*leftDebts[i] > 0 {
				continue
			}

			// choose
			leftDebts[i] += current
			// next layer
			cost = min(cost, bt(pos+1)+1)
			// cancel choose
			leftDebts[i] -= current

			// pruning
			if current+leftDebts[i] == 0 {
				break
			}
		}

		return cost
	}

	return bt(0)
}
