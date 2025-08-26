package backtrack

import "fmt"

/*
Key Points

backtrack version 1: in line palindrome checking

TC:
1. Possible cutting positions, O(2^(n-1)) = O(2^n). (There are n-1 cutting positions, with n numbers, each position can be cut or not)
2. For each possible cutting positions, copy strings, O(n)
3. For each recursive layer, check palindrome, with O(n), and up to O(n) of recursive depth, as, O(n^2)
4. Overall, O(2^n)*(O(n^2)+O(n)) = O(2^n*n^2)

SC:
1. Recursive depth, O(n)
2. Store the result, O(n) * O(2^n)
3. Overall as O(2^n*n)
*/

func partition(s string) [][]string {

	result := [][]string{}

	var bt func(int, []string)

	bt = func(start int, trace []string) {
		// base case
		if start == len(s) {
			temp := make([]string, len(trace))
			copy(temp, trace)
			result = append(result, temp)
			return
		}

		for idx := start; idx < len(s); idx++ { // up to 2^(n-1) possible cutting position sets

			//pruning
			if !isPalindrome(s[start : idx+1]) { // in-line palindrome checking. O(n)
				continue
			}

			// choose
			bt(idx+1, append(trace, s[start:idx+1])) // up to O(n) recursive depth
			// cancel
		}
	}

	bt(0, []string{})

	return result
}

func isPalindrome(s string) bool {
	if len(s) == 1 {
		return true
	}
	for i, j := 0, len(s)-1; i <= j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}

/*
Key Points

backtrack version 2: store palindrome result

TC:
1. Check the palindrome, and store it, with O(n^3)
1. Possible cutting positions, O(2^(n-1)) = O(2^n). (There are n-1 cutting positions, with n numbers, each position can be cut or not)
2. For each possible cutting positions, copy strings, O(n)
3. For each recursive layer, check palindrome, with O(1)
4. Overall, O(2^n)*(O(1)+O(n)) + O(n^3) = O(2^n*n) + O(n^3)

SC:
1. Recursive depth, O(n)
2. Store the result, O(n) * O(2^n)
3. Overall as O(2^n*n)
*/
func partitionWithCache(s string) [][]string {

	cache := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		cache[i] = make([]bool, len(s))
		for j := i; j < len(s); j++ {
			if isPalindrome(s[i : j+1]) {
				cache[i][j] = true
			}
		}
	}

	result := [][]string{}

	var bt func(int, []string)

	bt = func(start int, trace []string) {
		//base case
		if start == len(s) {
			temp := make([]string, len(trace))
			copy(temp, trace)
			result = append(result, temp)
			return
		}

		for idx := start; idx < len(s); idx++ {

			//pruning
			if !cache[start][idx] {
				continue
			}

			// choose
			bt(idx+1, append(trace, s[start:idx+1]))
			// cancel
		}
	}

	bt(0, []string{})

	return result

}

/*
Key Points

backtrack version 2: store palindrome result, with improved dp array

TC:
1. Check the palindrome, and store it, with O(n^3)
1. Possible cutting positions, O(2^(n-1)) = O(2^n). (There are n-1 cutting positions, with n numbers, each position can be cut or not)
2. For each possible cutting positions, copy strings, O(n)
3. For each recursive layer, check palindrome, with O(1)
4. Overall, O(2^n)*(O(1)+O(n)) + O(n^3) = O(2^n*n) + O(n^3)

SC:
1. Recursive depth, O(n)
2. Store the result, O(n) * O(2^n)
3. Overall as O(2^n*n)
*/

func partitionWithCacheV2(s string) [][]string {

	n := len(s)
	cache := make([][]bool, n)

	fmt.Println(s)
	// construct the cache with O(n^2)
	for i := n - 1; i >= 0; i-- {
		cache[i] = make([]bool, n)
		for j := i; j < n; j++ {
			if s[i] == s[j] && (j-i <= 2 || cache[i+1][j-1]) { // why?
				cache[i][j] = true
			}
		}
	}

	result := make([][]string, 0)

	var bt func(int, []string)

	bt = func(start int, trace []string) {

		// base case
		if start == n {
			temp := make([]string, len(trace))
			copy(temp, trace) // O(n)
			result = append(result, temp)
		}

		for i := start; i < n; i++ {

			// pruning
			if !cache[start][i] { // O(1)
				continue
			}

			// choose
			// next layer
			bt(i+1, append(trace, s[start:i+1])) // up to O(n) recursive depth
			// cancel
		}
	}

	bt(0, []string{})
	return result
}
