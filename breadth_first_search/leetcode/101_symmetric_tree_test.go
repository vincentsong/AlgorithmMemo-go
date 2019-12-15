package leetcode

import (
	ds "AlgorithmMemo/datastructures"
	"testing"
)

func Test_isSymmetric(t *testing.T) {
	type TreeNode = ds.TreeNode
	type args struct {
		root *TreeNode
	}
	treeNodes := []int{1, 2, 2, 3, 4, 4, 3}
	testNodes := make([](*int), len(treeNodes))
	for i := 0; i < len(treeNodes); i++ {
		testNodes[i] = &treeNodes[i]
	}
	root := ds.BuildTreeByBFS(testNodes)

	treeNodes1 := []int{1, 2, 2, -1, 3, -1, 3}
	testNodes1 := make([](*int), len(treeNodes1))
	for i := 0; i < len(treeNodes1); i++ {
		if treeNodes1[i] == -1 {
			testNodes1[i] = nil
		} else {
			testNodes1[i] = &treeNodes1[i]
		}
	}
	root1 := ds.BuildTreeByBFS(testNodes1)
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{root}, true},
		{"2", args{root1}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSymmetric(tt.args.root); got != tt.want {
				t.Errorf("isSymmetric() = %v, want %v", got, tt.want)
			}
		})
	}
}
