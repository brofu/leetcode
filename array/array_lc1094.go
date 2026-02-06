package array

func carPooling(trips [][]int, capacity int) bool {

	load := make([]int, 10)

	da := NewDifferArray(load)
	for _, trip := range trips {
		da.Increate(trip[1], trip[2]-1, trip[0])
	}
	load = da.Result()
	for _, l := range load {
		if l > capacity {
			return false
		}
	}
	return true
}
