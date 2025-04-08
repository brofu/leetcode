package uf

func removeStones(stones [][]int) int {

	locationToNumberMap := make(map[int]int)
	for idx, stone := range stones {
		locationToNumberMap[codeLocationToNumber(stone[0], stone[1])] = idx
	}

	sameRows := make(map[int][]int)
	sameColumns := make(map[int][]int)

	for _, stone := range stones {
		row, column := stone[0], stone[1]
		number := codeLocationToNumber(row, column)
		sameRows[row] = append(sameRows[row], number)
		sameColumns[column] = append(sameColumns[column], number)
	}

	uf := NewPathCompactedUF(len(stones))

	for _, numbers := range sameRows {
		if len(numbers) > 0 {
			for _, number := range numbers {
				uf.Connect(locationToNumberMap[numbers[0]], locationToNumberMap[number])
			}
		}
	}
	for _, numbers := range sameColumns {
		if len(numbers) > 0 {
			for _, number := range numbers {
				uf.Connect(locationToNumberMap[numbers[0]], locationToNumberMap[number])
			}
		}
	}

	return len(stones) - uf.Count()
}

func codeLocationToNumber(x, y int) int {
	return x*10000 + y
}
