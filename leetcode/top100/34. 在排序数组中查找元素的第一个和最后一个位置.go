package top100

// 给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。
//
// 你的算法时间复杂度必须是 O(log n) 级别。
//
// 如果数组中不存在目标值，返回 [-1, -1]。
//
// 示例 1:
//
// 输入: nums = [5,7,7,8,8,10], target = 8
// 输出: [3,4]
// 示例 2:
//
// 输入: nums = [5,7,7,8,8,10], target = 6
// 输出: [-1,-1]

// 思路：正常二分，找到后前后找到上界和下界返回
func searchRange(nums []int, target int) []int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			start, end := mid, mid
			for end < len(nums) && nums[end] == nums[mid] {
				end++
			}
			end--

			for start >= 0 && nums[start] == nums[mid] {
				start--
			}
			start++

			return []int{start, end}
		}
	}

	return []int{-1, -1}
}
