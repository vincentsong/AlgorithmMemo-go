package leetcode

// You are climbing a stair case. It takes n steps to reach to the top.

// Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?

// Note: Given n will be a positive integer.

// Example 1:

// Input: 2
// Output: 2
// Explanation: There are two ways to climb to the top.
// 1. 1 step + 1 step
// 2. 2 steps
// Example 2:

// Input: 3
// Output: 3
// Explanation: There are three ways to climb to the top.
// 1. 1 step + 1 step + 1 step
// 2. 1 step + 2 steps
// 3. 2 steps + 1 step

// https://leetcode-cn.com/problems/climbing-stairs

// Explaination:
// answer is the sum of distinct ways of n-1 stairs and n-2 stairs
// dp(n) = dp(n-1) + dp(n-2)

func climbStairs(n int) int {
	if n < 3 {
		return n
	}
	dp1, dp2 := 1, 2
	ans := 0
	for i := 3; i < n+1; i++ {
		ans = dp1 + dp2
		dp1, dp2 = dp2, ans
	}
	return ans
}
