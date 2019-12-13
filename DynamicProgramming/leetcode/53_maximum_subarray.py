# -*- coding: utf-8 -*-
# Given an integer array nums, find the contiguous subarray (containing at least one number) which has the largest sum and return its sum.

# Example:

# Input: [-2,1,-3,4,-1,2,1,-5,4],
# Output: 6
# Explanation: [4,-1,2,1] has the largest sum = 6.

# ref：https://leetcode-cn.com/problems/maximum-subarray

# Explaination:
# [x,  x,     y,    |    y,     y]
#      ^      ^     |
#     sum  current  |
# if current sum is not greater than 0 then it means we don't need to accumulate it anymore, just set it to current number

from typing import List

class Solution:
    def maxSubArray(self, nums: List[int]) -> int:
        ans = nums[0]
        sum = 0
        for n in nums:
            if sum > 0:
                sum += n
            else:
                sum = n
            ans = max(sum, ans)
        return ans


# test
s = Solution()
nums = [-2, 1, -3, 4, -1, 2, 1, -5, 4]
assert(s.maxSubArray(nums) == 6)