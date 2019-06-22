package leetcode

// 路径总和
// 给定一个二叉树和一个目标和，判断该树中是否存在根节点到叶子节点的路径，这条路径上所有节点值相加等于目标和。
// 说明: 叶子节点是指没有子节点的节点。
// 示例:
// 给定如下二叉树，以及目标和 sum = 22，
//               5
//              / \
//             4   8
//            /   / \
//           11  13  4
//          /  \      \
//         7    2      1
// 返回 true, 因为存在目标和为 22 的根节点到叶子节点的路径 5->4->11->2。

// 思路1：遍历二叉树每一条到叶子节点的路径，将路径和相加，如果能够达到目标和就返回true，达不到就返回false。

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func hasPathSum(root *TreeNode, sum int) bool {
	return _hasPathSum(root, 0, sum)
}

func _hasPathSum(node *TreeNode, temp, sum int) bool {
	if node == nil {
		return false
	}
	temp += node.Val

	// 提前截断
	if temp > sum {
		return false
	}
	if node.Left == nil && node.Right == nil {
		if temp == sum {
			return true
		}
		return false
	}
	return _hasPathSum(node.Left, temp, sum) || _hasPathSum(node.Right, temp, sum)
}