package sol

import (
	"reflect"
	"testing"
)

func CreateBinaryTree(input *[]int) *TreeNode {
	var tree, cur *TreeNode
	hashMap := make(map[int]*TreeNode)
	for idx, val := range *input {
		if idx == 0 {
			tree = &TreeNode{Val: val}
			hashMap[val] = tree
		}
		if node, ok := hashMap[val]; ok {
			cur = node
		} else {
			cur = &TreeNode{Val: val}
			hashMap[val] = cur
		}
		if 2*idx+1 < len(*input) {
			// cur.Left
			left_val := (*input)[2*idx+1]
			if left, exist := hashMap[left_val]; exist {
				cur.Left = left
			} else {
				left := &TreeNode{Val: left_val}
				hashMap[left_val] = left
				cur.Left = left
			}
		}
		if 2*idx+2 < len(*input) {
			right_val := (*input)[2*idx+2]
			if right, exist := hashMap[right_val]; exist {
				cur.Right = right
			} else {
				right := &TreeNode{Val: right_val}
				hashMap[right_val] = right
				cur.Right = right
			}
		}
	}
	return tree
}
func Test_invertTree(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "root = [2,1,3]",
			args: args{root: CreateBinaryTree(&[]int{2, 1, 3})},
			want: CreateBinaryTree(&[]int{2, 3, 1}),
		},
		{
			name: "root = []",
			args: args{root: CreateBinaryTree(&[]int{})},
			want: CreateBinaryTree(&[]int{}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := invertTree(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("invertTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
