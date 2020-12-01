package main

import (
	"fmt"
	s "strings"
)

type node struct {
	value string
	left  *node
	right *node
}

// RecursionTree function recursively built a binary tree.
// each path from the root to a leaf represents a valid interleaving of a and b.
// Build a binary tree at tnode according to strings a and b.
func RecursionTree(a, b string, tnode *node) {
	aLen, bLen := len(a), len(b)
	if aLen == 0 && bLen == 0 {
		// substrig a and b are both empty, no need to build tree
		return
	}

	if aLen == 0 && bLen > 0 {
		// a is empty but b is not
		nb := b[:1]
		tnode.left = &node{
			value: nb, // put the first character of b at current left node
		}
		RecursionTree(a, b[1:], tnode.left) // recursively build the left subtree
	} else if aLen > 0 && bLen == 0 {
		// b is empty but a is not
		na := a[:1]
		tnode.left = &node{
			value: na, // put the first character of a at current left node
		}
		RecursionTree(a[1:], b, tnode.left) // recursively build the left subtree
	} else {
		// a and b are both not empty
		na, nb := a[:1], b[:1]
		tnode.left = &node{
			value: na, // build left node
		}
		RecursionTree(a[1:], b, tnode.left) // recursively build left subtree

		tnode.right = &node{
			value: nb, // build right node
		}
		RecursionTree(a, b[1:], tnode.right) // recursively build right subtree
	}
}

// walkTreeMain function is the entry for walkTree function
func walkTreeMain(root *node, save *[]string) {
	paths := []string{}
	walkTree(root, save, &paths)
}

// walkTree function recursively walks each path and saves the path string
func walkTree(root *node, save, history *[]string) {
	left, right := root.left, root.right
	if left == nil && right == nil {
		// reach a leaf node, save the path string into the 'save' slice
		path := s.Join(*history, "") // combine the path into a string
		// TODO: use a map to filter duplicated path
		*save = append(*save, path) // save the path string
	} else {
		if left != nil {
			*history = append(*history, left.value)
			walkTree(root.left, save, history)
		}
		if right != nil {
			*history = append(*history, right.value)
			walkTree(root.right, save, history)
		}
	}

	// jump out of current subtree root
	if len(*history) > 0 {
		*history = (*history)[:len(*history)-1]
	}
}

// isInterleaving function check if c is interleaving from a anb b
func isInterleaving(a, b, c string) bool {
	aLen, bLen, cLen := len(a), len(b), len(c)
	if aLen+bLen != cLen {
		return false
	}
	if aLen == 0 || bLen == 0 {
		return true
	} else if aLen == 0 && bLen > 0 {
		return b == c
	} else if aLen > 0 && bLen == 0 {
		return a == c
	}

	// dp[ri][ki]表示a的前ri个字符 和 b的前ki个字符是否能够按照interleaving的方式构成 c的前ri+ki个字符
	dp := make([][]bool, aLen+1)
	for i := range dp {
		dp[i] = make([]bool, bLen+1)
	}

	// ignore b, only consider a and c
	for ri := 1; ri <= aLen; ri++ {
		if a[:ri] == c[:ri] {
			dp[ri][0] = true
		} else {
			dp[ri][0] = false
		}
	}

	// ignore a, only consider b and c
	for ki := 1; ki <= bLen; ki++ {
		if b[:ki] == c[:ki] {
			dp[0][ki] = true
		} else {
			dp[0][ki] = false
		}
	}

	// fill up the dp matrix
	for ri := 1; ri <= aLen; ri++ {
		for ki := 1; ki <= bLen; ki++ {
			// the index of a, b and c, respectively
			ai, bi, ci := ri-1, ki-1, ri+ki-1
			if a[ai] == c[ci] && b[bi] == c[ci] {
				dp[ri][ki] = dp[ri-1][ki] || dp[ri][ki-1]
			} else if a[ai] == c[ci] && b[bi] != c[ci] {
				dp[ri][ki] = dp[ri-1][ki]
			} else if a[ai] != c[ci] && b[bi] == c[ci] {
				dp[ri][ki] = dp[ri][ki-1]
			} else {
				dp[ri][ki] = false
			}
		}
	}

	return dp[aLen][bLen]
}

func main() {
	root := &node{}
	RecursionTree("mn", "xy", root)
	fmt.Printf("hhhhh")
}
