package main

func maxPathSumHelper(root *TreeNode, maxSum *int) int {
	if nil == root {
		return 0
	}

	left := maxPathSumHelper(root.Left, maxSum)
	right := maxPathSumHelper(root.Right, maxSum)

	*maxSum = max(*maxSum, root.Val+left+right)
	return max(0, root.Val+max(left, right))
}

func maxPathSum(root *TreeNode) int {
	maxSum := root.Val
	maxPathSumHelper(root, &maxSum)
	return maxSum
}