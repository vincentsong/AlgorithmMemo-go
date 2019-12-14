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

package leetcode

import "testing"

func Test_maxSubArray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"1", args{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}}, 6},
		{"2", args{[]int{2, 1, -3, -4, -1, 8, 1, 0, 4}}, 13},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSubArray(tt.args.nums); got != tt.want {
				t.Errorf("case %v: maxSubArray() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}
