package interviews

import (
	"github.com/brofu/leetcode/common"
)

/*
1. Give an integer, check if it's prime number, if no, return -1, if yes, return the index of it

*/
func CheckPrimeNumberIndex(number int) int {

	if !isPrimeV2(number) {
		return -1
	}
	count := 0
	for start := 2; start <= number; start++ {
		if isPrimeV2(start) {
			count++
		}
	}

	return count
}

/*
KP:
	1. Check a number is prime, need to check sqrt(number), at least
*/
func isPrimeV2(num int) bool {
	if num < 2 {
		return false
	}
	if num == 2 {
		return true
	}
	if num%2 == 0 {
		return false
	}

	for d := 3; d*d <= num; d += 2 {
		if num%d == 0 {
			return false
		}
	}
	return true
}

/*
Wrong version
*/
func isPrime(num int) bool {
	if num < 2 {
		return false
	}
	if num == 2 {
		return true
	}
	for _, n := range []int{2, 3, 5, 7} {
		if num == n {
			return true
		}
		if num%n == 0 {
			return false
		}
	}

	return true
}

/*
KP:
	1. Faster version. 埃氏筛选

*/
func CheckPrimeNumberIndexV2(number int) int {

	if number < 2 {
		return -1
	}
	if number == 2 {
		return 1
	}
	if number%2 == 0 {
		return -1
	}

	isPrime := make([]bool, number+1)

	for p := 2; p <= number; p++ {
		isPrime[p] = true
	}

	for p := 2; p*p <= number; p++ {
		if !isPrime[p] {
			continue
		}
		for m := p * p; m <= number; m += p {
			isPrime[m] = false //flag as not prime
		}
	}

	if !isPrime[number] {
		return -1
	}
	count := 0
	for start := 2; start <= number; start++ {
		if isPrime[start] {
			count++
		}
	}

	return count
}

/*
KP:
	1. A more faster version
	2. 埃氏筛选 + bitset
*/

func CheckPrimeNumberIndexV3(number int) int {

	if number < 2 {
		return -1
	}

	if number == 2 {
		return 1
	}

	isCompBS := common.NewBitSet(number + 1)

	isCompBS.Set(0)
	isCompBS.Set(1)

	// scan and filter all comp numbers
	for p := 2; p*p <= number; p++ {
		if isCompBS.Get(p) {
			continue
		}
		for m := p * p; m <= number; m += p {
			isCompBS.Set(m)
		}
	}

	// check
	if isCompBS.Get(number) {
		return -1
	}
	count := 0
	for i := 2; i <= number; i++ {
		if !isCompBS.Get(i) {
			count++
		}
	}
	return count
}
