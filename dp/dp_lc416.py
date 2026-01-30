
from typing import List 

"""
KP:
TC:
    1. O(M*N), M = len(nums)+1, N = sum(nums) // 2 + 1

SC:
    1. O(N), N = sum(nums) //2 +1
"""
class Solution:
    def canPartition(self, nums: List[int]) -> bool:

        total = sum(nums)
        if total % 2 != 0:
            return False

        n = total // 2
        dp = [False] * (n+1)
        dp[0] = True

        for i in range(1, len(nums)):
            wi = nums[i-1]
            for w in range(n, wi-1, -1):
                dp[w] = dp[w] or dp[w-wi]

        return dp[n]

