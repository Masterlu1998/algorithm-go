package codinginterviews

// leetcode公共函数和结构体

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Min(left, right int) int {
	if left < right {
		return left
	}
	return right
}

func Sum(p []int) int {
	result := 0
	for i := 0; i < len(p); i++ {
		result += p[i]
	}
	return result
}
