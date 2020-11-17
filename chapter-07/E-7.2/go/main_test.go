package main

import (
	"testing"
)

func TestHierarchySum(t *testing.T) {
	ts := []struct {
		input treeNode
		ouput treeNode
	}{
		{
			input: treeNode{
				value: 2,
				left: &treeNode{
					value: 4,
					left: &treeNode{
						value: 6,
					},
					right: &treeNode{
						value: 9,
						left: &treeNode{
							value: 3,
						},
					},
				},
				right: &treeNode{
					value: 1,
					right: &treeNode{
						value: 2,
					},
				},
			},
			ouput: treeNode{
				value: 27,
				left: &treeNode{
					value: 22,
					left: &treeNode{
						value: 6,
					},
					right: &treeNode{
						value: 12,
						left: &treeNode{
							value: 3,
						},
					},
				},
				right: &treeNode{
					value: 3,
					right: &treeNode{
						value: 2,
					},
				},
			},
		},
		{
			input: treeNode{
				value: -28,
				left: &treeNode{
					value: 6,
					left: &treeNode{
						value: 2,
						left: &treeNode{
							value: 0,
						},
						right: &treeNode{
							value: 1,
						},
					},
					right: &treeNode{
						value: 3,
					},
				},
				right: &treeNode{
					value: 7,
					left: &treeNode{
						value: 4,
					},
					right: &treeNode{
						value: 5,
					},
				},
			},
			ouput: treeNode{
				value: 0,
				left: &treeNode{
					value: 12,
					left: &treeNode{
						value: 3,
						left: &treeNode{
							value: 0,
						},
						right: &treeNode{
							value: 1,
						},
					},
					right: &treeNode{
						value: 3,
					},
				},
				right: &treeNode{
					value: 16,
					left: &treeNode{
						value: 4,
					},
					right: &treeNode{
						value: 5,
					},
				},
			},
		},
	}

	for idx, ttree := range ts {
		HierarchySum(&ttree.input)
		checkTree(t, idx, &ttree.input, &ttree.ouput)
	}
}

// checkTree function check if left subtree and right subtree have the same tree structure and node value
func checkTree(t *testing.T, idx int, left, right *treeNode) {
	if left != nil && right != nil {
		if left.value != right.value {
			t.Errorf("input index=%d, node value not match, %d != %d", idx, left.value, right.value)
		}
	} else if !(left == nil && right == nil) {
		t.Errorf("input index=%d, one of the tree is empty", idx)
	}

	if left.left != nil && right.left != nil {
		checkTree(t, idx, left.left, right.left)
	} else if !(left.left == nil && right.left == nil) {
		t.Errorf("input index=%d, tree structure(left node) not match", idx)
	}
	if left.right != nil && right.right != nil {
		checkTree(t, idx, left.right, right.right)
	} else if !(left.right == nil && right.right == nil) {
		t.Errorf("input index=%d, tree structure(right node) not match", idx)
	}
}
