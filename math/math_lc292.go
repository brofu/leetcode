package math

/**

KP.
	1.	Need to note that,
		a.	For the last round,
			1. If A want to win, there would 1 - 3 stones, then
			2. There should be 4 stones left for B.
		b.	For the 2nd last round,
			1. To leave 4 stones to B, there should be 5 - 7 stones for A to remove.
			2. And before that, there should be 8 stones left for B.
		c.	For the 3rd last round,
			1. To leave 8 stones to B, there should be 9 - 11 stones for A to remove.
			2. And before that, there should be 12 stones left for B.
		d.	And so on.


*/
func canWinNim(n int) bool {
	return n%4 != 0
}
