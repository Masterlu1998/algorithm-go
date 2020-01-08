package top100

// 给定一个字符串 s 和一个非空字符串 p，找到 s 中所有是 p 的字母异位词的子串，返回这些子串的起始索引。
//
// 字符串只包含小写英文字母，并且字符串 s 和 p 的长度都不超过 20100。
//
// 说明：
//
// 字母异位词指字母相同，但排列不同的字符串。
// 不考虑答案输出的顺序。
// 示例 1:
//
// 输入:
// s: "cbaebabacd" p: "abc"
//
// 输出:
// [0, 6]
//
// 解释:
// 起始索引等于 0 的子串是 "cba", 它是 "abc" 的字母异位词。
// 起始索引等于 6 的子串是 "bac", 它是 "abc" 的字母异位词。
//  示例 2:
//
// 输入:
// s: "abab" p: "ab"
//
// 输出:
// [0, 1, 2]
//
// 解释:
// 起始索引等于 0 的子串是 "ab", 它是 "ab" 的字母异位词。
// 起始索引等于 1 的子串是 "ba", 它是 "ab" 的字母异位词。
// 起始索引等于 2 的子串是 "ab", 它是 "ab" 的字母异位词。

// 滑动窗口，先动右指针，满足match，再动左指针，直到不满足match
func findAnagrams(s string, p string) []int {
	have := make(map[byte]int)
	for i := 0; i < len(p); i++ {
		have[p[i]]++
	}

	cache := make(map[byte]int)
	left, right := 0, 0
	match := 0
	result := make([]int, 0)

	for right < len(s) {
		char := s[right]
		if total, ok := have[char]; ok {
			cache[char]++
			if total == cache[char] {
				match++
			}
		}
		right++

		for match == len(have) {
			if (right - left) == len(p) {
				result = append(result, left)
			}

			char := s[left]
			if total, ok := have[char]; ok {
				cache[char]--
				if total > cache[char] {
					match--
				}
			}
			left++
		}
	}

	return result
}
