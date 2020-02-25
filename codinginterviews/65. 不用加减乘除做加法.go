package codinginterviews

// 写一个函数，求两个整数之和，要求在函数体内不得使用 “+”、“-”、“*”、“/” 四则运算符号。
//
//
//
// 示例:
//
// 输入: a = 1, b = 1
// 输出: 2
//
//
// 提示：
//
// a, b 均可能是负数或 0
// 结果不会溢出 32 位整数

func add(a int, b int) int {
	res := a ^ b
	more := a & b << 1
	if more != 0 {
		return add(res, more)
	}

	return res
}
