package top100

// 给定不同面额的硬币 coins 和一个总金额 amount。编写一个函数来计算可以凑成总金额所需的最少的硬币个数。如果没有任何一种硬币组合能组成总金额，返回 -1。
//
// 示例 1:
//
// 输入: coins = [1, 2, 5], amount = 11
// 输出: 3
// 解释: 11 = 5 + 5 + 1
// 示例 2:
//
// 输入: coins = [2], amount = 3
// 输出: -1
// 说明:
// 你可以认为每种硬币的数量是无限的。

// dp[i] = dp[i-val] + 1 val取多个值的时候需要保证最小
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	dp[0] = 0

	max := amount + 1
	for i := 1; i <= amount; i++ {
		dp[i] = max
	}

	for i := 1; i <= amount; i++ {
		for _, val := range coins {
			if i-val >= 0 {
				dp[i] = min(dp[i], dp[i-val]+1)
			}
		}
	}

	if dp[amount] > amount {
		return -1
	}

	return dp[amount]
}
