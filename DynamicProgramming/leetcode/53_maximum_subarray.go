// Given an integer array nums, find the contiguous subarray (containing at least one number) which has the largest sum and return its sum.

// Example:

// Input: [-2,1,-3,4,-1,2,1,-5,4],
// Output: 6
// Explanation: [4,-1,2,1] has the largest sum = 6.

// ref：https://leetcode-cn.com/problems/maximum-subarray

// Explaination:
// [x,  x,     y,    |    y,     y]
//      ^      ^     |
//     sum  current  |
// if current sum is not greater than 0 then it means we don't need to accumulate it anymore, just set it to current number

package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	ans := nums[0]
	sum := 0
	for _, n := range nums {
		if sum > 0 {
			sum += n
		} else {
			sum = n
		}
		ans = max(ans, sum)
	}
	return ans
}

func assert(condition bool, message string) {
	if !condition {
		panic(message)
	}
}

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	ans := maxSubArray(nums)
	assert(ans == 6, fmt.Sprintf("expect 6 but get %d", ans))
}
