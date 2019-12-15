package leetcode

import (
	ds "AlgorithmMemo/datastructures"
)

// Given a binary tree, check whether it is a mirror of itself (ie, symmetric around its center).

// For example, this binary tree [1,2,2,3,4,4,3] is symmetric:

//     1
//    / \
//   2   2
//  / \ / \
// 3  4 4  3
//

// But the following [1,2,2,null,3,null,3] is not:

//     1
//    / \
//   2   2
//    \   \
//    3    3
//

// Note:
// Bonus points if you could solve it both recursively and iteratively.

// 链接：https://leetcode-cn.com/problems/symmetric-tree

func isSymmetric(root *ds.TreeNode) bool {
	type TreeNode = ds.TreeNode
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		// check current level is symmetric or not
		left, right := 0, len(queue)-1
		for left < right {
			if (queue[left] == nil && queue[right] != nil) ||
				(queue[left] != nil && queue[right] == nil) ||
				(queue[left] != nil && queue[right] != nil && queue[left].Val != queue[right].Val) {
				return false
			}
			left++
			right--
		}
		tmpQ := make([]*TreeNode, 0)
		for i := 0; i < len(queue); i++ {
			node := queue[i]
			if node != nil {
				tmpQ = append(tmpQ, node.Left)
				tmpQ = append(tmpQ, node.Right)
			}
		}
		queue = tmpQ
	}
	return true
}
