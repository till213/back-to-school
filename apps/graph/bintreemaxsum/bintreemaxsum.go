package main

import "fmt"

const Min = -1 << 31

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	_, max := pathSum(root)
	return max
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func max3(a, b, c int) int {
	return (max(max(a, b), c))
}

// WRONG! We shall follow *paths*, not entire *subtrees* (except for the 
// **starting node**, from which onwards we follow *all* paths)
func pathSum(node *TreeNode) (sum, newMax int) {
	var leftSum, rightSum int
	var leftMax, rightMax int
	
	if node.Left != nil {
		leftSum, leftMax = pathSum(node.Left)
		leftMax = max3(leftSum, leftMax, node.Val)
		leftMax = max(leftMax, leftSum + node.Val)
	} else {
		leftSum = 0
		leftMax = Min
	}
	if node.Right != nil {
		rightSum, rightMax = pathSum(node.Right)
		rightMax = max3(rightSum, rightMax, node.Val)
		rightMax = max(rightMax, rightSum + node.Val)
	} else {
		rightSum = 0
		rightMax = Min
	}
	sum = node.Val + leftSum + rightSum
	newMax = max3(sum, leftMax, rightMax)

	return
}

var maxSum = Min;

// CORRECT
func maxGain(node *TreeNode) int {
	if (node == nil) {
		return 0
	}

	// max sum on the left and right sub-trees of node
	leftGain := max(maxGain(node.Left), 0)
	rightGain := max(maxGain(node.Right), 0)

	// the price to start a new path where `node` is a highest node
	priceNewPath := node.Val + leftGain + rightGain

	// update max_sum if it's better to start a new path
	maxSum = max(maxSum, priceNewPath);

	// for recursion :
	// return the max gain if continue the same path
	return node.Val + max(leftGain, rightGain);
}

func maxPathSum2(root *TreeNode) int {
	maxGain(root);
	return maxSum;
}

func main() {
	root := new(TreeNode)
	root.Val = 5
	root.Left = new(TreeNode)
	root.Left.Val = 4
	root.Left.Left = new(TreeNode)
	root.Left.Left.Val = 11
	root.Left.Left.Left = new(TreeNode)
	root.Left.Left.Left.Val = 7
	root.Left.Left.Right = new(TreeNode)
	root.Left.Left.Right.Val = 2

	root.Right = new(TreeNode)
	root.Right.Val = 8
	root.Right.Left = new(TreeNode)
	root.Right.Left.Val = 13
	root.Right.Right = new(TreeNode)
	root.Right.Right.Val = 4
	root.Right.Right.Right = new(TreeNode)
	root.Right.Right.Right.Val = 4

	// Wrong
	sum := maxPathSum(root)
	fmt.Println("Max Path Sum (wrong):", sum)

	// Correct
	sum = maxPathSum2(root)
	fmt.Println("Max Path Sum 2 (correct):", sum)
}
