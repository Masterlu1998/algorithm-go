package leetcode

// 跳跃游戏
// 给定一个非负整数数组，你最初位于数组的第一个位置。
// 数组中的每个元素代表你在该位置可以跳跃的最大长度。
// 判断你是否能够到达最后一个位置。
// 示例 1:
// 输入: [2,3,1,1,4]
// 输出: true
// 解释: 从位置 0 到 1 跳 1 步, 然后跳 3 步到达最后一个位置。
// 示例 2:
// 输入: [3,2,1,0,4]
// 输出: false
// 解释: 无论怎样，你总会到达索引为 3 的位置。但该位置的最大跳跃长度是 0 ， 所以你永远不可能到达最后一个位置


// 思路1：暴力回溯，遍历每一种跳法，超时
var resultCanjump bool

func canJumpFunc1(nums []int) bool {
    if len(nums) == 0 {
        return true
    }
    resultCanjump = false
    _canJump(nums, 0)
    return resultCanjump
}

func _canJump(nums []int, index int) {
    if resultCanjump == true {
        return
    }
    if index >= len(nums) - 1 {
        resultCanjump = true
        return
    }
    
    if index == len(nums) - 1 {
        resultCanjump = true
        return
    }
    
    if nums[index] == 0 {
        return
    }
    
    for i := 1; i <= nums[index]; i++ {
        _canJump(nums, index + i)
    }
    
}

// 思路2：从后向前遍历数组，如果当前元素比长度大就截断数组，长度从当前元素算起，否则继续往前。遍历到第一个元素的时候
// 长度还比1大，说明无法到达最后一个元素。
func canJumpFunc2(nums []int) bool {
	width := 1
	for i := len(nums)-2; i>= 0; i-- {
		if nums[i] >= width {
			width = 1
		} else {
			width++
		}

		if i == 0 && width > 1 {
			return false
		}
	}

	return true
}

