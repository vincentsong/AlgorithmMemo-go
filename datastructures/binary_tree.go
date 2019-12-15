package datastructures

import "container/list"

// TreeNode type
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// BuildTreeByBFS use BFS algo to create tree
func BuildTreeByBFS(nodes []*int) *TreeNode {
	length := len(nodes)
	if length == 0 {
		return nil
	}
	idx := 0
	root := TreeNode{*nodes[idx], nil, nil}
	idx++
	curr := &root
	queue := list.New()
	queue.PushBack(curr)
	for queue.Len() > 0 && idx < len(nodes) {
		e := queue.Front()
		node := e.Value.(*TreeNode)
		if node != nil {
			var val = nodes[idx]
			if val != nil {
				node.Left = &TreeNode{*val, nil, nil}
			}
			queue.PushBack(node.Left)
			idx++
			val = nodes[idx]
			if val != nil {
				node.Right = &TreeNode{*val, nil, nil}
			}
			queue.PushBack(node.Right)
		}
		queue.Remove(e)
		idx++
	}
	return &root
}
