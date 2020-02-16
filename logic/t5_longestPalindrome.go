package logic

import "fmt"

/*
#5 最长回文子串
给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。

示例 1：

输入: "babad"
输出: "bab"
注意: "aba" 也是一个有效答案。
示例 2：

输入: "cbbd"
输出: "bb"

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-palindromic-substring
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

//TestLongestPalindrome 测试 #5最长回文子串
func TestLongestPalindrome() {
	s := "ccc"
	fmt.Println("#5 TestLongestPalindrome Input:")
	fmt.Println(s)
	res := longestPalindrome(s)
	fmt.Println("#5 TestLongestPalindrome Res:")
	fmt.Println(res)
}

func longestPalindrome(s string) string {
	sLen := len(s)
	resStr := ""
	maxLen := 0
	//第一种回文 以某数字为中心 向两边扫描的方法
	for i := 0; i < sLen; i++ {
		m := i
		n := i
		for {
			m--
			n++
			if m < 0 || n >= sLen || s[m] != s[n] {
				if (n - m - 1) > maxLen {
					resStr = s[m+1 : n]
					maxLen = n - m - 1
				}
				break
			}
		}

	}

	//第二种回文 以两个相同的数字分别向两边扫描的方法
	for i := 0; i < sLen; i++ {
		m := i + 1
		n := i
		for {
			m--
			n++
			if m < 0 || n >= sLen || s[m] != s[n] {
				if (n - m - 1) > maxLen {
					resStr = s[m+1 : n]
					maxLen = n - m - 1
				}
				break
			}
		}
	}

	return resStr
}

//该分立计算方式速度超出其他用户90%以上 空间占用超出其他用户100%以上
//可以考虑优化方法并进行合并
