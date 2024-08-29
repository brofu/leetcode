package math

/**

KP.
	How Alice can always win if she can take first?
	1.	Check the stone sum of `ODD` and `EVEN` number index of the piles. That's sum(1st, 3rd, 5th, 7th...) and sum(2nd, 4th, 6th, 8th...)
	2.	If ODD sum > EVEN sum, take from the 1st pile. (Then, Bob can only take a EVEN index pile, the 2nd, or Nth)
	3.  If EVEN sum > ODD sum, take from the last pile (an EVEN pile). (Then, Bob can only take an ODD index pile, the 1st or N-1th)
*/

func stoneGame(piles []int) bool {
	return true
}
