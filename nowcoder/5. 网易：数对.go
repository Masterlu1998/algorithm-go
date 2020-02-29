package nowcoder

// 牛牛以前在老师那里得到了一个正整数数对(x, y), 牛牛忘记他们具体是多少了。
//
// 但是牛牛记得老师告诉过他x和y均不大于n, 并且x除以y的余数大于等于k。
// 牛牛希望你能帮他计算一共有多少个可能的数对。
//
// 输入描述:
// 输入包括两个正整数n,k(1 <= n <= 10^5, 0 <= k <= n - 1)。
// 输出描述:
// 对于每个测试用例, 输出一个正整数表示可能的数对数量。
// 示例1：
// 输入：5 2
// 输出：7
// 说明：
// 满足条件的数对有(2,3),(2,4),(2,5),(3,4),(3,5),(4,5),(5,3)

func handle(n, k int) int {

	if k == 0 {
		return n * n
	}

	count := 0
	for x := k + 1; x <= n; x++ {
		// 除数从k+1开始，对于每一个x都有n/x个余数循环，其中从k~x-1共x-k个数符合要求
		count += (n / x) * (x - k)

		// 如果最后一组的最大余数也符合要求，就将k~last计算进去
		last := n % x
		if last >= k {
			count += last - k + 1
		}
	}

	return count
}
