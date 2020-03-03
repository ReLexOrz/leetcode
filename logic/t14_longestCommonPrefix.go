package logic

import "fmt"

/*
#14 最长公共前缀
编写一个函数来查找字符串数组中的最长公共前缀。

如果不存在公共前缀，返回空字符串 ""。

示例 1:

输入: ["flower","flow","flight"]
输出: "fl"
示例 2:

输入: ["dog","racecar","car"]
输出: ""
解释: 输入不存在公共前缀。
说明:

所有输入只包含小写字母 a-z 。
*/

//TestLongestCommonPrefix 测试 #14最长公共前缀
func TestLongestCommonPrefix() {
	strings := []string{"aaa", "aa", "aaa"}
	fmt.Println("#14 TestLongestCommonPrefix Input:")
	fmt.Println(strings)
	res := longestCommonPrefix2(strings)
	fmt.Println("#14 TestLongestCommonPrefix Res:")
	fmt.Println(res)
}

func longestCommonPrefix(strs []string) string {
	var cpfStr string
	for idx, str := range strs {
		if str == "" {
			return "" //技巧：如果遇到""说明公共前缀一定是不存在的 无需继续向下比较了
		}
		if idx == 0 {
			cpfStr = str
		} else {
			cpfStr = compareCpf(str, cpfStr)
			if cpfStr == "" {
				return "" //技巧如上
			}
		}
	}
	return cpfStr
}

func compareCpf(str string, cpfStr string) string {
	for i := 0; ; {
		if cpfStr[i] != str[i] {
			cpfStr = str[0:i]
			break
		}
		i++
		if i >= len(cpfStr) {
			break
		}
		if i >= len(str) {
			cpfStr = str
			break
		}
	}
	return cpfStr
}

//优化算法 无需任何申请的变量 直接对数组进行操作 原理同上
func longestCommonPrefix2(strs []string) string {
	//空组 直接返回""
	if len(strs) == 0 {
		return ""
	}
	for i := 0; i < len(strs); i++ {
		//遇到空字符串 直接返回
		if strs[i] == "" {
			return ""
		}
		if i == 0 {
			continue
		}
		for j := 0; ; j++ {
			//三种情况 字符比对到末尾
			//前一个字符已经到头了 当前字符保留前缀
			//有不同字符 当前字串保留前缀
			if j >= len(strs[i-1]) {
				strs[i] = strs[i][0:j]
				break
			}
			if j >= len(strs[i]) {
				strs[i] = strs[i][0:j]
				break
			}
			if strs[i-1][j] != strs[i][j] {
				strs[i] = strs[i][0:j]
				break
			}
		}
		if i >= len(strs)-1 {
			return strs[i]
		}
	}
	return strs[0]
}
