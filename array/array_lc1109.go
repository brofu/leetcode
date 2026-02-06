package array

func corpFlightBookings(bookings [][]int, n int) []int {

	flights := make([]int, n+1)
	da := NewDifferArray(flights)
	for _, book := range bookings {
		da.Increate(book[0], book[1], book[2])
	}

	r := da.Result()
	return r[1:]
}
