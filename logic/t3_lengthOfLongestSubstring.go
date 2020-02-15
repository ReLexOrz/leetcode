package logic

import "fmt"

/*
#3 无重复字符的最长子串
给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。

示例 1:

输入: "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
示例 2:

输入: "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
示例 3:

输入: "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-substring-without-repeating-characters
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

//TestLengthOfLongestSubstring 测试 #3无重复字符的最长子串
func TestLengthOfLongestSubstring() {
	s := "abcabcbb"
	fmt.Println("#3 TestLengthOfLongestSubstring Input:")
	fmt.Println(s)
	res := lengthOfLongestSubstring(s)
	fmt.Println("#3 TestLengthOfLongestSubstring Res:")
	fmt.Println(res)
}

func lengthOfLongestSubstring(s string) int {
	idxMap := make(map[byte]int, 0)
	maxLen := 0
	minIdx := 0
	maxIdx := 0
	sLen := len(s)
	for ; maxIdx < sLen; maxIdx++ {
		curChar := s[maxIdx]
		existIdx, ok := idxMap[curChar]
		//以前有过这个字符且以前这个字符位置要大于等于前坐标
		if ok && existIdx >= minIdx {
			//比较并记录长度
			curLen := maxIdx - minIdx
			if curLen > maxLen {
				maxLen = curLen
			}
			minIdx = existIdx + 1
		}
		//记录一下有过的字符
		idxMap[curChar] = maxIdx
	}
	//无重复字符的特情处理
	if (maxIdx - minIdx) > maxLen {
		return maxIdx - minIdx
	}
	return maxLen
}
