package array

func getModifiedArray(length int, updates [][]int) []int {

	result := make([]int, length)
	da := NewDifferArray(result)
	for _, op := range updates {
		da.Increate(op[0], op[1], op[2])
	}

	return da.Result()
}
