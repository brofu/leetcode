# Dynamic Programming

**Workflows**
1. Base case
2. Status and Status changing
3. Choice
4. dp function or dp array

**Key Points**

1. 最优子结构
    * If there is 最优子结构
    
2. 状态转移方程
    * (Recurrence or Recursion) 前后项之前关联+base case的描述
    * Examples:
        * Fibonacci Problems: https://labuladong.github.io/algo/images/%E5%8A%A8%E6%80%81%E8%A7%84%E5%88%92%E8%AF%A6%E8%A7%A3%E8%BF%9B%E9%98%B6/fib.png
        * Coins Change Problem: https://labuladong.github.io/algo/images/%E5%8A%A8%E6%80%81%E8%A7%84%E5%88%92%E8%AF%A6%E8%A7%A3%E8%BF%9B%E9%98%B6/coin.png
        
3. Overlapping sub problems. For example, `Fibonacci Number`
    * Solution:  Cache the middle result
    * Optimization: Space used by Cache

4. Memorization or Tabulation. 
    * LC300. `dp[i]`, the length of longest sub sequence ends with `nums[i]` in `nums`
5. Formulation of State Transition Equation 
    * LC300. `dp[i] = max(dp[i], dp[j])+1), where j = [i-1,i-2,…,0] && nums[j] < nums[i]`

**Problmes**
* LC509. Fibonacci Number
    * Recursive with middle result cached.
    * Recurrence
* LC322. Coins Change 
    * Can be resolve by Recursion or Recurrence
* LC300. Longest Increasing Subsequence
    * dp. O(N^2)
    * binary search <= O(N*logN)
* LC354. Russian Doll Envelopes
    * Iiheat from LC300
    * Sort by width and height with different order
