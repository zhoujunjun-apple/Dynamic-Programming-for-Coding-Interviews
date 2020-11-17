package main

type treeNode struct {
	value int
	left  *treeNode
	right *treeNode
}

// HierarchySum function use recursion method to solve the problem
// it calculate the sum of root tree and update the value of root node to the sum
func HierarchySum(root *treeNode) {
	HierarchySumRecursion(root)
}

// HierarchySumRecursion function returns the sum of root tree
func HierarchySumRecursion(root *treeNode) int {
	leftSum, rightSum := 0, 0

	// check if child node is empty. get rid of invalid recusive call
	if root.left != nil {
		leftSum = HierarchySumRecursion(root.left)
	}
	if root.right != nil {
		rightSum = HierarchySumRecursion(root.right)
	}

	// update the value of current root node
	root.value = root.value + leftSum + rightSum
	return root.value
}

func main() {

}
